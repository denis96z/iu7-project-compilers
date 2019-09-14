package syntax

//easyjson:json
type ConstVal struct {
	Type  *BasicType  `json:"type"`
	Value interface{} `json:"value"`
}

func (v ConstVal) GetValueType() Type {
	return v.Type
}
