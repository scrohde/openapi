package objects

type Response struct {
	Reference
	Description string               //	REQUIRED. A description of the response. CommonMark syntax MAY be used for rich text representation.
	Headers     map[string]Header    //	Maps a header name to its definition. [RFC7230] states header names are case insensitive. If a response header is defined with the name "Content-Type", it SHALL be ignored.
	Content     map[string]MediaType // A map containing descriptions of potential response payloads. The key is a media type or media type range and the value describes it. For responses that match multiple keys, only the most specific key is applicable. e.g. text/plain overrides text/*
	Links       map[string]Link      //	A map of operations links that can be followed from the response. The key of the map is a short name for the link, following the naming constraints of the names for Component Objects.
}
