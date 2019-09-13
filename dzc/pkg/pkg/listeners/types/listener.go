package types

import (
	"log"

	"dzc/pkg/parser"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	types           map[string]syntax.Type
	incompleteTypes map[string]syntax.Type

	currentEnumName    string
	currentEnumOptions map[string]byte

	currentStructName       string
	currentStructIsComplete bool
	currentStructAttrs      map[string]*syntax.Attr
}

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkg = info

	v.types = make(map[string]syntax.Type)
	v.incompleteTypes = make(map[string]syntax.Type)

	for tName, t := range syntax.GetBasicTypes() {
		v.types[tName] = t
	}
}

func (v *Listener) EnterTypeDecl(ctx *parser.TypeDeclContext) {
	name := ctx.GetName().GetText()
	if v.types[name] != nil {
		log.Fatalf("type %q is redeclared", name)
	}

	tName := ctx.GetTName().GetText()
	if tName == name {
		log.Fatalf("type %q is self-defining", name)
	}
	if tBase := v.types[tName]; tBase != nil {
		v.types[name] = syntax.NewNamedType(name, tBase)
		return
	}

	if syntax.IsNamedType(tName) {
		v.incompleteTypes[name] =
			&syntax.NamedType{
				Name:     tName,
				BaseType: nil,
			}
	} else if syntax.IsRefType(tName) {
		bt := v.types[syntax.ParseTypeNameFromRefTypeName(tName)]
		if bt == nil {
			v.incompleteTypes[name] =
				&syntax.Ref{
					Name:      tName,
					ValueType: nil,
				}
			return
		}

		if bt.IsBasic() || bt.IsEnum() || bt.IsStruct() {
			v.types[tName] = syntax.NewRef(bt)
			v.types[name] = syntax.NewNamedType(name, v.types[tName])
			return
		}

		log.Fatalf("invalid definition %q for type %q", tName, name)
	} else if syntax.IsArrayType(tName) {
		var size int
		switch vSize := syntax.ParseSizeFromArrayTypeName(tName).(type) {
		case int:
			size = vSize
		case string:
			if c := v.pkg.Consts[vSize]; c != nil {
				if c.Type.GetName() != syntax.BasicTypeSize {
					log.Fatalf(
						"const %q in definition of %q has wrong type %q - expected %q",
						c.Name, name, c.Type.GetName(), syntax.BasicTypeSize,
					)
				}
				size = int(c.Value.(uint64))
			} else {
				log.Fatalf("const %q in definition of %q is unknown", vSize, name)
			}
		}

		bt := v.types[syntax.ParseTypeNameFromArrayTypeName(tName)]
		if bt == nil {
			v.incompleteTypes[name] =
				&syntax.Array{
					Name:     tName,
					ItemType: nil,
					Size:     size,
				}
			return
		}

		if bt.IsRef() || bt.IsArray() || bt.IsSlice() {
			log.Fatalf("invalid definition %q for type %q", tName, name)
		}

		v.types[tName] = syntax.NewArray(bt, size)
		v.types[name] = syntax.NewNamedType(name, v.types[tName])
	} else if syntax.IsSliceType(tName) {
		bt := v.types[syntax.ParseTypeNameFromSliceTypeName(tName)]
		if bt == nil {
			v.incompleteTypes[name] =
				&syntax.Slice{
					Name:     tName,
					ItemType: nil,
				}
			return
		}

		if bt.IsRef() || bt.IsArray() || bt.IsSlice() {
			log.Fatalf("invalid definition %q for type %q", tName, name)
		}

		v.types[tName] = syntax.NewSlice(bt)
		v.types[name] = syntax.NewNamedType(name, v.types[tName])
	}
}

func (v *Listener) EnterEnumDecl(ctx *parser.EnumDeclContext) {
	name := ctx.GetName().GetText()
	if v.types[name] != nil {
		log.Fatalf("type %q is redeclared", name)
	}

	v.currentEnumName = name
	v.currentEnumOptions = make(map[string]byte, 0)
}

func (v *Listener) ExitEnumDecl(ctx *parser.EnumDeclContext) {
	v.types[v.currentEnumName] =
		syntax.NewEnum(v.currentEnumName, v.currentEnumOptions)
}

func (v *Listener) EnterEnumOption(ctx *parser.EnumOptionContext) {
	name := ctx.GetName().GetText()
	if v.currentEnumOptions[name] != 0 {
		log.Fatalf("option %q is redeclared for enum type %q", name, v.currentEnumName)
	}
	v.currentEnumOptions[name] = byte(len(v.currentEnumOptions)) + 1
}

func (v *Listener) EnterStructDecl(ctx *parser.StructDeclContext) {
	name := ctx.GetName().GetText()
	if v.types[name] != nil {
		log.Fatalf("type %q is redeclared", name)
	}

	v.currentStructName = name
	v.currentStructIsComplete = true
	v.currentStructAttrs = make(map[string]*syntax.Attr)
}

func (v *Listener) ExitStructDecl(ctx *parser.StructDeclContext) {
	t := syntax.NewStruct(
		v.currentStructName, v.currentStructAttrs,
	)

	if v.currentStructIsComplete {
		v.types[v.currentStructName] = t
	} else {
		v.incompleteTypes[v.currentStructName] = t
	}
}

func (v *Listener) EnterStructAttr(ctx *parser.StructAttrContext) {
	name := ctx.GetName().GetText()
	if v.currentStructAttrs[name] != nil {
		log.Fatalf("attr %q of type %q is redeclared", name, v.currentStructName)
	}

	tName := ctx.GetTName().GetText()
	if t := v.types[tName]; t != nil {
		if !t.IsBasic() && !t.IsEnum() {
			log.Fatalf("type %q not allowed for attr %q of type %q", tName, name, v.currentStructName)
		}
		v.currentStructAttrs[name] =
			&syntax.Attr{
				Name: name,
				Type: t,
			}
		return
	}

	v.currentStructAttrs[name] =
		&syntax.Attr{
			Name: name,
			Type: nil,
		}
	v.currentStructIsComplete = false
}

func (v *Listener) fixIncomplete() {
	//TODO
}

func (v *Listener) Finalize() {
	v.pkg.Types = v.types
}
