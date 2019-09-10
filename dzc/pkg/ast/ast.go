package ast

import (
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/parser"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/golang-collections/collections/stack"
)

type GlobalParser struct {
	*parser.BaseDZListener

	stateStack *stack.Stack

	PkgInfo pkginfo.PkgInfo
}

type GlobalParserState struct {
	State int
	Item  interface{}
}

const (
	GlobalParserStatePkg = iota
	GlobalParserStateProcDecl
	GlobalParserStateFuncDecl
)

func NewGlobalParser() *GlobalParser {
	return &GlobalParser{
		stateStack: stack.New(),
		PkgInfo: pkginfo.PkgInfo{
			Types:      makeTypes(),
			Procedures: make(map[string]*pkginfo.Procedure),
			Functions:  make(map[string]*pkginfo.Function),
		},
	}
}

func (p *GlobalParser) EnterPkg(ctx *parser.PkgContext) {
	state := &GlobalParserState{
		State: GlobalParserStatePkg,
		Item: &pkginfo.Pkg{
			Name: fixPkgName(ctx.GetName().GetText()),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) ExitPkg(ctx *parser.PkgContext) {
	state := p.popState()
	pkg := state.Item.(*pkginfo.Pkg)
	p.PkgInfo.Pkg = pkg
}

func (p *GlobalParser) EnterProcheader(ctx *parser.ProcheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateProcDecl,
		Item: &pkginfo.Procedure{
			Name: fixProcName(ctx.GetId().GetText()),
			Args: make(map[string]*pkginfo.Arg, 0),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) EnterFuncheader(ctx *parser.FuncheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateFuncDecl,
		Item: &pkginfo.Function{
			Name: fixFuncName(ctx.GetId().GetText()),
			Args: make(map[string]*pkginfo.Arg, 0),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) EnterArgdecl(ctx *parser.ArgdeclContext) {
	state := p.peekState()
	if state == nil {
		return //TODO functions
	}

	switch state.State {
	case GlobalParserStateProcDecl:
		decl := state.Item.(*pkginfo.Procedure)

		name := ctx.GetId().GetText()
		decl.Args[name] = &pkginfo.Arg{
			Name: name,
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		}

	case GlobalParserStateFuncDecl:
		decl := state.Item.(*pkginfo.Function)

		name := ctx.GetId().GetText()
		decl.Args[name] = &pkginfo.Arg{
			Name: name,
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		}
	}
}

func (p *GlobalParser) EnterFuncret(ctx *parser.FuncretContext) {
	state := p.peekState()
	decl := state.Item.(*pkginfo.Function)
	decl.RetType = p.findOrMakeType(ctx.GetT().GetText())
}

func (p *GlobalParser) ExitProcheader(ctx *parser.ProcheaderContext) {
	state := p.popState()
	decl := state.Item.(*pkginfo.Procedure)
	p.PkgInfo.Procedures[decl.Name] = decl
}

func (p *GlobalParser) ExitFuncheader(ctx *parser.FuncheaderContext) {
	state := p.popState()
	decl := state.Item.(*pkginfo.Function)
	p.PkgInfo.Functions[decl.Name] = decl
}

func (p *GlobalParser) EnterTypedecl(ctx *parser.TypedeclContext) {
	name := fixTypeName(ctx.GetId().GetText())
	if p.findType(name) != nil {
		panic(fmt.Sprintf("type %q already exists", name))
	}

	t := &pkginfo.SimpleType{
		Name: name,
	}

	exp := fixTypeName(ctx.GetT().GetText())
	if tExp := p.findType(exp); tExp != nil {
		t.BaseType = tExp
	} else if isSimpleType(exp) {
		t.BaseType = p.findType(exp) //XXX if nil will be set later
	} else if isRefType(exp) {
		nameBase := parseBaseNameFromRefTypeName(exp)

		bt := p.findType(nameBase)
		if bt != nil {
			rt := &pkginfo.RefType{
				Name:     exp,
				BaseType: bt,
			}
			p.PkgInfo.Types[exp] = rt
			t.BaseType = rt
		}
	} else if isArrayType(exp) {
		nameBase := parseBaseNameFromArrayTypeName(exp)

		size := parseSizeFromArrayTypeName(exp).(int) //TODO check if constparser name

		bt := p.findType(nameBase)
		if bt != nil {
			rt := &pkginfo.ArrayType{
				Name:     exp,
				Size:     size,
				BaseType: bt,
			}
			p.PkgInfo.Types[exp] = rt
			t.BaseType = rt
		}
	} else {
		panic("unknown type") //XXX should not happen
	}

	p.PkgInfo.Types[name] = t
}

func (p *GlobalParser) popState() *GlobalParserState {
	state := p.stateStack.Pop()
	if state == nil {
		return nil
	}
	return state.(*GlobalParserState)
}

func (p *GlobalParser) peekState() *GlobalParserState {
	state := p.stateStack.Peek()
	if state == nil {
		return nil
	}
	return state.(*GlobalParserState)
}

func (p *GlobalParser) findType(name string) pkginfo.Type {
	name = fixTypeName(name)
	return p.PkgInfo.Types[name]
}

func (p *GlobalParser) findOrMakeType(name string) pkginfo.Type {
	name = fixTypeName(name)

	t := p.PkgInfo.Types[name]
	if t == nil {
		if isSimpleType(name) {
			return &pkginfo.SimpleType{
				Name: name,
			}
		} else if isRefType(name) {
			nameBase := parseBaseNameFromRefTypeName(name)

			bt := p.PkgInfo.Types[nameBase]
			t = &pkginfo.RefType{
				Name:     name,
				BaseType: bt,
			}

			if bt != nil {
				p.PkgInfo.Types[name] = t
			}
		} else if isArrayType(name) {
			nameBase := parseBaseNameFromArrayTypeName(name)

			size := parseSizeFromArrayTypeName(name).(int) //TODO check if constparser name

			bt := p.PkgInfo.Types[nameBase]
			t = &pkginfo.ArrayType{
				Name:     name,
				Size:     size,
				BaseType: bt,
			}

			if bt != nil {
				p.PkgInfo.Types[name] = t
			}
		} else {
			panic("unknown type") //XXX should not happen
		}
	}

	return t
}

func isSimpleType(name string) bool {
	return !strings.HasPrefix(name, pkginfo.RefSymbol) && //TODO improve
		!strings.HasPrefix(name, "[") //TODO constparser?
}

func isRefType(name string) bool {
	return strings.HasPrefix(name, pkginfo.RefSymbol)
}

func isArrayType(name string) bool {
	return strings.HasPrefix(name, "[")
}

func fixPkgName(name string) string {
	return name //TODO ?
}

func fixTypeName(name string) string {
	return name //TODO ?
}

func parseBaseNameFromRefTypeName(name string) string {
	name = fixTypeName(name)
	return name[1:]
}

func parseBaseNameFromArrayTypeName(name string) string {
	name = fixTypeName(name)
	return name[1:strings.Index(name, ":")]
}

func parseSizeFromArrayTypeName(name string) interface{} {
	name = fixTypeName(name)

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

func makeTypes() map[string]pkginfo.Type {
	m := make(map[string]pkginfo.Type)

	types := []string{
		"i8_t", "u8_t", "i16_t", "u16_t",
		"i32_t", "u32_t", "i64_t", "u64_t",
		"char_t", "bool_t", "size_t", "ssize_t",
	}

	for _, t := range types {
		name := fixTypeName(t)
		m[name] = &pkginfo.BasicType{
			Name: t,
		}
	}

	return m
}

func fixProcName(name string) string {
	return name //TODO ?
}

func fixFuncName(name string) string {
	return name //TODO ?
}
