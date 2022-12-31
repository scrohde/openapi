package objects

type Components struct {
	Schemas         *map[string]Schema         // An object to hold reusable Schema Objects.
	Responses       *map[string]Response       // An object to hold reusable Response Objects.
	Parameters      *map[string]Parameter      // An object to hold reusable Parameter Objects.
	Examples        *map[string]Example        // An object to hold reusable Example Objects.
	RequestBodies   *map[string]RequestBody    // An object to hold reusable Request Body Objects.
	Headers         *map[string]Header         // An object to hold reusable Header Objects.
	SecuritySchemes *map[string]SecurityScheme // An object to hold reusable Security Scheme Objects.
	Links           *map[string]Link           // An object to hold reusable Link Objects.
	Callbacks       *map[string]Callback       // An object to hold reusable Callback Objects.
	PathItems       *map[string]PathItem       // An object to hold reusable Path Item Object.
}
