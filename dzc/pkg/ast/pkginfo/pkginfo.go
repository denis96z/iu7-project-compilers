package pkginfo

import (
	"unsafe"
)

var PtrSize int

func init() {
	var v int
	PtrSize = int(unsafe.Sizeof(&v))
}

type PkgInfo struct {
	Pkg        *Pkg
	Types      map[string]Type
	Procedures map[string]*Procedure
	Functions  map[string]*Function
}

type Pkg struct {
	Name string //TODO add imports
}

type Const struct {
	Name  string
	Type  Type
	Value interface{}
}

type Type interface {
	Text() string
	Base() Type
}

const (
	BoolFalse = "false"
	BoolTrue  = "true"
)

const (
	BasicTypeI8    = "i8_t"
	BasicTypeU8    = "u8_t"
	BasicTypeI16   = "i16_t"
	BasicTypeU16   = "u16_t"
	BasicTypeI32   = "i32_t"
	BasicTypeU32   = "u32_t"
	BasicTypeI64   = "i64_t"
	BasicTypeU64   = "u64_t"
	BasicTypeSize  = "size_t"
	BasicTypeSSize = "ssize_t"
	BasicTypeBool  = "bool_t"
	BasicTypeChar  = "char_t"
)

type BasicType struct {
	Name string
}

func (t BasicType) Text() string {
	return t.Name
}

func (t BasicType) Base() Type {
	return nil
}

type SimpleType struct {
	Name     string
	BaseType Type
}

func (t SimpleType) Text() string {
	return t.Name
}

func (t SimpleType) Base() Type {
	return t.BaseType
}

type RefType struct {
	Name     string
	BaseType Type
}

const (
	RefSymbol = "@"
)

func (t RefType) Text() string {
	return t.Name
}

func (t RefType) Base() Type {
	return t.BaseType
}

type ArrayType struct {
	Name     string
	Size     int
	BaseType Type
}

func (t ArrayType) Text() string {
	return t.Name
}

func (t ArrayType) Base() Type {
	return t.BaseType
}

type Procedure struct {
	Name string
	Args map[string]*Arg
}

type Function struct {
	Name    string
	Args    map[string]*Arg
	RetType Type
}

type Arg struct {
	Name string
	Type Type
}
