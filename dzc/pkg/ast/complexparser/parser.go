package complexparser

import "fmt"

type Parser struct {
	Enums   map[string]*pkginfo.Enum
	Structs map[string]*pkginfo.Struct

	types map[string]pkginfo.Type

	currentStructName  string
	currentStructAttrs map[string]*pkginfo.Attr

	currentEnumName    string
	currentEnumOptions []string
}

func NewParser(types map[string]pkginfo.Type) *Parser {
	return &Parser{
		types:   types,
		Enums:   make(map[string]*pkginfo.Enum),
		Structs: make(map[string]*pkginfo.Struct),
	}
}

func (p *Parser) EnterEnumdecl(ctx *parser.EnumdeclContext) {
	name := ctx.GetId().GetText()
	if !p.checkTypeIsNew(name) {
		panic(fmt.Sprintf("type %q is redeclared", name))
	}
	p.currentEnumName = name
}

func (p *Parser) ExitEnumdecl(ctx *parser.EnumdeclContext) {
	p.Enums[p.currentEnumName] =
		&pkginfo.Enum{
			Name:    p.currentEnumName,
			Options: p.currentEnumOptions,
		}
}

func (p *Parser) EnterEnumoptions(ctx *parser.EnumoptionsContext) {
	p.currentEnumOptions = make([]string, 0)
}

func (p *Parser) EnterEnumoption(ctx *parser.EnumOptionContext) {
	name := ctx.GetId().GetText()
	for _, o := range p.currentEnumOptions {
		if name == o {
			panic(fmt.Sprintf("option %q is redeclared in enum %q", name, p.currentEnumName))
		}
		p.currentEnumOptions = append(p.currentEnumOptions, name)
	}
}

func (p *Parser) EnterStructdecl(ctx *parser.StructdeclContext) {
	name := ctx.GetId().GetText()
	if !p.checkTypeIsNew(name) {
		panic(fmt.Sprintf("type %q is redeclared", name))
	}
	p.currentStructName = name
}

func (p *Parser) ExistStructdecl(ctx *parser.StructdeclContext) {
	p.Structs[p.currentStructName] =
		&pkginfo.Struct{
			Name:  p.currentStructName,
			Attrs: p.currentStructAttrs,
		}
}

func (p *Parser) EnterStructattrs(ctx *parser.StructattrsContext) {
	p.currentStructAttrs = make(map[string]*pkginfo.Attrs)
}

func (p *Parser) EnterStructattr(ctx *parser.StructattrContext) {
	name := ctx.GetId().GetText()
	if p.currentStructAttrs[name] != nil {
		panic(fmt.Sprintf("attr %q of type %q is redeclared", name, p.currentStructName))
	}

	tName := ctx.GetT().GetText()

	//TODO add support for nested structs
	t := p.types[tName]
	if t == nil {
		panic(fmt.Sprintf("attr %q of type %q has unknown type %q", name, p.currentStructName, tName))
	}

	p.currentStructAttrs[name] =
		&pkginfo.Attr{
			Name: name,
			Type: t,
		}
}

func (p *Parser) checkTypeIsNew(name string) {
	return p.Enums[name] == nil && p.Structs[name] == nil && p.types[name] == nil
}
