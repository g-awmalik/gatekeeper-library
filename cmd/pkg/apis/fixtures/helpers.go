package fixtures

import (
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func yamlToUnstructured(data []byte) (*unstructured.Unstructured, error) {
	u := &unstructured.Unstructured{}

	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, fmt.Errorf("YAMLToJSON: %w", err)
	}

	err = u.UnmarshalJSON(jsonData)
	if err != nil {
		return nil, fmt.Errorf("UnmarshalJSON: %w", err)
	}

	return u, nil
}
