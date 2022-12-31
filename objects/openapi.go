package objects

type OpenAPI struct {
	// OpenAPI MUST be the version number of the OpenAPI Specification that the OpenAPI document uses. The openapi field SHOULD be used by tooling to interpret the OpenAPI document. This is not related to the API Info.Version string.
	OpenAPI string `yaml:"openapi,omitempty" json:"openapi,omitempty"`

	// Info provides metadata about the API. The metadata MAY be used by tooling as required.
	Info              Info
	JsonSchemaDialect *string
	Servers           *[]Server
	Paths             *map[string]PathItem
	Webhooks          *map[string]PathItem
	Components        *Components
	Tags              []Tag
}
