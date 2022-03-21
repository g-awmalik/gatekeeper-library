package fixtures

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

var K8sPSPProcMount *apis.ConstraintTemplateDoc

const k8spspprocmountConstraint = `
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sPSPProcMount
metadata: # kpt-merge: /psp-proc-mount
  name: psp-proc-mount
spec:
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
  parameters:
    procMount: Default
`

const k8spspprocmountDisallowed = `
apiVersion: v1
kind: Pod
metadata: # kpt-merge: /nginx-proc-mount-disallowed
  name: nginx-proc-mount-disallowed
  labels:
    app: nginx-proc-mount
spec:
  containers:
  - name: nginx
    image: nginx
    securityContext:
      procMount: Unmasked #Default
`

func init() {
	constraint, err := yamlToUnstructured([]byte(k8spspprocmountConstraint))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	disallowed, err := yamlToUnstructured([]byte(k8spspprocmountDisallowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}

	K8sPSPProcMount = &apis.ConstraintTemplateDoc{
		Name:        "K8sPSPProcMount",
		Description: "Controls the allowed `procMount` types for the container. Corresponds to the `allowedProcMountTypes` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#allowedprocmounttypes",
		Validation: &templates.Validation{
			LegacySchema: pointer.Bool(true),
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
				XPreserveUnknownFields: pointer.BoolPtr(true),
				Type:                   "object",
				Description:            "Controls the allowed `procMount` types for the container. Corresponds to the `allowedProcMountTypes` field in a PodSecurityPolicy. For more information, see https://kubernetes.io/docs/concepts/policy/pod-security-policy/#allowedprocmounttypes",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"procMount": {
						Type:        "string",
						Description: "Defines the strategy for the security exposure of certain paths in `/proc` by the container runtime. Setting to `Default` uses the runtime defaults, where `Unmasked` bypasses the default behavior.",
						Enum:        []apiextensions.JSON{"Default", "Unmasked"},
					},
				},
			},
		},
		Samples: []*apis.Sample{
			{
				Constraint: &apis.ObjectWithContext{Object: constraint},
				Disallowed: []*apis.ObjectWithContext{{Object: disallowed}},
			},
		},
	}
}
