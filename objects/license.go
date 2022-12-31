package objects

type License struct {
	Name       string `json:"name,omitempty" yaml:"name"`
	Identifier string `json:"identifier,omitempty" yaml:"identifier"`
	URL        string `json:"url,omitempty" yaml:"url"`
}
