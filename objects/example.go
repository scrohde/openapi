package objects

type Example struct {
	Reference
	Summary       string `yaml:"summary" json:"summary,omitempty"`             // Short description for the example.
	Description   string `yaml:"description" json:"description,omitempty"`     // Long description for the example. CommonMark syntax MAY be used for rich text representation.
	Value         any    `yaml:"value" json:"value,omitempty"`                 // Embedded literal example. The value field and externalValue field are mutually exclusive. To represent examples of media types that cannot naturally represented in JSON or YAML, use a string value to contain the example, escaping where necessary.
	ExternalValue string `yaml:"externalValue" json:"externalValue,omitempty"` // A URI that points to the literal example. This provides the capability to reference examples that cannot easily be included in JSON or YAML documents. The value field and externalValue field are mutually exclusive. See the rules for resolving Relative References.
}
