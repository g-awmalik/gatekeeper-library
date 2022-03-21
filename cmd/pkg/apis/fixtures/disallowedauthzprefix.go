package fixtures

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

var DisallowedAuthzPrefix *apis.ConstraintTemplateDoc

const disallowedAuthzPrefixConstraint = `
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: DisallowedAuthzPrefix
metadata:
  name: disallowed-authz-prefix-constraint
spec:
  enforcementAction: dryrun
  match:
    kinds:
      - apiGroups: ["security.istio.io"]
        kinds: ["AuthorizationPolicy"]
  parameters:
    disallowedprefixes: ["badprefix", "reallybadprefix"]
`

const disallowedAuthzPrefixDisallowed = `
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: bad-source-namespace
  namespace: foo
spec:
  selector:
    matchLabels:
      app: httpbin
      version: v1
  rules:
    - from:
        - source:
            principals: ["cluster.local/ns/default/sa/sleep"]
        - source:
            namespaces: ["badprefix-test"]
      to:
        - operation:
            methods: ["GET"]
            paths: ["/info*"]
        - operation:
            methods: ["POST"]
            paths: ["/data"]
      when:
        - key: request.auth.claims[iss]
          values: ["https://accounts.google.com"]
`

func init() {
	constraint, err := yamlToUnstructured([]byte(disallowedAuthzPrefixConstraint))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	disallowed, err := yamlToUnstructured([]byte(disallowedAuthzPrefixDisallowed))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}

	DisallowedAuthzPrefix = &apis.ConstraintTemplateDoc{
		Name:        "DisallowedAuthzPrefix",
		Description: "Requires that principals and namespaces in Istio `AuthorizationPolicy` rules not have a prefix from a specified list.\nhttps://istio.io/latest/docs/reference/config/security/authorization-policy/",
		Validation: &templates.Validation{
			LegacySchema: pointer.Bool(true),
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
				XPreserveUnknownFields: pointer.BoolPtr(true),
				Properties: map[string]apiextensions.JSONSchemaProps{
					"disallowedprefixes": {
						Type:        "array",
						Description: "Disallowed prefixes of principals and namespaces.",
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
				Disallowed: []*apis.ObjectWithContext{{Object: disallowed}},
			},
		},
	}
}
