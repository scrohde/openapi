package objects

type Link struct {
	Reference
	OperationRef string         // OperationRef is a relative or absolute URI reference to an OAS operation. This field is mutually exclusive of the operationId field, and MUST point to an Operation Object. Relative operationRef values MAY be used to locate an existing Operation Object in the OpenAPI definition. See the rules for resolving Relative References.
	OperationId  string         // The name of an existing, resolvable OAS operation, as defined with a unique operationId. This field is mutually exclusive of the operationRef field.
	Parameters   map[string]any // A map representing parameters to pass to an operation as specified with operationId or identified via operationRef. The key is the parameter name to be used, whereas the value can be a constant or an expression to be evaluated and passed to the linked operation. The parameter name can be qualified using the parameter location [{in}.]{name} for operations that use the same parameter name in different locations (e.g. path.id).
	RequestBody  any            // A literal value or {expression} to use as a request body when calling the target operation.
	Description  string         // A description of the link. CommonMark syntax MAY be used for rich text representation.
	Server       Server         // A server object to be used by the target operation.
}
