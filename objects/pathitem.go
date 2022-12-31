package objects

type PathItem struct {
	Reference
	Summary     *string
	Description *string
	Get         *Operation
	Put         *Operation
	Post        *Operation
	Delete      *Operation
	Options     *Operation
	Head        *Operation
	Patch       *Operation
	Trace       *Operation
	Servers     *[]Server
	Parameters  *[]Parameter
}
