package syntax

type Slice struct {
	Name string
	Type Type
}

func (v Slice) GetName() string {
	return v.Name
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
