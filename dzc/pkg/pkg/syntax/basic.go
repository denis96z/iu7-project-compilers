package syntax

//go:generate easyjson

//easyjson:json
type BasicType struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewBasicType(name string) *BasicType {
	return &BasicType{
		Name: name,
		Type: TypeBasic,
	}
}

func (v BasicType) GetName() string {
	return v.Name
}

func (v BasicType) GetType() string {
	return v.Type
}

func (v BasicType) IsBasic() bool {
	return true
}

func (v BasicType) IsNamed() bool {
	return true
}

func (v BasicType) IsRef() bool {
	return false
}

func (v BasicType) IsArray() bool {
	return false
}

func (v BasicType) IsSlice() bool {
	return false
}

func (v BasicType) IsEnum() bool {
	return false
}

func (v BasicType) IsStruct() bool {
	return false
}
