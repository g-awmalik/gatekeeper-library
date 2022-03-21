package apis

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

// ConstraintTemplateDoc is all the meaningful information from one CT library template directory,
// including the template and relevant examples.
type ConstraintTemplateDoc struct {
	Name        string
	Description string
	Validation  *templates.Validation
	Samples     []*Sample
}

// Sample is single example Constraint and the allowed and disallowed objects that provide context.
type Sample struct {
	Constraint *ObjectWithContext
	Allowed    []*ObjectWithContext
	Disallowed []*ObjectWithContext
}

// Merge is a simple combination of two Sample structs.  It makes sure that
// both Samples have valid Constraints and that the constraints have no
// cmp.Diff output.  It does not provide assurances for the Allowed or
// Disallowed objects, doing a simple append of the slices of those objects.
func Merge(s1 *Sample, s2 *Sample) (*Sample, error) {
	if s1.Constraint == nil {
		return nil, fmt.Errorf("s1 Constraint cannot be nil")
	}
	if s2.Constraint == nil {
		return nil, fmt.Errorf("s2 Constraint cannot be nil")
	}

	if diff := cmp.Diff(s1.Constraint, s2.Constraint); diff != "" {
		return nil, fmt.Errorf("s1 and s2 Constraint values must be same to merge: %v", diff)
	}

	return &Sample{
		Constraint: s1.Constraint,
		Allowed:    append(s1.Allowed, s2.Allowed...),
		Disallowed: append(s1.Disallowed, s2.Disallowed...),
	}, nil
}

// ObjectWithContext is a struct that contains the unstructured content of a
// k8s object and any other meaningful fields that might be useful to a
// consumer.
type ObjectWithContext struct {
	Object          *unstructured.Unstructured
	ReferentialData []*unstructured.Unstructured
}

// ObjectString returns the ObjectWithContext Object field as a string
func (owc *ObjectWithContext) ObjectString() (string, error) {
	if owc == nil {
		return "", fmt.Errorf("unable to convert nil to string")
	}

	bytes, err := yaml.Marshal(owc.Object)
	if err != nil {
		return "", fmt.Errorf("marshalling object: %w", err)
	}

	return string(bytes), nil
}

// ReferentialDataStrings returns the ObjectWithContext ReferentialData field
// as a slice of strings
func (owc *ObjectWithContext) ReferentialDataStrings() ([]string, error) {
	if owc == nil {
		return nil, fmt.Errorf("unable to convert nil to string")
	}

	strings := make([]string, 0, len(owc.ReferentialData))

	for _, data := range owc.ReferentialData {
		bytes, err := yaml.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("marshalling referential data: %w", err)
		}

		strings = append(strings, string(bytes))
	}

	return strings, nil
}

func (owc *ObjectWithContext) Name() string {
	if owc == nil || owc.Object == nil {
		return ""
	}

	return owc.Object.GetName()
}
