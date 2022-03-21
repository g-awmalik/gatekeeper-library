package fixtures

import (
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/pointer"
)

var K8sRequireNamespaceNetworkPolicies *apis.ConstraintTemplateDoc

const constraintText = `
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sRequireNamespaceNetworkPolicies
metadata:
  name: require-namespace-network-policies
spec:
  enforcementAction: dryrun
`

const namespaceText = `
apiVersion: v1
kind: Namespace
metadata:
  name: require-namespace-network-policies-example
`

const networkPolicyText = `
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: require-namespace-network-policies-example
`

const description = `Requires that every namespace defined in the cluster has a NetworkPolicy.
Note: This constraint is referential. See https://cloud.google.com/anthos-config-management/docs/how-to/creating-constraints#referential for details.`

func init() {
	constraint, err := yamlToUnstructured([]byte(constraintText))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	namespace, err := yamlToUnstructured([]byte(namespaceText))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}
	networkPolicy, err := yamlToUnstructured([]byte(networkPolicyText))
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal yaml: %v", err))
	}

	K8sRequireNamespaceNetworkPolicies = &apis.ConstraintTemplateDoc{
		Name:        "K8sRequireNamespaceNetworkPolicies",
		Description: description,
		Validation: &templates.Validation{
			LegacySchema:    pointer.Bool(true),
			OpenAPIV3Schema: &apiextensions.JSONSchemaProps{XPreserveUnknownFields: pointer.BoolPtr(true)},
		},
		Samples: []*apis.Sample{
			{
				Constraint: &apis.ObjectWithContext{Object: constraint},
				Allowed: []*apis.ObjectWithContext{{
					Object:          namespace,
					ReferentialData: []*unstructured.Unstructured{networkPolicy},
				}},
				Disallowed: []*apis.ObjectWithContext{{Object: namespace}},
			},
		},
	}
}
