package fixtures

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

var K8sRestrictLabels *apis.ConstraintTemplateDoc

const k8sRestrictLabelsConstraint = `
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRestrictLabels
metadata:
  name: restrict-label-example
spec:
  enforcementAction: dryrun
  parameters:
    restrictedLabels:
      - label-example
    exceptions:
      - group: ""
        kind: Pod
        namespace: default
        name: allowed-example
`

const k8sRestrictLabelsAllowed = `
apiVersion: v1
kind: Pod
metadata:
  name: allowed-example
  namespace: default
  labels:
    label-example: example
spec:
  containers:
  - name: nginx
    image: nginx
`

const k8sRestrictLabelsDisallowed = `
apiVersion: v1
kind: Pod
metadata:
  name: disallowed-example
  namespace: default
  labels:
    label-example: example
spec:
  containers:
  - name: nginx
    image: nginx
`

func init() {
	constraint, err := yamlToUnstructured([]byte(k8sRestrictLabelsConstraint))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	allowed, err := yamlToUnstructured([]byte(k8sRestrictLabelsAllowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	disallowed, err := yamlToUnstructured([]byte(k8sRestrictLabelsDisallowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}

	K8sRestrictLabels = &apis.ConstraintTemplateDoc{
		Name:        "K8sRestrictLabels",
		Description: "Disallows resources with any of the specified `restrictedLabels`. Matches on label key names only.  Single object exceptions can be included, identified by their group, kind, namespace, and name.",
		Validation: &templates.Validation{
			LegacySchema: pointer.Bool(false),
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"exceptions": {
						Type:        "array",
						Description: "A list of objects that are exempted from the label restrictions.",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type:        "object",
								Description: "A single object's identification, based on group, kind, namespace, and name.",
								Properties: map[string]apiextensions.JSONSchemaProps{
									"group":     {Type: "string"},
									"kind":      {Type: "string"},
									"name":      {Type: "string"},
									"namespace": {Type: "string"},
								},
							},
						},
					},
					"restrictedLabels": {
						Type:        "array",
						Description: "A list of label keys strings.",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
				},
			},
		},
		Samples: []*apis.Sample{
			{
				Constraint: &apis.ObjectWithContext{Object: constraint},
				Allowed:    []*apis.ObjectWithContext{{Object: allowed}},
				Disallowed: []*apis.ObjectWithContext{{Object: disallowed}},
			},
		},
	}
}
