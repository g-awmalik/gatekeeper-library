package presenter

import (
	"bytes"
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"
	"strings"

	"github.com/open-policy-agent/gatekeeper/pkg/target"
)

// Present accepts a slice of ConstraintTemplateDoc objects and generates
// markdown text output, returned as a string.
func Present(docs []*apis.ConstraintTemplateDoc) (string, error) {
	buf := &bytes.Buffer{}
	for _, doc := range docs {
		sch, err := schemaText(doc)
		if err != nil {
			return "", fmt.Errorf("converting schema of doc %q to text: %w", doc.Name, err)
		}

		samples, err := marshalSamples(doc.Samples)
		if err != nil {
			return "", fmt.Errorf("parsing samples for doc '%s': %w", doc.Name, err)
		}

		err = constraintTemplateTemplate.Execute(buf, struct {
			Name        string
			Schema      string
			Description string
			Samples     []*sample
		}{
			Name:        doc.Name,
			Schema:      sch,
			Description: doc.Description,
			Samples:     samples,
		})
		if err != nil {
			return "", fmt.Errorf("executing template %q: %w", doc.Name, err)
		}
	}

	return strings.TrimSpace(buf.String()), nil
}

// Match returns the generated portion of the Policy Controller `match`
// section.  This is derived from code in the gatekeeper repository.
func Match() (string, error) {
	buf := &bytes.Buffer{}

	k := &target.K8sValidationTarget{}
	ms := k.MatchSchema()

	lines, err := parametersTextSlice(&ms, 0)
	if err != nil {
		return "", fmt.Errorf("generating lines for Match Schema: %w", err)
	}

	err = matchTemplate.Execute(buf, struct {
		Content string
	}{
		Content: strings.Join(lines, "\n"),
	})
	if err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return strings.TrimSpace(buf.String()), nil
}

func schemaText(doc *apis.ConstraintTemplateDoc) (string, error) {
	pt, err := parametersText(doc.Validation, "    ")
	if err != nil {
		return "", fmt.Errorf("generating parametersText: %w", err)
	}

	buf := &bytes.Buffer{}
	err = schemaTemp.Execute(buf, struct {
		Parameters string
		Kind       string
	}{
		Parameters: pt,
		Kind:       doc.Name,
	})
	if err != nil {
		return "", fmt.Errorf("executing schema template: %w", err)
	}

	return buf.String(), nil
}
