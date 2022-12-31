package objects

type OAuthFlow struct {
	AuthorizationUrl string            // oauth2 ("implicit", "authorizationCode")	REQUIRED. The authorization URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	TokenUrl         string            // oauth2 ("password", "clientCredentials", "authorizationCode")	REQUIRED. The token URL to be used for this flow. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	RefreshUrl       string            // oauth2 The URL to be used for obtaining refresh tokens. This MUST be in the form of a URL. The OAuth2 standard requires the use of TLS.
	Scopes           map[string]string // oauth2 REQUIRED. The available scopes for the OAuth2 security scheme. A map between the scope name and a short description for it. The map MAY be empty.
}

type OAuthFlows struct {
	Implicit          OAuthFlow //Object Configuration for the OAuth Implicit flow
	Password          OAuthFlow //Object Configuration for the OAuth Resource Owner Password flow
	ClientCredentials OAuthFlow // Object Configuration for the OAuth Client Credentials flow. Previously called application in OpenAPI 2.0.
	AuthorizationCode OAuthFlow // Object Configuration for the OAuth Authorization Code flow. Previously called accessCode in OpenAPI 2.0.
}

type SecurityScheme struct {
	Reference
	// Type of the security scheme. Valid values are "apiKey", "http", "mutualTLS", "oauth2", "openIdConnect".
	Type string
	// Description for security scheme. CommonMark syntax MAY be used for rich text representation.
	Description string
	// Name of the header, query or cookie parameter to be used.
	Name string
	// In is the location of the API key. Valid values are "query", "header" or "cookie".
	In string
	// Scheme is the name of the HTTP Authorization scheme to be used in the Authorization header as defined in [RFC7235]. The values used SHOULD be registered in the IANA Authentication Scheme registry.
	Scheme string
	// BearerFormat is a hint to the client to identify how the bearer token is formatted. Bearer tokens are usually generated by an authorization server, so this information is primarily for documentation purposes.
	BearerFormat string
	// Flows is an object containing configuration information for the flow types supported.
	Flows OAuthFlows
	// OpenIdConnectURL to discover OAuth2 configuration values. This MUST be in the form of a URL. The OpenID Connect standard requires the use of TLS.
	OpenIdConnectURL string `yaml:"openIdConnectUrl" json:"openIdConnectUrl"`
}
