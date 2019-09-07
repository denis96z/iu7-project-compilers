package syntax

//go:generate easyjson

import (
	"fmt"
)

//easyjson:json
type Slice struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	ItemType Type   `json:"item_type"`
}

func NewSlice(itemType Type) *Slice {
	return &Slice{
		Name:     fmt.Sprintf("[%s]", itemType.GetName()),
		Type:     TypeSlice,
		ItemType: itemType,
	}
}

func (v Slice) GetName() string {
	return v.Name
}

func (v Slice) GetType() string {
	return v.Type
}

func (v Slice) IsBasic() bool {
	return false
}

func (v Slice) IsNamed() bool {
	return false
}

func (v Slice) IsRef() bool {
	return false
}

func (v Slice) IsArray() bool {
	return false
}

func (v Slice) IsSlice() bool {
	return true
}

func (v Slice) IsEnum() bool {
	return false
}

func (v Slice) IsStruct() bool {
	return false
}
