package presenter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kylelemons/godebug/diff"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

func TestHandleBlankSchema(t *testing.T) {
	tcs := []struct {
		name       string
		validation *templates.Validation
		want       string
		wantErr    bool
	}{
		{
			name: "legacySchema is true, schema is nil",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(true),
				OpenAPIV3Schema: nil,
			},
			wantErr: true,
		},
		{
			name: "legacySchema is true, schema is empty",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(true),
				OpenAPIV3Schema: empty,
			},
		},
		{
			name: "legacySchema is true, schema is empty with XPreserveUnknownFields: true",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(true),
				OpenAPIV3Schema: emptyWithXPreserve,
			},
		},
		{
			name: "legacySchema is true, schema is empty with XPreserveUnknownFields: true and a description",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(true),
				OpenAPIV3Schema: descriptiveEmptyWithXPreserve,
			},
		},
		{
			name: "legacySchema is false, schema is nil",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(false),
				OpenAPIV3Schema: nil,
			},
			wantErr: true,
		},
		{
			name: "legacySchema is false, schema is empty",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(false),
				OpenAPIV3Schema: empty,
			},
		},
		{
			name: "legacySchema is false, schema is empty with XPreserveUnknownFields: true and no description",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(false),
				OpenAPIV3Schema: emptyWithXPreserve,
			},
			wantErr: true,
		},
		{
			name: "legacySchema is false, schema is empty with XPreserveUnknownFields: true and and a description",
			validation: &templates.Validation{
				LegacySchema:    pointer.Bool(false),
				OpenAPIV3Schema: descriptiveEmptyWithXPreserve,
			},
			want: descriptiveEmptyWithXPreserveText,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := handleBlankSchema(tc.validation)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Errorf("generating handleBlankSchema: %q", err)
			} else if tc.wantErr {
				t.Errorf("want non-nil err, got: %q", err)
			}

			doubleDiff(t, "handleBlankSchema", tc.want, got)
		})
	}
}

func TestParametersTextSlice(t *testing.T) {
	tcs := []struct {
		name    string
		schema  *apiextensions.JSONSchemaProps
		want    string
		wantErr bool
	}{
		{
			name:    "A nil schema",
			schema:  nilSchema,
			want:    "",
			wantErr: true,
		},
		{
			name:   "An empty schema",
			schema: empty,
			want:   "<untyped>",
		},
		{
			name:   "An empty schema with XPreserveUnknownFields: true",
			schema: emptyWithXPreserve,
			want:   emptyWithXPreserveText,
		},
		{
			name:   "Single primitive",
			schema: singlePrimitive,
			want:   singlePrimitiveText,
		},
		{
			name:   "Map with empty value",
			schema: mapWithEmptyVal,
			want:   mapWithEmptyValText,
		},
		{
			name:   "Map with primitive types",
			schema: primitiveTypeMap,
			want:   primitiveTypeMapText,
		},
		{
			name:   "Map with primitive types and descriptions",
			schema: descriptivePrimitiveTypeMap,
			want:   descriptivePrimitiveTypeMapText,
		},
		{
			name:   "Map with primitive types and an enum",
			schema: enumPrimitiveTypeMap,
			want:   enumPrimitiveTypeMapText,
		},
		{
			name:   "Map with primitive types, an enum, and descriptions",
			schema: descriptiveEnumPrimitiveTypeMap,
			want:   descriptiveEnumPrimitiveTypeMapText,
		},
		{
			name:   "Array of primitive type",
			schema: primitiveTypeArray,
			want:   primitiveTypeArrayText,
		},

		{
			name:   "Array of primitive type with description",
			schema: descriptivePrimitiveTypeArray,
			want:   descriptivePrimitiveTypeArrayText,
		},

		{
			name:   "Array of primitive type with enum",
			schema: enumPrimitiveTypeArray,
			want:   enumPrimitiveTypeArrayText,
		},
		{
			name:   "Array of primitive type with enum and description",
			schema: descriptiveEnumPrimitiveTypeArray,
			want:   descriptiveEnumPrimitiveTypeArrayText,
		},

		{
			name:   "Array of maps",
			schema: arrayWithMap,
			want:   arrayWithMapText,
		},
		{
			name:   "Array of maps with description",
			schema: descriptiveArrayWithMap,
			want:   descriptiveArrayWithMapText,
		},
		{
			name:   "Nested maps",
			schema: nestedMap,
			want:   nestedMapText,
		},
		{
			name:   "Nested maps with description",
			schema: descriptiveNestedMap,
			want:   descriptiveNestedMapText,
		},
		{
			name:   "Map containing an array of primitives",
			schema: primitiveArrayWithinMap,
			want:   primitiveArrayWithinMapText,
		},
		{
			name:   "Map containing an array of primitives with description",
			schema: descriptivePrimitiveArrayWithinMap,
			want:   descriptivePrimitiveArrayWithinMapText,
		},
		{
			name:   "Map containing an array of maps",
			schema: mapSchemaArrayWithinMap,
			want:   mapSchemaArrayWithinMapText,
		},
		{
			name:   "Map containing an array of maps with descriptions",
			schema: descriptiveMapSchemaArrayWithinMap,
			want:   descriptiveMapSchemaArrayWithinMapText,
		},
		{
			name:   "AdditionalProperties of primitive type",
			schema: additionalPropertiesPrimitive,
			want:   additionalPropertiesPrimitiveText,
		},
		{
			name:   "AdditionalProperties of primitive type with description",
			schema: descriptiveAdditionalPropertiesPrimitive,
			want:   descriptiveAdditionalPropertiesPrimitiveText,
		},
		{
			name:   "AdditionalProperties of map type",
			schema: additionalPropertiesMap,
			want:   additionalPropertiesMapText,
		},
		{
			name:   "AdditionalProperties of map type with descriptions",
			schema: descriptiveAdditionalPropertiesMap,
			want:   descriptiveAdditionalPropertiesMapText,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			lines, err := parametersTextSlice(tc.schema, 0)
			if err != nil {
				if tc.wantErr {
					return
				}
				t.Errorf("generating parametersTextSlice: %q", err)
			} else if tc.wantErr {
				t.Errorf("want non-nil err, got: %q", err)
			}
			got := strings.Join(lines, "\n")
			doubleDiff(t, "parametersTextSlice", tc.want, got)
		})
	}
}

