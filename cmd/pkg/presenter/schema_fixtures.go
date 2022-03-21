package presenter

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/utils/pointer"
)

var nilSchema *apiextensions.JSONSchemaProps

var empty = &apiextensions.JSONSchemaProps{}

var emptyWithXPreserve = &apiextensions.JSONSchemaProps{XPreserveUnknownFields: pointer.BoolPtr(true)}

const emptyWithXPreserveText = `<untyped>`

var descriptiveEmptyWithXPreserve = &apiextensions.JSONSchemaProps{
	XPreserveUnknownFields: pointer.BoolPtr(true),
	Description:            "pepperoni pizza",
}

const descriptiveEmptyWithXPreserveText = `# unknown fields <untyped>: pepperoni pizza`

var singlePrimitive = &apiextensions.JSONSchemaProps{
	Type: "string",
}

const singlePrimitiveText = `<string>`

var mapWithEmptyVal = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"foo": {
			XPreserveUnknownFields: pointer.BoolPtr(true),
		},
		"bar": {
			Type: "boolean",
		},
	},
}

const mapWithEmptyValText = `bar: <boolean>
foo: <untyped>`

var primitiveTypeMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"foo": {
			Type: "string",
		},
		"bar": {
			Type: "boolean",
		},
		"taco": {
			Type: "number",
		},
	},
}

const primitiveTypeMapText = `bar: <boolean>
foo: <string>
taco: <number>`

var descriptivePrimitiveTypeMap = &apiextensions.JSONSchemaProps{
	Description: "refried beans",
	Properties: map[string]apiextensions.JSONSchemaProps{
		"foo": {
			Type:        "string",
			Description: "foo is a string",
		},
		"bar": {
			Type:        "boolean",
			Description: "pizza pie",
		},
		"taco": {
			Type:        "number",
			Description: "big pink pumpkin",
		},
	},
}

const descriptivePrimitiveTypeMapText = `# bar <boolean>: pizza pie
bar: <boolean>
# foo <string>: foo is a string
foo: <string>
# taco <number>: big pink pumpkin
taco: <number>`

var enumPrimitiveTypeMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"foo": {
			Type: "string",
			Enum: []apiextensions.JSON{"bar", "pizza"},
		},
		"bar": {
			Type: "boolean",
		},
		"taco": {
			Type: "number",
		},
	},
}

const enumPrimitiveTypeMapText = `bar: <boolean>
# Allowed Values: bar, pizza
foo: <string>
taco: <number>`

var descriptiveEnumPrimitiveTypeMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"foo": {
			Type:        "string",
			Description: "foo is a string",
			Enum:        []apiextensions.JSON{"bar", "pizza"},
		},
		"bar": {
			Type:        "boolean",
			Description: "pizza pie",
		},
		"taco": {
			Type:        "number",
			Description: "big pink pumpkin",
		},
	},
}

const descriptiveEnumPrimitiveTypeMapText = `# bar <boolean>: pizza pie
bar: <boolean>
# foo <string>: foo is a string
# Allowed Values: bar, pizza
foo: <string>
# taco <number>: big pink pumpkin
taco: <number>`

var primitiveTypeArray = &apiextensions.JSONSchemaProps{
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: &apiextensions.JSONSchemaProps{
			Type: "string",
		},
	},
}

const primitiveTypeArrayText = `- <string>`

var descriptivePrimitiveTypeArray = &apiextensions.JSONSchemaProps{
	Description: "a list of tofus",
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: &apiextensions.JSONSchemaProps{
			Type:        "string",
			Description: "deep fried tofu",
		},
	},
}

const descriptivePrimitiveTypeArrayText = `# <list item: string>: deep fried tofu
- <string>`

var enumPrimitiveTypeArray = &apiextensions.JSONSchemaProps{
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: &apiextensions.JSONSchemaProps{
			Type: "string",
			Enum: []apiextensions.JSON{"bar", "pizza"},
		},
	},
}

const enumPrimitiveTypeArrayText = `# Allowed Values: bar, pizza
- <string>`

var descriptiveEnumPrimitiveTypeArray = &apiextensions.JSONSchemaProps{
	Description: "a list of tofus",
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: &apiextensions.JSONSchemaProps{
			Type:        "string",
			Description: "deep fried tofu",
			Enum:        []apiextensions.JSON{"bar", "pizza"},
		},
	},
}

const descriptiveEnumPrimitiveTypeArrayText = `# <list item: string>: deep fried tofu
# Allowed Values: bar, pizza
- <string>`

