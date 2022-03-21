package presenter

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
)

const untyped = "untyped"

type templateContext struct {
	Name          string
	KeyName       string
	Type          string
	Description   string
	Value         string
	AllowedValues []string
}

// ExecuteToString injects templateContext data into the passed-in template,
// returning the resulting string.  It returns an error if the template fails
// to evaluate.
func (tc *templateContext) ExecuteToString(t *template.Template) (string, error) {
	buf := &bytes.Buffer{}
	err := t.Execute(buf, tc)
	if err != nil {
		return "", fmt.Errorf("executing template %q: %w", t.Name(), err)
	}
	return strings.TrimRight(buf.String(), " "), nil
}

// newTemplateContext creates a templateContext from the included schema and
// keyName.  It sets some logical defaults and does slight data transformation.
func newTemplateContext(schema *apiextensions.JSONSchemaProps, keyName string) (*templateContext, error) {
	// remove any newlines from descriptions.  Handling these well will require
	// additional work.
	cleanedDescription := strings.ReplaceAll(schema.Description, "\n", " ")

	tc := &templateContext{
		Name:        keyName,
		KeyName:     keyName,
		Type:        strings.ToLower(schema.Type),
		Description: cleanedDescription,
	}

	if schema.Enum != nil && len(schema.Enum) > 0 {
		for _, e := range schema.Enum {
			switch v := e.(type) {
			case int:
				tc.AllowedValues = append(tc.AllowedValues, strconv.Itoa(v))
			case string:
				tc.AllowedValues = append(tc.AllowedValues, v)
			default:
				return nil, fmt.Errorf("enum value is neither int nor string: %v", v)
			}
		}
	}

	// TODO: This section is a necessary evil for the moment.
	// Guessing the type based on structure is possible, but ultimately we
	// should just be requiring all the types to exist.  V1 CT will get us some
	// of this.  Once the linting work in the associated bug is complete, we
	// may be able to remove this section.
	if tc.Type == "" {
		if schema.Properties != nil {
			tc.Type = "object"
		} else if schema.Items != nil {
			tc.Type = "array"
		} else {
			tc.Type = untyped
		}
	}

	return tc, nil
}

func (tc *templateContext) DelimitedValues() string {
	return strings.Join(tc.AllowedValues, ", ")
}
