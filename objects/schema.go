package objects

type Schema struct {

	// Core vocabulary
	Schema        *string          `yaml:"$schema" json:"$schema"`
	Vocabulary    *map[string]bool `yaml:"$vocabulary" json:"$vocabulary"`
	ID            *string          `yaml:"$id" json:"$id"`
	Ref           *string          `yaml:"$ref" json:"$ref"`
	DynamicRef    *string          `yaml:"$dynamicRef" json:"$dynamicRef"`
	DynamicAnchor *string
	Defs          *Schema `yaml:"$defs" json:"$defs"`
	Comment       *string `yaml:"$comment" json:"$comment"`

	// Keywords for Applying Subschemas With Logic
	AllOf *[]Schema
	AnyOf *[]Schema
	OneOf *[]Schema
	Not   *Schema

	// Keywords for Applying Subschemas Conditionally
	If               *Schema
	Then             *Schema
	Else             *Schema
	DependentSchemas *map[string]Schema

	// Keywords for Applying Subschemas to Arrays
	PrefixItems *[]Schema
	Items       *Schema
	Contains    *Schema

	// Keywords for Applying Subschemas to Objects
	Properties           *map[string]Schema
	PatternProperties    *map[string]Schema
	AdditionalProperties *Schema
	PropertyNames        *Schema

	// Validation keywords for any instance type
	Type  any
	Enum  []any
	Const any

	// Validation Keywords for Numeric Instances (number and integer)
	MultipleOf       int
	Maximum          int
	ExclusiveMaximum int
	Minimum          int
	ExclusiveMinimum int

	// Validation Keywords for Strings
	MaxLength int
	MinLength int
	Pattern   string

	// Validation Keywords for Arrays
	MaxItems    int
	MinItems    int
	UniqueItems int
	MaxContains int
	MinContains int

	// Validation Keywords for Objects
	MaxProperties     int
	MinProperties     int
	Required          []string
	DependentRequired any
}
