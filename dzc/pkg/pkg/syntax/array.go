package syntax

type Array struct {
	Name string
	Type Type
	Size int
}

func (v Array) GetName() string {
	return v.Name
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
