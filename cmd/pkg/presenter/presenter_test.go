package presenter

import (
	"gatekeeperlibrary/cmd/pkg/apis"
	"gatekeeperlibrary/cmd/pkg/apis/fixtures"
	"os"
	"strings"
	"testing"
)

var (
	k8srestrictlabelsMD                  string
	k8spodsrequiresecuritycontextMD      string
	disallowedauthzprefixMD              string
	doubleinfoMD                         string
	matchText                            string
	k8srequirenamespacenetworkpoliciesMD string
)

func init() {
	data, err := os.ReadFile("./test_data/k8srestrictlabels.md")
	if err != nil {
		panic(err)
	}
	k8srestrictlabelsMD = strings.TrimSpace(string(data))

	data, err = os.ReadFile("./test_data/k8spodsrequiresecuritycontext.md")
	if err != nil {
		panic(err)
	}
	k8spodsrequiresecuritycontextMD = strings.TrimSpace(string(data))

	data, err = os.ReadFile("./test_data/disallowedauthzprefix.md")
	if err != nil {
		panic(err)
	}
	disallowedauthzprefixMD = strings.TrimSpace(string(data))

	data, err = os.ReadFile("./test_data/double-info.md")
	if err != nil {
		panic(err)
	}
	doubleinfoMD = strings.TrimSpace(string(data))

	data, err = os.ReadFile("./test_data/match.md")
	if err != nil {
		panic(err)
	}
	matchText = strings.TrimSpace(string(data))

	data, err = os.ReadFile("./test_data/k8srequirenamespacenetworkpolicies.md")
	if err != nil {
		panic(err)
	}
	k8srequirenamespacenetworkpoliciesMD = strings.TrimSpace(string(data))
}

func TestFullSingleConstraintTemplate(t *testing.T) {
	tcs := []struct {
		name  string
		input []*apis.ConstraintTemplateDoc
		want  string
	}{
		{
			name:  "Single Info - k8srestrictlabels",
			input: []*apis.ConstraintTemplateDoc{fixtures.K8sRestrictLabels},
			want:  k8srestrictlabelsMD,
		},
		{
			name:  "Single Info - k8spodsrequiresecuritycontext",
			input: []*apis.ConstraintTemplateDoc{fixtures.K8sPodsRequireSecurityContext},
			want:  k8spodsrequiresecuritycontextMD,
		},
		{
			name:  "Single Info - disallowedauthzprefix",
			input: []*apis.ConstraintTemplateDoc{fixtures.DisallowedAuthzPrefix},
			want:  disallowedauthzprefixMD,
		},
		{
			name:  "Single Info - k8srequirenamespacenetworkpolicies",
			input: []*apis.ConstraintTemplateDoc{fixtures.K8sRequireNamespaceNetworkPolicies},
			want:  k8srequirenamespacenetworkpoliciesMD,
		},
		{
			name:  "Double Info - k8srestrictlabels, k8spodsrequiresecuritycontext",
			input: []*apis.ConstraintTemplateDoc{fixtures.K8sRestrictLabels, fixtures.K8sPodsRequireSecurityContext},
			want:  doubleinfoMD,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Present(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			doubleDiff(t, "Present", tc.want, got)
		})
	}
}

func TestMatch(t *testing.T) {
	tcs := []struct {
		name string
		want string
	}{
		{
			name: "base case",
			want: matchText,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Match()
			if err != nil {
				t.Fatal(err)
			}

			doubleDiff(t, "Match", tc.want, got)
		})
	}
}
