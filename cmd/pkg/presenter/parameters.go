package presenter

import (
	"fmt"
	"sort"
	"strings"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
)

const allowedDepth = 80

// parametersText generates the parameters section of a CT validation with the
// desired level of indentation.
//
// Note that the indentation logic at each level of recursion makes this
// logical path n^2, as each line is iterated over at each level of
// indentation.  This is obviously an algorithmic negative, but yields big
// benefits in code complexity.
//
// That said, the "n" is the depth of the parameters section, not the number of
// templates in the library.  As the depth of templates isn't something that's
// going to grow large, we don't have much to worry about.
func parametersText(validation *templates.Validation, indent string) (string, error) {
	if validation.OpenAPIV3Schema == nil {
		return "", nil
	}

	// Check if we're dealing with an empty schema
	blank, err := blankSchema(validation.OpenAPIV3Schema)
	if err != nil {
		return "", err
	}
	if blank {
		return handleBlankSchema(validation)
	}

	lines, err := parametersTextSlice(validation.OpenAPIV3Schema, len(indent))
	if err != nil {
		return "", err
	}

	// indent to the desired level
	prefix(lines, indent)

	return strings.Join(lines, "\n"), nil
}

func prefix(lines []string, prefix string) {
	for i, line := range lines {
		lines[i] = prefix + line
	}
}

// blankSchema returns where the schema has non-nil Items or Properties or a
// non-blank type.
func blankSchema(schema *apiextensions.JSONSchemaProps) (bool, error) {
	if schema == nil {
		return false, fmt.Errorf("expected non-nil schema but got %v", schema)
	}

	return (schema.Items == nil && schema.Properties == nil && schema.Type == ""), nil
}

// handleBlankSchema encompasses special case logic around blank or
// nearly-blank, root level schemas.
//
// In a v1beta1 CT, legacySchema is defaulted to true.  If legacySchema is true
// and the schema is nil,  the conversion logic sets the nil schema to a
// non-nil, empty schema object.  The same logic then adds
// `XPreserveUnknownFields: true` to that non-nil object, as its type is
// unknown.  That confuses us when rendering the schema section, as we have no
// way to differentiate between a v1beta1 CT with no validation section and one
// that contains only `x-kubernetes-preserve-unknown-fields: true`.
//
// On the other hand, v1 CTs default legacySchema to false.  In that case, the
// conversion logic leaves the nil schema unchanged, and no `XPres...` is
// added.  Thus, when legacySchema is false, we _are_ able to differentiate
// between a nil `validation` and a validation with a `XPreserveUnknownFields:
// true` schema.
//
// Consequently, in the case of `legacySchema: true` (the common case for
// v1beta1 CTs), we will assume that the user meant empty if there is nothing
// in the schema besides `XPreserveUnknownFields: true`.  If `legacySchema:
// false`, we will honor the presence of `XPreserveUnknownFields: true`, as it
// was not added by our conversion logic.  We can reliably say that the user
// entered it.
//
// This function also handles the other blank or close-to-blank cases, which
// are dependent on the the composition of these different fields.
func handleBlankSchema(validation *templates.Validation) (string, error) {
	schema := validation.OpenAPIV3Schema
	blank, err := blankSchema(validation.OpenAPIV3Schema)
	if err != nil {
		return "", err
	}
	if !blank {
		return "", fmt.Errorf("wanted blank schema but got: %v", schema)
	}

	// In LegacySchema mode, we can't tell if an empty schema was meant by
	// the user or was machine generated.  So, we assume the CT had an
	// empty validation section, the most common case.
	if *validation.LegacySchema {
		return "", nil
	}

	// Outside of legacySchema mode, we have to handle the case of the
	// user intentionally creating an empty validation with
	// `x-kubernetes-preserve-unknown-fields: true` set.  While it may
	// be ill-advised, it is a valid structural schema.

	// If XPreserveUnknownFields is falsey, we're dealing with a regular empty schema.
	if schema.XPreserveUnknownFields == nil || !*schema.XPreserveUnknownFields {
		return "", nil
	}

	// To usefully a document an empty schema, one must include a
	// description. An empty validation section will look more like an
	// error than an intentionally open-ended set of parameters.
	if schema.Description == "" {
		return "", fmt.Errorf("a non-blank description is required in an unbounded schema section")
	}

	tc, err := newTemplateContext(schema, "unknown fields")
	if err != nil {
		return "", fmt.Errorf("generating template content for empty schema: %w", err)
	}
	line, err := tc.ExecuteToString(descriptionTemp)
	if err != nil {
		return "", fmt.Errorf("writing description for empty schema: %w", err)
	}

	return line, nil
}

