package syntax

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

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

const (
	RefSymbol = "@"
)

var (
	basicTypes map[string]Type
)

var PtrSize int

func init() {
	PtrSize = int(unsafe.Sizeof(&PtrSize))
}

func init() {
	bTypes := []string{
		BasicTypeI8, BasicTypeU8,
		BasicTypeI16, BasicTypeU16,
		BasicTypeI32, BasicTypeU32,
		BasicTypeI64, BasicTypeU64,
		BasicTypeSize, BasicTypeSSize,
		BasicTypeChar, BasicTypeBool,
	}

	basicTypes = make(map[string]Type)
	for _, t := range bTypes {
		basicTypes[t] = NewBasicType(t)
	}
}

func GetBasicTypes() map[string]Type {
	return basicTypes
}

func GetBasicType(name string) *BasicType {
	return basicTypes[name].(*BasicType)
}

func IsConst(name string) bool {
	for _, c := range name {
		if (!unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_') || (c != unicode.ToUpper(c)) {
			return false
		}
	}
	return true
}

func IsBasicType(name string) bool {
	return basicTypes[name] != nil
}

var (
	typeNameRegExp = `[a-z]([_]?[a-z0-9])*[_][t]`

	namedTypeRegExp = regexp.MustCompile(`^` + typeNameRegExp + `$`)
	refTypeRegExp   = regexp.MustCompile(`^\` + RefSymbol + typeNameRegExp + `$`)
)

func IsNamedType(name string) bool {
	return namedTypeRegExp.MatchString(name)
}

func IsRefType(name string) bool {
	return refTypeRegExp.MatchString(name)
}

func IsArrayType(name string) bool {
	//TODO check with regexp
	return strings.HasPrefix(name, "[") &&
		strings.Contains(name, ":")
}

func IsSliceType(name string) bool {
	//TODO check with regexp
	return strings.HasPrefix(name, "[") &&
		!strings.Contains(name, ":")
}

func ParseTypeNameFromRefTypeName(name string) string {
	return name[1:]
}

func ParseTypeNameFromArrayTypeName(name string) string {
	return name[1:strings.Index(name, ":")]
}

func ParseSizeFromArrayTypeName(name string) interface{} {
	sizeStr := name[strings.Index(name, ":")+1 : len(name)-1]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		if IsConst(sizeStr) {
			return sizeStr
		}
		log.Fatalf("%q in %q is not a constant name", sizeStr, name)
	}
	return size
}

func ParseTypeNameFromSliceTypeName(name string) string {
	return name[1 : len(name)-1]
}
