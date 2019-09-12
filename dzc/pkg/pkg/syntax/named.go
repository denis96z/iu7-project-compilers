package syntax

type NamedType struct {
	Name string
	Type Type
}

func (v NamedType) GetName() string {
	return v.Name
}

func (v NamedType) IsBasic() bool {
	return v.Type.IsBasic()
}

func (v NamedType) IsNamed() bool {
	return true
}

func (v NamedType) IsRef() bool {
	return v.Type.IsRef()
}

func (v NamedType) IsArray() bool {
	return v.Type.IsArray()
}

func (v NamedType) IsSlice() bool {
	return v.Type.IsSlice()
}

func (v NamedType) IsEnum() bool {
	return v.Type.IsEnum()
}

func (v NamedType) IsStruct() bool {
	return v.Type.IsStruct()
}
