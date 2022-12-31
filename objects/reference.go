package objects

type Reference struct {
	Ref         string `yaml:"$ref" json:"$ref"`
	Summary     *string
	Description *string
}
