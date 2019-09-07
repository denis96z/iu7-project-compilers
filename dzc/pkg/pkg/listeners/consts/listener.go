package consts

import (
	"fmt"
	"log"
	"strconv"

	"dzc/pkg/parser"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"

	"github.com/pkg/errors"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	consts           map[string]*syntax.Const
	incompleteConsts map[string]*syntax.Const
}

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkg = info
	v.consts = make(map[string]*syntax.Const)
	v.incompleteConsts = make(map[string]*syntax.Const)
}

func (v *Listener) EnterConstDecl(ctx *parser.ConstDeclContext) {
	name := ctx.GetName().GetText()

	if v.consts[name] != nil {
		log.Fatalf("constant %q is redeclared", name)
	}

	tName := ctx.GetTName().GetText()
	if !syntax.IsBasicType(tName) {
		log.Fatalf("constant %q is not of basic type (%q)", name, tName)
	}
	t := syntax.GetBasicType(tName)

	vText := ctx.GetValue().GetText()

	var err error
	var value interface{}

	switch tName {
	case syntax.BasicTypeI8:
		value, err = strconv.ParseInt(vText, 10, 8)
	case syntax.BasicTypeU8:
		value, err = strconv.ParseUint(vText, 10, 8)
	case syntax.BasicTypeI16:
		value, err = strconv.ParseInt(vText, 10, 16)
	case syntax.BasicTypeU16:
		value, err = strconv.ParseUint(vText, 10, 16)
	case syntax.BasicTypeI32:
		value, err = strconv.ParseInt(vText, 10, 32)
	case syntax.BasicTypeU32:
		value, err = strconv.ParseUint(vText, 10, 32)
	case syntax.BasicTypeI64:
		value, err = strconv.ParseInt(vText, 10, 64)
	case syntax.BasicTypeU64:
		value, err = strconv.ParseUint(vText, 10, 64)
	case syntax.BasicTypeSize:
		value, err = strconv.ParseUint(vText, 10, syntax.PtrSize)
	case syntax.BasicTypeSSize:
		value, err = strconv.ParseInt(vText, 10, syntax.PtrSize)
	case syntax.BasicTypeBool:
		if vText == syntax.BoolFalse {
			value = false
		} else if vText == syntax.BoolTrue {
			value = true
		} else {
			err = errors.Errorf("invalid bool value %q", vText)
		}
	case syntax.BasicTypeChar:
		constVal := syntax.ParseChar(ctx, vText)
		value = constVal.Value //TODO fix
	default:
		panic("not implemented")
	}

	c := &syntax.Const{
		Name: name, Type: t,
	}

	if err == nil {
		c.Value = value
		v.consts[name] = c
		return
	}

	if !syntax.IsConst(vText) {
		log.Fatal(err)
	}

	if vc := v.consts[vText]; vc != nil {
		if vc.Type != c.Type {
			panic(fmt.Sprintf("definition of %q type mismatch", name))
		}

		c.Value = vc.Value
		v.consts[name] = c

		return
	}

	c.Value = vText
	v.incompleteConsts[name] = c
}

func (v *Listener) fixIncomplete() {
	for name, c := range v.incompleteConsts {
		if v.consts[name] != nil {
			log.Fatalf("constant %q is redeclared", name)
		}

		if vc := v.consts[c.Value.(string)]; vc != nil {
			if vc.Type != c.Type {
				log.Fatalf("definition of %q type mismatch", name)
			}

			c.Value = vc.Value
			v.consts[name] = c

			continue
		}

		log.Fatalf("definition of constant %q is incomplete", name)
	}
}

func (v *Listener) Finalize() {
	v.fixIncomplete()
	v.pkg.Consts = v.consts
}
