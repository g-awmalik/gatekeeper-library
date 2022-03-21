package presenter

import (
	"bytes"
	"fmt"
	"gatekeeperlibrary/cmd/pkg/apis"
)

type sample struct {
	Constraint     string
	ConstraintName string
	Allowed        []string
	Disallowed     []string
}

func marshalSamples(structureds []*apis.Sample) ([]*sample, error) {
	marshalled := make([]*sample, 0, len(structureds))

	for _, structured := range structureds {
		constraintName := structured.Constraint.Name()
		constraint, err := structured.Constraint.ObjectString()
		if err != nil {
			return nil, fmt.Errorf("converting Constraint '%v' to string: %w", constraintName, err)
		}

		s := sample{
			Constraint:     constraint,
			ConstraintName: constraintName,
			Allowed:        make([]string, 0, len(structured.Allowed)),
			Disallowed:     make([]string, 0, len(structured.Disallowed)),
		}

		for _, allowed := range structured.Allowed {
			obj, err := allowed.ObjectString()
			if err != nil {
				return nil, fmt.Errorf("converting allowed example '%v' for contraint '%v' to string: %w", allowed.Name(), constraintName, err)
			}

			refData, err := refDataFormatted(allowed)
			if err != nil {
				return nil, fmt.Errorf("formatting referential data for allowed example '%v' for contraint '%v': %w", allowed.Name(), constraintName, err)
			}

			s.Allowed = append(s.Allowed, obj+refData)
		}
		for _, disallowed := range structured.Disallowed {
			obj, err := disallowed.ObjectString()
			if err != nil {
				return nil, fmt.Errorf("converting disallowed example '%v' for contraint '%v' to string: %w", disallowed.Name(), constraintName, err)
			}

			refData, err := refDataFormatted(disallowed)
			if err != nil {
				return nil, fmt.Errorf("formatting referential data for disallowed example '%v' for contraint '%v': %w", disallowed.Name(), constraintName, err)
			}

			s.Disallowed = append(s.Disallowed, obj+refData)
		}

		marshalled = append(marshalled, &s)
	}

	return marshalled, nil
}

func refDataFormatted(owc *apis.ObjectWithContext) (string, error) {
	buf := &bytes.Buffer{}

	refDatas, err := owc.ReferentialDataStrings()
	if err != nil {
		return "", fmt.Errorf("converting referential data to string: %w", err)
	}

	err = referentialDataTemplate.Execute(buf, struct {
		Objects []string
	}{
		Objects: refDatas,
	})
	if err != nil {
		return "", fmt.Errorf("executing referential data template: %w", err)
	}

	return buf.String(), nil
}
