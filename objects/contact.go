package objects

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name"`
	URL   string `json:"url,omitempty" yaml:"url"`
	Email string `json:"email,omitempty" yaml:"email"`
}
