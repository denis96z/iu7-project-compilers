package syntax

//go:generate easyjson

//easyjson:json
type Const struct {
	Name  string      `json:"name"`
	Type  *BasicType  `json:"type"`
	Value interface{} `json:"value"`
}
