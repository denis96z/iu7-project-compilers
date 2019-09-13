package syntax

//go:generate easyjson

//easyjson:json
type Enum struct {
	Name    string          `json:"name"`
	Type    string          `json:"type"`
	Options map[string]byte `json:"options"`
}

func NewEnum(name string, options map[string]byte) *Enum {
	return &Enum{
		Name:    name,
		Type:    TypeEnum,
		Options: options,
	}
}

func (v Enum) GetName() string {
	return v.Name
}

func (v Enum) GetType() string {
	return v.Type
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
