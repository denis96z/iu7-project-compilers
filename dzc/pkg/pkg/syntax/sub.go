package syntax

//go:generate easyjson

//easyjson:json
type Proc struct {
	Name string          `json:"name"`
	Args map[string]*Arg `json:"args"`
	Body *Block          `json:"body"`
}

//easyjson:json
type Func struct {
	Name    string          `json:"name"`
	Args    map[string]*Arg `json:"args"`
	Body    *Block          `json:"body"`
	RetType Type            `json:"type"`
}

//easyjson:json
type Arg struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}
