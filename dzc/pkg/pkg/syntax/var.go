package syntax

//easyjson:json
type Var struct {
	Name  string     `json:"name"`
	Type  Type       `json:"type"`
	Value Expression `json:"value"`
}

func (v Var) GetValueType() Type {
	return v.Type
}
