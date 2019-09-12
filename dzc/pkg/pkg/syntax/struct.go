package syntax

type Struct struct {
	Name  string
	Attrs map[string]*Attr
}

func (v Struct) GetName() string {
	return v.Name
}

func (v Struct) IsBasic() bool {
	return false
}

func (v Struct) IsNamed() bool {
	return true
}

func (v Struct) IsRef() bool {
	return false
}

func (v Struct) IsArray() bool {
	return false
}

func (v Struct) IsSlice() bool {
	return false
}

func (v Struct) IsEnum() bool {
	return false
}

func (v Struct) IsStruct() bool {
	return true
}

type Attr struct {
	Name string
	Type Type
}
