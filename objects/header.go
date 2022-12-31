package objects

type Header struct {
	Reference
	Description     *string
	Required        *bool // Required determines whether this parameter is mandatory. If the parameter location is "path", this property is REQUIRED and its value MUST be true. Otherwise, the property MAY be included and its default value is false.
	Deprecated      *bool // Deprecated specifies that a parameter is deprecated and SHOULD be transitioned out of usage. Default value is false.
	AllowEmptyValue *bool // AllowEmptyValue sets the ability to pass empty-valued parameters. This is valid only for query parameters and allows sending a parameter with an empty value. Default value is false. If style is used, and if behavior is n/a (cannot be serialized), the value of allowEmptyValue SHALL be ignored. Use of this property is NOT RECOMMENDED, as it is likely to be removed in a later revision.
}
