package ast

import (
	"dzc/pkg/parser"
	"strings"

	"github.com/golang-collections/collections/stack"
)

type PkgInfo struct {
	Pkg        *Pkg
	Types      map[string]*TypeDecl
	Procedures map[string]*ProcDecl
	Functions  map[string]*FuncDecl
}

type Pkg struct {
	Name    string
	Imports map[string]*PkgInfo //TODO use
}

type TypeDecl struct {
	Name     string
	BaseType *TypeDecl
}

type ProcDecl struct {
	Name string
	Args []Arg
}

type FuncDecl struct {
	Name    string
	Args    []Arg
	RetType *TypeDecl
}

type Arg struct {
	Name string
	Type *TypeDecl
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
			Pkg:        makePkg(),
			Types:      makeTypes(),
			Procedures: make(map[string]*ProcDecl),
		},
	}
}

func (p *GlobalParser) EnterPkg(ctx *parser.PkgContext) {
	state := &GlobalParserState{
		State: GlobalParserStatePkg,
		Item: &Pkg{
			Name: ctx.GetName().GetText(), //TODO add imports
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
		Item: &ProcDecl{
			Name: ctx.GetId().GetText(),
			Args: make([]Arg, 0),
		},
	}
	p.stateStack.Push(state)
}

func (p *GlobalParser) EnterFuncheader(ctx *parser.ProcheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateFuncDecl,
		Item: &FuncDecl{
			Name: ctx.GetId().GetText(),
			Args: make([]Arg, 0),
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
		decl := state.Item.(*ProcDecl)
		decl.Args = append(decl.Args, Arg{
			Name: ctx.GetId().GetText(),
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		})
	case GlobalParserStateFuncDecl:
		decl := state.Item.(*FuncDecl)
		decl.Args = append(decl.Args, Arg{
			Name: ctx.GetId().GetText(),
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		})
	}
}

func (p *GlobalParser) EnterFuncret(ctx *parser.FuncretContext) {
	state := p.peekState()
	decl := state.Item.(*FuncDecl)
	decl.RetType = p.findOrMakeType(ctx.GetT().GetText())
}

func (p *GlobalParser) ExitProcheader(ctx *parser.ProcheaderContext) {
	state := p.popState()
	decl := state.Item.(*ProcDecl)
	p.PkgInfo.Procedures[decl.Name] = decl
}

func (p *GlobalParser) ExitFuncheader(ctx *parser.FuncheaderContext) {
	state := p.popState()
	decl := state.Item.(*FuncDecl)
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

func makePkg() *Pkg {
	return &Pkg{
		Imports: make(map[string]*PkgInfo),
	}
}

func fixTypeName(name string) string {
	return strings.ReplaceAll(name, " ", "")
}

func (p *GlobalParser) findOrMakeType(name string) *TypeDecl {
	name = fixTypeName(name)

	t := p.PkgInfo.Types[name]
	if t == nil {
		return &TypeDecl{
			Name:     name,
			BaseType: nil,
		}
	}

	return nil
}

func makeTypes() map[string]*TypeDecl {
	m := make(map[string]*TypeDecl)

	types := []string{
		"i8_t", "u8_t", "i16_t", "u16_t",
		"i32_t", "u32_t", "i64_t", "u64_t",
		"char_t", "bool_t", "size_t", "ssize_t",
	}

	for _, t := range types {
		m[t] = &TypeDecl{
			Name:     t,
			BaseType: nil,
		}
	}

	return m
}
