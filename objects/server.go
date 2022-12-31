package objects

type Server struct {
	URL         string                     `yaml:"url" json:"url,omitempty"`
	Description *string                    `yaml:"description" json:"description,omitempty"`
	Variables   *map[string]ServerVariable `yaml:"variables" json:"variables,omitempty"`
}
