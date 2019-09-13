package syntax

//easyjson:json
type Expr struct {
	Type  string     `json:"type"`
	Value Expression `json:"value"`
}

type Expression interface {
	GetType()
}
