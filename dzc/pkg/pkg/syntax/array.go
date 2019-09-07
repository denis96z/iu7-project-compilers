package syntax

//go:generate easyjson

import (
	"fmt"
)

//easyjson:json
type Array struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	ItemType Type   `json:"item_type"`
	Size     int    `json:"size"`
}

func NewArray(itemType Type, size int) *Array {
	return &Array{
		Name:     fmt.Sprintf("[%s:%d]", itemType.GetName(), size),
		Type:     TypeArray,
		ItemType: itemType,
		Size:     size,
	}
}

func (v Array) GetName() string {
	return v.Name
}

func (v Array) GetType() string {
	return v.Type
}

func (v Array) IsBasic() bool {
	return false
}

func (v Array) IsNamed() bool {
	return false
}

func (v Array) IsRef() bool {
	return false
}

func (v Array) IsArray() bool {
	return true
}

func (v Array) IsSlice() bool {
	return true //XXX so as to be passed as arg
}

func (v Array) IsEnum() bool {
	return false
}

func (v Array) IsStruct() bool {
	return false
}
