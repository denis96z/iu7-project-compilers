package constparser

import (
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/ast/utils"
	"dzc/pkg/parser"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

type ConstParser struct {
	*parser.BaseDZListener

	KnownConsts   map[string]*pkginfo.Const
	unknownConsts map[string]*pkginfo.Const
}

func NewParser() *ConstParser {
	return &ConstParser{
		KnownConsts:   make(map[string]*pkginfo.Const),
		unknownConsts: make(map[string]*pkginfo.Const),
	}
}

func (p *ConstParser) EnterConstdecl(ctx *parser.ConstdeclContext) {
	name := ctx.GetId().GetText()

	if p.KnownConsts[name] != nil {
		panic("constant %q is redeclared")
	}

	tName := utils.FixTypeName(ctx.GetT().GetText())
	if !utils.IsBasicType(tName) {
		panic("constant of type %q not allowed: basic type is required")
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
		p.KnownConsts[name] = c
		return
	}

	if !utils.IsConst(vText) {
		panic(err)
	}

	if vc := p.KnownConsts[vText]; vc != nil {
		if vc.Type != c.Type {
			panic(fmt.Sprintf("definition of %q type mismatch", name))
		}

		c.Value = vc.Value
		p.KnownConsts[name] = c

		return
	}

	c.Value = vText
	p.unknownConsts[name] = c
}

func (p *ConstParser) FixDefinitions() {
	for name, c := range p.unknownConsts {
		if p.KnownConsts[name] != nil {
			panic(fmt.Sprintf("constant %q is redeclared", name))
		}

		if vc := p.KnownConsts[c.Value.(string)]; vc != nil {
			if vc.Type != c.Type {
				panic(fmt.Sprintf("definition of %q type mismatch", name))
			}

			c.Value = vc.Value
			p.KnownConsts[name] = c

			delete(p.unknownConsts, name)
			continue
		}

		panic(fmt.Sprintf("definition of constant %q is incomplete", name))
	}
}
