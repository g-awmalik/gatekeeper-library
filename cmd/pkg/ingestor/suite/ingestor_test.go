package suite

import (
	"gatekeeperlibrary/cmd/pkg/apis"
	"gatekeeperlibrary/cmd/pkg/apis/fixtures"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIngestor(t *testing.T) {
	tcs := []struct {
		name string
		path string
		want []*apis.ConstraintTemplateDoc
	}{
		{
			name: "flat file structure",
			path: "./testdata/flat-folder-structure",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.DisallowedAuthzPrefix,
			},
		},
		{
			name: "traditional file structure",
			path: "./testdata/traditional-folder-structure",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.K8sPSPProcMount,
			},
		},
		{
			name: "multiple suites for one template",
			path: "./testdata/multiple-suites-for-one-template",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.K8sPodsRequireSecurityContext,
			},
		},
		{
			name: "multiple tests for one constraint",
			path: "./testdata/multiple-tests-for-one-constraint",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.K8sRestrictLabels,
			},
		},
		{
			name: "constraint with referential data",
			path: "./testdata/referential-data",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.K8sRequireNamespaceNetworkPolicies,
			},
		},
		{
			name: "all",
			path: "./testdata",
			want: []*apis.ConstraintTemplateDoc{
				fixtures.DisallowedAuthzPrefix,
				fixtures.K8sPSPProcMount,
				fixtures.K8sPodsRequireSecurityContext,
				fixtures.K8sRestrictLabels,
				fixtures.K8sRequireNamespaceNetworkPolicies,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// call the ingestor
			got, err := Ingest(tc.path)
			if err != nil {
				t.Fatal(err)
			}

			// sort both got and want
			sort.Slice(got, func(i, j int) bool {
				return got[i].Name < got[j].Name
			})
			sort.Slice(tc.want, func(i, j int) bool {
				return tc.want[i].Name < tc.want[j].Name
			})

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Ingest() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