func TestWrapComment(t *testing.T) {
	tcs := []struct {
		name       string
		unwrapped  string
		lineLength int
		want       string
		wantErr    bool
	}{
		{
			name:       "comment that is shorter than length is not wrapped",
			unwrapped:  "# pizza foo bar",
			lineLength: 1000,
			want:       "# pizza foo bar",
		},
		{
			name:       "comment that is longer than length is wrapped",
			unwrapped:  unwrappedLongComment,
			lineLength: 80,
			want:       wrappedLongComment,
		},
		{
			name:       "comment containing a token that's longer than lineLength is wrapped correctly",
			unwrapped:  unwrappedLongCommentWithLargeToken,
			lineLength: 80,
			want:       wrappedLongCommentWithLargeToken,
		},
		{
			name:       "comment that contains a pound character is wrapped correctly",
			unwrapped:  unwrappedLongCommentWithPound,
			lineLength: 80,
			want:       wrappedLongCommentWithPound,
		},
		{
			name:       "line length zero",
			lineLength: 0,
			wantErr:    true,
		},
		{
			name:       "input is not a comment",
			lineLength: 100,
			wantErr:    true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := wrapComment(tc.unwrapped, tc.lineLength)
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("wanted nil err but got: %v", err)
				}
			} else if tc.wantErr {
				t.Fatalf("wanted err but got nil")
			}

			doubleDiff(t, "wrapComment", tc.want, strings.Join(got, "\n"))
		})
	}
}

const (
	unwrappedLongComment = "# namespaces <array>: namespaces is a list of namespace names. If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix-based glob.  For example, namespaces: [kube-*] matches both kube-system and kube-public."
	wrappedLongComment   = `# namespaces <array>: namespaces is a list of namespace names. If defined, a
# constraint only applies to resources in a listed namespace.  Namespaces also
# supports a prefix-based glob.  For example, namespaces: [kube-*] matches both
# kube-system and kube-public.`

	unwrappedLongCommentWithLargeToken = "# namespaces <array>: namespaces is a list of namespace names. https://team-review.git.corp.google.com/c/nomos-team/policy-controller-constraint-library/+/1263609 If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix-based glob.  For example, namespaces: [kube-*] matches both kube-system and kube-public.  "
	wrappedLongCommentWithLargeToken   = `# namespaces <array>: namespaces is a list of namespace names.
# https://team-review.git.corp.google.com/c/nomos-team/policy-controller-constraint-library/+/1263609
# If defined, a constraint only applies to resources in a listed namespace.
# Namespaces also supports a prefix-based glob.  For example, namespaces:
# [kube-*] matches both kube-system and kube-public.`

	unwrappedLongCommentWithPound = "# namespaces <array>: namespaces is a list of # namespace names. If #defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix-based glob.  For example, namespaces: [kube-*] matches both kube-system and kube-public."
	wrappedLongCommentWithPound   = `# namespaces <array>: namespaces is a list of # namespace names. If #defined, a
# constraint only applies to resources in a listed namespace.  Namespaces also
# supports a prefix-based glob.  For example, namespaces: [kube-*] matches both
# kube-system and kube-public.`
)

// doubleDiff runs both cmp.Diff and diff.Diff when want and got are unequal.
// In comparing multiline strings, both of these tools are imperfect.  But,
// having the output of both yields a significantly easier debugging
// experience.
func doubleDiff(t *testing.T, funcName, want, got string) {
	if cmp.Equal(want, got) {
		return
	}

	fmt.Println(got)

	t.Errorf("%s() mismatch (-want +got):\n%q", funcName, cmp.Diff(want, got))
	t.Error("\n" + diff.Diff(want, got))
}
