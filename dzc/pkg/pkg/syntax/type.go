package syntax

type Type interface {
	GetName() string
	GetType() string
	IsBasic() bool
	IsNamed() bool
	IsRef() bool
	IsArray() bool
	IsSlice() bool
	IsEnum() bool
	IsStruct() bool
}

const (
	TypeBasic  = "basic"
	TypeNamed  = "named"
	TypeRef    = "ref"
	TypeArray  = "array"
	TypeSlice  = "slice"
	TypeEnum   = "enum"
	TypeStruct = "struct"
)