func parametersTextSlice(schema *apiextensions.JSONSchemaProps, depth int) ([]string, error) {
	switch {
	case schema == nil:
		return nil, fmt.Errorf("cannot generate text from nil schema")
	case schema.Properties != nil:
		lines, err := handleProperties(schema, depth)
		if err != nil {
			return nil, err
		}
		return lines, nil
	case schema.Items != nil:
		lines, err := handleItems(schema, depth)
		if err != nil {
			return nil, err
		}
		return lines, nil
	case schema.AdditionalProperties != nil:
		lines, err := handleAdditionalProperties(schema, depth)
		if err != nil {
			return nil, err
		}
		return lines, nil
	default:
		t := untyped
		if schema.Type != "" {
			t = schema.Type
		}
		return []string{"<" + t + ">"}, nil
	}
}

// handleProperties generates parameters schema text lines for a schema which
// has `Properties` defined, and is thus of type `object`.  It iterates over
// each key in the Properties map, recursively generating schema text for each
// value within the map.
func handleProperties(schema *apiextensions.JSONSchemaProps, depth int) ([]string, error) {
	var lines []string

	// Sort keys for predictable ordering
	keys := make([]string, 0, len(schema.Properties))
	for k := range schema.Properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		property := schema.Properties[k]

		tc, err := newTemplateContext(&property, k)
		if err != nil {
			return nil, fmt.Errorf("generating template content for key %q: %w", k, err)
		}

		subLines, err := handleProperty(&property, tc, depth)
		if err != nil {
			return nil, fmt.Errorf("for property %q: %w", k, err)
		}
		lines = append(lines, subLines...)
	}

	return lines, nil
}

// handleProperty generates parameters schema text lines for a single value in
// a `Properties` map or for the schema defined in `AdditionalProperties`.
// These elements are essentially the same, other than slight differences in
// the key and comment/description presentations.
func handleProperty(property *apiextensions.JSONSchemaProps, tc *templateContext, depth int) ([]string, error) {
	var lines []string

	commentLines, err := comments(tc, allowedDepth-depth)
	if err != nil {
		return nil, fmt.Errorf("writing comments: %w", err)
	}
	lines = append(lines, commentLines...)

	subLines, err := parametersTextSlice(property, depth+2)
	if err != nil {
		return nil, fmt.Errorf("writing schema: %w", err)
	}

	// If the value is a terminating one, we want to write its contents
	// directly behind the key.
	termination := property.Properties == nil && property.Items == nil && property.AdditionalProperties == nil
	if termination {
		if len(subLines) != 1 {
			return nil, fmt.Errorf("expected sublines length 1, got %v", len(subLines))
		}

		tc.Value = strings.TrimSpace(subLines[0])
	}

	// put the property name (if there is one) as a line, append
	line, err := tc.ExecuteToString(propertyTemp)
	if err != nil {
		return nil, fmt.Errorf("writing property template: %w", err)
	}
	lines = append(lines, line)

	// if we weren't dealing with the termination case, we indent all the sublines and append them
	if !termination {
		prefix(subLines, "  ")
		lines = append(lines, subLines...)
	}

	return lines, nil
}

func handleItems(schema *apiextensions.JSONSchemaProps, depth int) ([]string, error) {
	var lines []string

	tc, err := newTemplateContext(schema.Items.Schema, "")
	if err != nil {
		return nil, fmt.Errorf("generating template content for items: %w", err)
	}

	commentLines, err := comments(tc, allowedDepth-depth)
	if err != nil {
		return nil, fmt.Errorf("writing for item schema: %w", err)
	}
	lines = append(lines, commentLines...)

	subLines, err := parametersTextSlice(schema.Items.Schema, depth+4)
	if err != nil {
		return nil, fmt.Errorf("writing schema for items: %w", err)
	}

	// indent the entire list item
	prefix(subLines, "  ")

	// make the first character of the first line the array character
	subLines[0] = "-" + subLines[0][1:]

	lines = append(lines, subLines...)

	return lines, nil
}

