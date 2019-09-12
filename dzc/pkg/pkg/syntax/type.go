package syntax

type Type interface {
	GetName() string
	IsBasic() bool
	IsNamed() bool
	IsRef() bool
	IsArray() bool
	IsSlice() bool
	IsEnum() bool
	IsStruct() bool
}
