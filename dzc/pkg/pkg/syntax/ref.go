package syntax

//go:generate easyjson

import (
	"fmt"
)

//easyjson:json
type Ref struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	ValueType Type   `json:"value_type"`
}

func NewRef(valueType Type) *Ref {
	return &Ref{
		Name:      fmt.Sprintf("%s%s", RefSymbol, valueType.GetName()),
		Type:      TypeRef,
		ValueType: valueType,
	}
}

func (v Ref) GetName() string {
	return v.Name
}

func (v Ref) GetType() string {
	return v.Type
}

func (v Ref) IsBasic() bool {
	return false
}

func (v Ref) IsNamed() bool {
	return false
}

func (v Ref) IsRef() bool {
	return true
}

func (v Ref) IsArray() bool {
	return false
}

func (v Ref) IsSlice() bool {
	return false
}

func (v Ref) IsEnum() bool {
	return false
}

func (v Ref) IsStruct() bool {
	return false
}
