package syntax

type Enum struct {
	Name    string
	Options map[string]byte
}

func (v Enum) GetName() string {
	return v.Name
}

func (v Enum) IsBasic() bool {
	return false
}

func (v Enum) IsNamed() bool {
	return true
}

func (v Enum) IsRef() bool {
	return false
}

func (v Enum) IsArray() bool {
	return false
}

func (v Enum) IsSlice() bool {
	return false
}

func (v Enum) IsEnum() bool {
	return true
}

func (v Enum) IsStruct() bool {
	return false
}
