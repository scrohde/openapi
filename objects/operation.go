package objects

type Callback map[string]PathItem

type SecurityRequirement map[string][]string

type Operation struct {
	Tags         []string              // A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.
	Summary      string                // A short summary of what the operation does.
	Description  string                // A verbose explanation of the operation behavior. CommonMark syntax MAY be used for rich text representation.
	ExternalDocs ExternalDocumentation // Object Additional external documentation for this operation.
	OperationId  string                // Unique string used to identify the operation. The id MUST be unique among all operations described in the API. The operationId value is case-sensitive. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is RECOMMENDED to follow common programming naming conventions.
	Parameters   []Parameter           // A list of parameters that are applicable for this operation. If a parameter is already defined at the Path Item, the new definition will override it but can never remove it. The list MUST NOT include duplicated parameters. A unique parameter is defined by a combination of a name and location. The list can use the Reference Object to link to parameters that are defined at the OpenAPI Object’s components/parameters.
	RequestBody  RequestBody           // RequestBody applicable for this operation. The RequestBody is fully supported in HTTP methods where the HTTP 1.1 specification [RFC7231] has explicitly defined semantics for request bodies. In other cases where the HTTP spec is vague (such as GET, HEAD and DELETE), requestBody is permitted but does not have well-defined semantics and SHOULD be avoided if possible.
	Responses    map[string]Response   // The list of possible responses as they are returned from executing this operation.
	Callbacks    map[string]Callback   // A map of possible out-of band callbacks related to the parent operation. The key is a unique identifier for the Callback Object. Each value in the map is a Callback Object that describes a request that may be initiated by the API provider and the expected responses.
	Deprecated   bool                  // Declares this operation to be deprecated. Consumers SHOULD refrain from usage of the declared operation. Default value is false.
	Security     []SecurityRequirement // A declaration of which security mechanisms can be used for this operation. The list of values includes alternative security requirement objects that can be used. Only one of the security requirement objects need to be satisfied to authorize a request. To make security optional, an empty security requirement ({}) can be included in the array. This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used.
	Servers      []Server              // An alternative server array to service this operation. If an alternative server object is specified at the Path Item Object or Root level, it will be overridden by this value.
}