var arrayWithMap = &apiextensions.JSONSchemaProps{
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: primitiveTypeMap,
	},
}

const arrayWithMapText = `- bar: <boolean>
  foo: <string>
  taco: <number>`

var descriptiveArrayWithMap = &apiextensions.JSONSchemaProps{
	Items: &apiextensions.JSONSchemaPropsOrArray{
		Schema: descriptivePrimitiveTypeMap,
	},
}

const descriptiveArrayWithMapText = `# <list item: object>: refried beans
- # bar <boolean>: pizza pie
  bar: <boolean>
  # foo <string>: foo is a string
  foo: <string>
  # taco <number>: big pink pumpkin
  taco: <number>`

var nestedMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type: "boolean",
		},
		"mappy": *primitiveTypeMap,
	},
}

const nestedMapText = `burrito: <boolean>
mappy:
  bar: <boolean>
  foo: <string>
  taco: <number>`

var descriptiveNestedMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type: "boolean",
		},
		"mappy": *descriptivePrimitiveTypeMap,
	},
}

const descriptiveNestedMapText = `burrito: <boolean>
# mappy <object>: refried beans
mappy:
  # bar <boolean>: pizza pie
  bar: <boolean>
  # foo <string>: foo is a string
  foo: <string>
  # taco <number>: big pink pumpkin
  taco: <number>`

var primitiveArrayWithinMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type: "boolean",
		},
		"listy": *primitiveTypeArray,
	},
}

const primitiveArrayWithinMapText = `burrito: <boolean>
listy:
  - <string>`

var descriptivePrimitiveArrayWithinMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type: "boolean",
		},
		"listy": *descriptivePrimitiveTypeArray,
	},
}

const descriptivePrimitiveArrayWithinMapText = `burrito: <boolean>
# listy <array>: a list of tofus
listy:
  # <list item: string>: deep fried tofu
  - <string>`

var mapSchemaArrayWithinMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type: "boolean",
		},
		"listy": *arrayWithMap,
	},
}

const mapSchemaArrayWithinMapText = `burrito: <boolean>
listy:
  - bar: <boolean>
    foo: <string>
    taco: <number>`

var descriptiveMapSchemaArrayWithinMap = &apiextensions.JSONSchemaProps{
	Properties: map[string]apiextensions.JSONSchemaProps{
		"burrito": {
			Type:        "boolean",
			Description: "carne asada",
		},
		"listy": *descriptiveArrayWithMap,
	},
}

const descriptiveMapSchemaArrayWithinMapText = `# burrito <boolean>: carne asada
burrito: <boolean>
listy:
  # <list item: object>: refried beans
  - # bar <boolean>: pizza pie
    bar: <boolean>
    # foo <string>: foo is a string
    foo: <string>
    # taco <number>: big pink pumpkin
    taco: <number>`

var additionalPropertiesPrimitive = &apiextensions.JSONSchemaProps{
	AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
		Allows: true,
		Schema: &apiextensions.JSONSchemaProps{
			Type: "string",
		},
	},
}

const additionalPropertiesPrimitiveText = `[key]: <string>`

var descriptiveAdditionalPropertiesPrimitive = &apiextensions.JSONSchemaProps{
	AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
		Allows: true,
		Schema: &apiextensions.JSONSchemaProps{
			Description: "a delicious taco",
			Type:        "string",
		},
	},
}

const descriptiveAdditionalPropertiesPrimitiveText = `# additional user-defined keys <string>: a delicious taco
[key]: <string>`

var additionalPropertiesMap = &apiextensions.JSONSchemaProps{
	AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
		Allows: true,
		Schema: &apiextensions.JSONSchemaProps{
			Type: "object",
			Properties: map[string]apiextensions.JSONSchemaProps{
				"foo": {
					Type: "string",
				},
			},
		},
	},
}

const additionalPropertiesMapText = `[key]:
  foo: <string>`

var descriptiveAdditionalPropertiesMap = &apiextensions.JSONSchemaProps{
	Description: "An arbitrary set of keys",
	AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
		Allows: true,
		Schema: &apiextensions.JSONSchemaProps{
			Type:        "object",
			Description: "A defined set of keys",
			Properties: map[string]apiextensions.JSONSchemaProps{
				"foo": {
					Description: "A specific string",
					Type:        "string",
				},
			},
		},
	},
}

const descriptiveAdditionalPropertiesMapText = `# additional user-defined keys <object>: A defined set of keys
[key]:
  # foo <string>: A specific string
  foo: <string>`
