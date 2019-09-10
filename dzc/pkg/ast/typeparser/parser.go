package typeparser

import (
	"fmt"

	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/ast/utils"
	"dzc/pkg/parser"
)

type Parser struct {
	*parser.BaseDZListener

	Types map[string]pkginfo.Type

	consts          map[string]*pkginfo.Const
	incompleteTypes map[string]pkginfo.Type
}

func NewParser(consts map[string]*pkginfo.Const) *Parser {
	return &Parser{
		Types: make(map[string]pkginfo.Type),

		consts:          consts,
		incompleteTypes: make(map[string]pkginfo.Type),
	}
}

func (p *Parser) EnterTypedecl(ctx *parser.TypedeclContext) {
	name := utils.FixTypeName(ctx.GetId().GetText())
	if p.Types[name] != nil {
		panic(fmt.Sprintf("type %q is redeclared", name))
	}

	tName := utils.FixTypeName(ctx.GetT().GetText())
	if tName == name {
		panic(fmt.Sprintf("type %q is self-defining", name))
	}
	if tBase := p.Types[tName]; tBase != nil {
		p.Types[name] =
			&pkginfo.SimpleType{
				Name:     name,
				BaseType: tBase,
			}
		return
	}

	if utils.IsBasicType(tName) {
		p.Types[name] =
			&pkginfo.SimpleType{
				Name:     name,
				BaseType: utils.GetBasicType(tName),
			}
	} else if utils.IsSimpleType(tName) {
		p.addIncompleteType(name, tName)
	} else if utils.IsRefType(tName) {
		bName := utils.ParseBaseNameFromRefTypeName(tName)

		if utils.IsBasicType(bName) {
			p.Types[tName] =
				&pkginfo.RefType{
					Name:     tName,
					BaseType: utils.GetBasicType(bName),
				}
		} else if utils.IsSimpleType(bName) {
			if tBase := p.Types[bName]; tBase == nil {
				p.addIncompleteType(name, tName)
				return
			} else {
				p.Types[tName] =
					&pkginfo.RefType{
						Name:     tName,
						BaseType: tBase,
					}
			}
		}

		p.Types[name] =
			&pkginfo.SimpleType{
				Name:     name,
				BaseType: p.Types[tName],
			}
	} else if utils.IsArrayType(tName) {
		bName := utils.ParseBaseNameFromArrayTypeName(tName)

		var size int

		vSize := utils.ParseSizeFromArrayTypeName(tName)
		switch v := vSize.(type) {
		case int:
			size = v
		case string:
			if c := p.consts[v]; c != nil {
				if c.Type.Text() != pkginfo.BasicTypeSize {
					panic(fmt.Sprintf("const %q is not of type %q in %q", v, pkginfo.BasicTypeSize, tName))
				}
				size = int(c.Value.(uint64))
			} else {
				panic(fmt.Sprintf("const %q in %q is unknown", v, tName))
			}
		}

		if utils.IsBasicType(bName) {
			p.Types[tName] =
				&pkginfo.ArrayType{
					Name:     tName,
					Size:     size,
					BaseType: utils.GetBasicType(bName),
				}
		} else if utils.IsSimpleType(bName) {
			if tBase := p.Types[bName]; tBase == nil {
				p.addIncompleteType(name, tName)
				return
			} else {
				p.Types[tName] =
					&pkginfo.ArrayType{
						Name:     tName,
						Size:     size,
						BaseType: tBase,
					}
			}
		}

		p.Types[name] =
			&pkginfo.SimpleType{
				Name:     name,
				BaseType: p.Types[tName],
			}
	}
}

func (p *Parser) addIncompleteType(name string, baseTypeName string) {
	p.incompleteTypes[name] =
		&incompleteType{
			name:         name,
			baseTypeName: baseTypeName,
		}
}

type incompleteType struct {
	name         string
	baseTypeName string
}

func (t *incompleteType) Text() string {
	return t.name
}

func (t *incompleteType) Base() pkginfo.Type {
	return nil
}

func (p *Parser) FixIncomplete() {
	for name, t := range p.incompleteTypes {
		t := t.(*incompleteType)

		tBase := p.Types[t.baseTypeName]
		if tBase != nil {
			//the type is found
		} else if utils.IsSimpleType(t.baseTypeName) {
			panic(fmt.Sprintf("type %q in %q declaration is unknown", t.baseTypeName, name))
		} else if utils.IsRefType(t.baseTypeName) {
			bName := utils.ParseBaseNameFromRefTypeName(t.baseTypeName)
			if tB := p.Types[bName]; tB != nil {
				tBase = &pkginfo.RefType{
					Name:     t.baseTypeName,
					BaseType: tB,
				}
				p.Types[t.baseTypeName] = tBase
			} else {
				panic(fmt.Sprintf("type %q in %q declaration is unknown", bName, name))
			}
		} else if utils.IsArrayType(t.baseTypeName) {
			bName := utils.ParseBaseNameFromArrayTypeName(t.baseTypeName)

			var size int

			vSize := utils.ParseSizeFromArrayTypeName(t.baseTypeName)
			switch v := vSize.(type) {
			case int:
				size = v
			case string:
				//already checked by now
				size = int(p.consts[v].Value.(uint64)) //TODO ugly - needs to be fixed
			}

			if tB := p.Types[bName]; tB != nil {
				tBase = &pkginfo.ArrayType{
					Name:     t.baseTypeName,
					Size:     size,
					BaseType: tB,
				}
				p.Types[t.baseTypeName] = tBase
			} else {
				panic(fmt.Sprintf("type %q in %q declaration is unknown", bName, name))
			}
		}

		if tBase == nil {
			panic(fmt.Sprintf("no base type for %q", name))
		}

		p.Types[name] = &pkginfo.SimpleType{
			Name:     name,
			BaseType: tBase,
		}
	}
}

func (p *Parser) AddBasic() {
	for k, v := range utils.GetBasicTypes() {
		p.Types[k] = v
	}
}
