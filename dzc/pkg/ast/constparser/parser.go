package constparser

import (
	"fmt"
	"log"
	"strconv"

	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/ast/utils"
	"dzc/pkg/parser"

	"github.com/pkg/errors"
)

type Parser struct {
	*parser.BaseDZListener

	Consts           map[string]*pkginfo.Const
	incompleteConsts map[string]*pkginfo.Const
}

func New() *Parser {
	return &Parser{
		Consts:           make(map[string]*pkginfo.Const),
		incompleteConsts: make(map[string]*pkginfo.Const),
	}
}

func (p *Parser) EnterConstdecl(ctx *parser.ConstdeclContext) {
	name := ctx.GetId().GetText()

	if p.Consts[name] != nil {
		log.Fatalf("constant %q is redeclared", name)
	}

	tName := utils.FixTypeName(ctx.GetT().GetText())
	if !utils.IsBasicType(tName) {
		log.Fatalf("constant %q is not of basic type (%q)", name, tName)
	}
	t := utils.GetBasicType(tName)

	vText := ctx.GetV().GetText()

	var err error
	var value interface{}

	switch tName {
	case pkginfo.BasicTypeI8:
		value, err = strconv.ParseInt(vText, 10, 8)
	case pkginfo.BasicTypeU8:
		value, err = strconv.ParseUint(vText, 10, 8)
	case pkginfo.BasicTypeI16:
		value, err = strconv.ParseInt(vText, 10, 16)
	case pkginfo.BasicTypeU16:
		value, err = strconv.ParseUint(vText, 10, 16)
	case pkginfo.BasicTypeI32:
		value, err = strconv.ParseInt(vText, 10, 32)
	case pkginfo.BasicTypeU32:
		value, err = strconv.ParseUint(vText, 10, 32)
	case pkginfo.BasicTypeI64:
		value, err = strconv.ParseInt(vText, 10, 64)
	case pkginfo.BasicTypeU64:
		value, err = strconv.ParseUint(vText, 10, 64)
	case pkginfo.BasicTypeSize:
		value, err = strconv.ParseUint(vText, 10, pkginfo.PtrSize)
	case pkginfo.BasicTypeSSize:
		value, err = strconv.ParseInt(vText, 10, pkginfo.PtrSize)
	case pkginfo.BasicTypeBool:
		if vText == pkginfo.BoolFalse {
			value = false
		} else if vText == pkginfo.BoolTrue {
			value = true
		} else {
			err = errors.Errorf("invalid bool value %q", vText)
		}
		//TODO add other types
	default:
		panic(fmt.Sprintf("unknown basic type %q", tName))
	}

	c := &pkginfo.Const{
		Name: name, Type: t,
	}

	if err == nil {
		c.Value = value
		p.Consts[name] = c
		return
	}

	if !utils.IsConst(vText) {
		log.Fatal(err)
	}

	if vc := p.Consts[vText]; vc != nil {
		if vc.Type != c.Type {
			panic(fmt.Sprintf("definition of %q type mismatch", name))
		}

		c.Value = vc.Value
		p.Consts[name] = c

		return
	}

	c.Value = vText
	p.incompleteConsts[name] = c
}

func (p *Parser) FixIncomplete() {
	for name, c := range p.incompleteConsts {
		if p.Consts[name] != nil {
			panic(fmt.Sprintf("constant %q is redeclared", name))
		}

		if vc := p.Consts[c.Value.(string)]; vc != nil {
			if vc.Type != c.Type {
				panic(fmt.Sprintf("definition of %q type mismatch", name))
			}

			c.Value = vc.Value
			p.Consts[name] = c

			delete(p.incompleteConsts, name)
			continue
		}

		panic(fmt.Sprintf("definition of constant %q is incomplete", name))
	}
}