func handleAdditionalProperties(schema *apiextensions.JSONSchemaProps, depth int) ([]string, error) {
	var lines []string

	if !schema.AdditionalProperties.Allows {
		return nil, fmt.Errorf("unable to interpret AdditionalProperties when Allows is false")
	}
	if schema.AdditionalProperties.Schema == nil {
		return nil, fmt.Errorf("unable to interpret AdditionalProperties when schema is nil")
	}

	tc, err := newTemplateContext(schema.AdditionalProperties.Schema, "[key]")
	if err != nil {
		return nil, fmt.Errorf("generating template content for additional properties: %w", err)
	}
	tc.Name = "additional user-defined keys"

	subLines, err := handleProperty(schema.AdditionalProperties.Schema, tc, depth)
	if err != nil {
		return nil, fmt.Errorf("for AdditionalProperties: %w", err)
	}

	return append(lines, subLines...), nil
}

// comments generates the commented information that goes above literal schema
// info at each level.  That includes both a description of the field and the
// values that are allowed for that field.  Both of those elements have the
// potential to be wrapped over multiple lines if they exceed the number of
// characters specified in the lineLength function parameter.
func comments(tc *templateContext, lineLength int) ([]string, error) {
	var lines []string

	if tc.Description != "" {
		line, err := tc.ExecuteToString(descriptionTemp)
		if err != nil {
			return nil, fmt.Errorf("writing description: %w", err)
		}
		wrapped, err := wrapComment(line, lineLength)
		if err != nil {
			return nil, fmt.Errorf("wrapping line %q", line)
		}
		lines = append(lines, wrapped...)
	}

	if len(tc.AllowedValues) > 0 {
		line, err := tc.ExecuteToString(allowedValuesTemplate)
		if err != nil {
			return nil, fmt.Errorf("writing allowedValues: %w", err)
		}
		wrapped, err := wrapComment(line, lineLength)
		if err != nil {
			return nil, fmt.Errorf("wrapping line %q", line)
		}
		lines = append(lines, wrapped...)
	}

	return lines, nil
}

// wrapComment will take a comment string (essentially any string beginning
// with "# " and wrap it to the length specified in the length parameter).
func wrapComment(comment string, length int) ([]string, error) {
	if length < 3 {
		return nil, fmt.Errorf("unwrappable line length %d, wanted 3 or greater", length)
	}
	if !strings.HasPrefix(comment, "# ") {
		return nil, fmt.Errorf("input does not begin with '# ': %q", comment)
	}

	var lines []string

	remaining := comment
	for len(remaining) >= length {
		// find the space before and after the length
		before := -1
		after := -1

		// seek back
		i := length - 1
		for i >= 2 {
			if remaining[i] == ' ' {
				before = i
				break
			}
			i--
		}

		// seek forward
		i = length - 1
		for i < len(remaining) {
			if remaining[i] == ' ' {
				after = i
				break
			}
			i++
		}

		// if we found a space before, let's wrap and get ready for the next
		// iteration.  This is our simplest case.
		if before != -1 {
			line := remaining[:before] // exclusive use of the `before` index
			lines = append(lines, strings.TrimSpace(line))

			// make whatever's left into the next line of comment
			remaining = "# " + remaining[before+1:] // note the inclusive use of the `before` index

			continue
		}

		// if we didn't find a valid space before, then we're dealing with a very
		// long token.  That will get left on its own line

		// if after didn't get set, then there isn't anything to wrap.  We're
		// either at the end or the token is too long.
		if after == -1 {
			break
		}

		// otherwise, we should move to to the end of the long token (leaving
		// it on one line) and wrap afterwards
		line := remaining[:after] // exclusive use of `after`
		lines = append(lines, strings.TrimSpace(line))

		// make whatever's left into the next line of comment
		remaining = "# " + remaining[after+1:] // note the inclusive use of the `after` index
	}

	// by now, we've done all the wrapping we can do.  Let's add whatever is left over.
	return append(lines, strings.TrimSpace(remaining)), nil
}
