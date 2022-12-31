package objects

type Info struct {
	Title          string `yaml:"title" json:"title"`
	Summary        string `yaml:"summary" json:"summary"`
	Description    string
	TermsOfService string `yaml:"termsOfService" json:"termsOfService"`
	Contact        Contact
	License        License
	Version        string
}
