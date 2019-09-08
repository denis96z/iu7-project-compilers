package ast

import (
	"dzc/pkg/parser"

	"github.com/golang-collections/collections/stack"
)

type PkgInfo struct {
	Pkg        *Pkg
	Types      map[string]*Type
	Procedures map[string]*Procedure
	Functions  map[string]*Function
}

type Pkg struct {
	Name string //TODO add imports
}

type Type struct {
	Name     string
	BaseType *Type
}

type Procedure struct {
	Name string
	Args map[string]*Arg
}

type Function struct {
	Name    string
	Args    map[string]*Arg
	RetType *Type
}

type Arg struct {
	Name string
	Type *Type
}

type GlobalParser struct {
	*parser.BaseDZListener

	stateStack *stack.Stack

	PkgInfo PkgInfo
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
		PkgInfo: PkgInfo{
			Types:      makeTypes(),
			Procedures: make(map[string]*Procedure),
			Functions:  make(map[string]*Function),
		},
	}
}

func (p *GlobalParser) EnterPkg(ctx *parser.PkgContext) {
	state := &GlobalParserState{
		State: GlobalParserStatePkg,
		Item: &Pkg{
			Name: fixPkgName(ctx.GetName().GetText()),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) ExitPkg(ctx *parser.PkgContext) {
	state := p.popState()
	pkg := state.Item.(*Pkg)
	p.PkgInfo.Pkg = pkg
}

func (p *GlobalParser) EnterProcheader(ctx *parser.ProcheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateProcDecl,
		Item: &Procedure{
			Name: fixProcName(ctx.GetId().GetText()),
			Args: make(map[string]*Arg, 0),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) EnterFuncheader(ctx *parser.FuncheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateFuncDecl,
		Item: &Function{
			Name: fixFuncName(ctx.GetId().GetText()),
			Args: make(map[string]*Arg, 0),
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
		decl := state.Item.(*Procedure)

		name := ctx.GetId().GetText()
		decl.Args[name] = &Arg{
			Name: name,
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		}

	case GlobalParserStateFuncDecl:
		decl := state.Item.(*Function)

		name := ctx.GetId().GetText()
		decl.Args[name] = &Arg{
			Name: name,
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		}
	}
}

func (p *GlobalParser) EnterFuncret(ctx *parser.FuncretContext) {
	state := p.peekState()
	decl := state.Item.(*Function)
	decl.RetType = p.findOrMakeType(ctx.GetT().GetText())
}

func (p *GlobalParser) ExitProcheader(ctx *parser.ProcheaderContext) {
	state := p.popState()
	decl := state.Item.(*Procedure)
	p.PkgInfo.Procedures[decl.Name] = decl
}

func (p *GlobalParser) ExitFuncheader(ctx *parser.FuncheaderContext) {
	state := p.popState()
	decl := state.Item.(*Function)
	p.PkgInfo.Functions[decl.Name] = decl
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

func (p *GlobalParser) findOrMakeType(name string) *Type {
	name = fixTypeName(name)

	t := p.PkgInfo.Types[name]
	if t == nil {
		return &Type{
			Name: name,
		}
	}

	return t
}

func fixPkgName(name string) string {
	return name //TODO ?
}

func fixTypeName(name string) string {
	return name //TODO ?
}

func makeTypes() map[string]*Type {
	m := make(map[string]*Type)

	types := []string{
		"i8_t", "u8_t", "i16_t", "u16_t",
		"i32_t", "u32_t", "i64_t", "u64_t",
		"char_t", "bool_t", "size_t", "ssize_t",
	}

	for _, t := range types {
		m[t] = &Type{
			Name:     t,
			BaseType: nil,
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
