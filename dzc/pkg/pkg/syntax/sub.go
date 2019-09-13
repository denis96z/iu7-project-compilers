package syntax

//go:generate easyjson

//easyjson:json
type Procedure struct {
	Name string          `json:"name"`
	Args map[string]*Arg `json:"args"`
}

//easyjson:json
type Function struct {
	Name    string          `json:"name"`
	Args    map[string]*Arg `json:"args"`
	RetType Type            `json:"type"`
}

//easyjson:json
type Arg struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}
