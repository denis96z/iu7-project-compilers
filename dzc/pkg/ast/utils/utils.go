package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"dzc/pkg/ast/pkginfo"
)

var (
	basicTypes map[string]pkginfo.Type
)

func init() {
	bTypes := []string{
		pkginfo.BasicTypeI8, pkginfo.BasicTypeU8,
		pkginfo.BasicTypeI16, pkginfo.BasicTypeU16,
		pkginfo.BasicTypeI32, pkginfo.BasicTypeU32,
		pkginfo.BasicTypeI64, pkginfo.BasicTypeU64,
		pkginfo.BasicTypeSize, pkginfo.BasicTypeSSize,
		pkginfo.BasicTypeChar, pkginfo.BasicTypeBool,
	}

	basicTypes = make(map[string]pkginfo.Type)
	for _, t := range bTypes {
		name := FixTypeName(t)
		basicTypes[name] = &pkginfo.BasicType{
			Name: t,
		}
	}
}

func GetBasicTypes() map[string]pkginfo.Type {
	return basicTypes //TODO return copy
}

func GetBasicType(name string) pkginfo.Type {
	return basicTypes[FixTypeName(name)]
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
	return basicTypes[FixTypeName(name)] != nil
}

func IsSimpleType(name string) bool {
	return !strings.HasPrefix(name, pkginfo.RefSymbol) && //TODO improve
		!strings.HasPrefix(name, "[") //TODO const?
}

func IsRefType(name string) bool {
	return strings.HasPrefix(name, pkginfo.RefSymbol)
}

func IsArrayType(name string) bool {
	return strings.HasPrefix(name, "[")
}

func FixPkgName(name string) string {
	return name //TODO ?
}

func FixTypeName(name string) string {
	return name //TODO ?
}

func ParseBaseNameFromRefTypeName(name string) string {
	name = FixTypeName(name)
	return name[1:]
}

func ParseBaseNameFromArrayTypeName(name string) string {
	name = FixTypeName(name)
	return name[1:strings.Index(name, ":")]
}

func ParseSizeFromArrayTypeName(name string) interface{} {
	name = FixTypeName(name)

	sizeStr := name[strings.Index(name, ":")+1 : len(name)-1]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		// TODO improve check
		for _, r := range sizeStr {
			if r != unicode.ToUpper(r) {
				panic(fmt.Sprintf("%q is not a constant", sizeStr)) //TODO handle gracefully
			}
		}
		return sizeStr //XXX constant name is returned
	}

	return size
}
