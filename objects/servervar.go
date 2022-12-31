package objects

type ServerVariable struct {
	// Enum of string values to be used if the substitution options are from a limited set. The array MUST NOT be empty.
	Enum []string
	// Default value to use for substitution, which SHALL be sent if an alternate value is not supplied. Note this behavior is different than the Schema Object’s treatment of default values, because in those cases parameter values are optional. If the Enum is defined, the value MUST exist in the Enum’s values.
	Default string
	// Description for the server variable. CommonMark syntax MAY be used for rich text representation.
	Description string
}
