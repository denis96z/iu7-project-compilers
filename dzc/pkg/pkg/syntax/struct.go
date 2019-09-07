package syntax

//go:generate easyjson

//easyjson:json
type Struct struct {
	Name  string           `json:"name"`
	Type  string           `json:"type"`
	Attrs map[string]*Attr `json:"attrs"`
}

func NewStruct(name string, attrs map[string]*Attr) *Struct {
	return &Struct{
		Name:  name,
		Type:  TypeStruct,
		Attrs: attrs,
	}
}

func (v Struct) GetName() string {
	return v.Name
}

func (v Struct) GetType() string {
	return v.Type
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

//easyjson:json
type Attr struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}
