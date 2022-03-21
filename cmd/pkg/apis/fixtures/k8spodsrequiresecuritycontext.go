package fixtures

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

var K8sPodsRequireSecurityContext *apis.ConstraintTemplateDoc

const k8sPodsRequireSecurityContextConstraint = `
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPodsRequireSecurityContext
metadata:
  name: pods-require-security-context
`

const k8sPodsRequireSecurityContextAllowed = `
apiVersion: v1
kind: Pod
metadata:
  name: allowed-example
spec:
  containers:
    - name: nginx
      image: nginx
      securityContext:
        runAsUser: 2000
`

const k8sPodsRequireSecurityContextDisallowed = `
apiVersion: v1
kind: Pod
metadata:
  name: disallowed-example
spec:
  containers:
    - name: nginx
      image: nginx
`

func init() {
	constraint, err := yamlToUnstructured([]byte(k8sPodsRequireSecurityContextConstraint))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	allowed, err := yamlToUnstructured([]byte(k8sPodsRequireSecurityContextAllowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	disallowed, err := yamlToUnstructured([]byte(k8sPodsRequireSecurityContextDisallowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}

	K8sPodsRequireSecurityContext = &apis.ConstraintTemplateDoc{
		Name:        "K8sPodsRequireSecurityContext",
		Description: "Requires all Pods and containers to have a SecurityContext defined at the Pod or container level.",
		Validation: &templates.Validation{
			LegacySchema:    pointer.Bool(true),
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{XPreserveUnknownFields: pointer.BoolPtr(true)},
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
