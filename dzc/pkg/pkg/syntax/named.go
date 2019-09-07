package syntax

//go:generate easyjson

//easyjson:json
type NamedType struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	BaseType Type   `json:"base_type"`
}

func NewNamedType(name string, baseType Type) *NamedType {
	return &NamedType{
		Name:     name,
		Type:     TypeNamed,
		BaseType: baseType,
	}
}

func (v NamedType) GetName() string {
	return v.Name
}

func (v NamedType) GetType() string {
	return v.Type
}

func (v NamedType) GetBaseType() Type {
	t := v.BaseType
	for {
		if !t.IsNamed() {
			break
		}
		t = (t.(*NamedType)).BaseType
	}
	return t
}

func (v NamedType) IsBasic() bool {
	return v.BaseType.IsBasic()
}

func (v NamedType) IsNamed() bool {
	return true
}

func (v NamedType) IsRef() bool {
	return v.BaseType.IsRef()
}

func (v NamedType) IsArray() bool {
	return v.BaseType.IsArray()
}

func (v NamedType) IsSlice() bool {
	return v.BaseType.IsSlice()
}

func (v NamedType) IsEnum() bool {
	return v.BaseType.IsEnum()
}

func (v NamedType) IsStruct() bool {
	return v.BaseType.IsStruct()
}
