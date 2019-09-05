package cat

// Cat represents a cat in the petstore
type Cat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
