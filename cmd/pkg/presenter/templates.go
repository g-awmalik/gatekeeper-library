package presenter

import (
	_ "embed"
	"text/template"
)

//go:embed templates/schema.tmpl
var schemaTemplateText string

//go:embed templates/constraint_template.tmpl
var constraintTemplateTemplateText string

const descriptionTemplateText = `
{{- if .Name -}}
# {{.Name}} <{{.Type}}>: {{.Description -}}
{{ else -}}
# <list item: {{.Type}}>: {{.Description -}}
{{ end -}}
`

// This calls a function that joins the values on a comma
const allowedValuesTemplateText = "# Allowed Values: {{.DelimitedValues}}"

//go:embed templates/match.tmpl
var matchTemplateText string

const referentialDataText = `
{{- range $obj := .Objects -}}
---
# Referential Data
{{$obj}}
{{- end -}}
`

var (
	descriptionTemp            *template.Template
	propertyTemp               *template.Template
	schemaTemp                 *template.Template
	constraintTemplateTemplate *template.Template
	allowedValuesTemplate      *template.Template
	matchTemplate              *template.Template
	referentialDataTemplate    *template.Template
)

func init() {
	descriptionTemp = template.Must(template.New("property-description").Parse(descriptionTemplateText))

	propertyTemp = template.Must(template.New("property").Parse("{{.KeyName}}: {{.Value -}}"))

	schemaTemp = template.Must(template.New("schema").Parse(schemaTemplateText))

	constraintTemplateTemplate = template.Must(template.New("constrainttemplate").Parse(constraintTemplateTemplateText))

	allowedValuesTemplate = template.Must(template.New("allowedvalues").Parse(allowedValuesTemplateText))

	matchTemplate = template.Must(template.New("match").Parse(matchTemplateText))

	referentialDataTemplate = template.Must(template.New("referential-data").Parse(referentialDataText))
}
