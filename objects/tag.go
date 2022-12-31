package objects

type Tag struct {
	Name         string                 //REQUIRED. The name of the tag.
	Description  *string                //	A description for the tag. CommonMark syntax MAY be used for rich text representation.
	ExternalDocs *ExternalDocumentation // Object	Additional external documentation for this tag.
}
