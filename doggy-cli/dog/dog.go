package dog

// Dog represents a dog in the petstore
type Dog struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
