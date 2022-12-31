package objects

type ExternalDocumentation struct {
	// Description of the target documentation. CommonMark syntax MAY be used for rich text representation.
	Description *string
	// URL for the target documentation. This MUST be in the form of a URL.
	URL string //REQUIRED.
}
