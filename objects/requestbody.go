package objects

type RequestBody struct {
	Reference
	Description *string              // Description of the request body. This could contain examples of use. CommonMark syntax MAY be used for rich text representation.
	Content     map[string]MediaType // Content of the request body. The key is a media type or media type range and the value describes it. For requests that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
	Required    *bool                //	Required determines if the request body is required in the request. Defaults to false.
}
