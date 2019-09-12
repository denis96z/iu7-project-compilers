package syntax

type Ref struct {
	Name string
	Type Type
}

func (v Ref) GetName() string {
	return v.Name
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
