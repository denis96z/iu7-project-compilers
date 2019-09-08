package ast

import (
	"dzc/pkg/parser"

	"github.com/golang-collections/collections/stack"
)

type PkgInfo struct {
	Pkg        *Pkg
	Types      map[string]*TypeDecl
	Procedures map[string]*ProcDecl
}

type Pkg struct {
	Name string //TODO add imports
}

type TypeDecl struct {
	Name     string
	BaseType *TypeDecl
}

type ProcDecl struct {
	Name string
	Args map[string]*Arg
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

func (p *GlobalParser) EnterProcheader(ctx *parser.ProcheaderContext) {
	state := &GlobalParserState{
		State: GlobalParserStateProcDecl,
		Item: &ProcDecl{
			Name: ctx.GetId().GetText(),
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
		decl := state.Item.(*ProcDecl)

		name := ctx.GetId().GetText()
		decl.Args[name] = &Arg{
			Name: name,
			Type: p.findOrMakeType(ctx.GetT().GetText()),
		}
	}
}

func (p *GlobalParser) ExitProcheader(ctx *parser.ProcheaderContext) {
	state := p.popState()
	decl := state.Item.(*ProcDecl)
	p.PkgInfo.Procedures[decl.Name] = decl
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

func (p *GlobalParser) findOrMakeType(name string) *TypeDecl {
	name = fixTypeName(name)

	t := p.PkgInfo.Types[name]
	if t == nil {
		return &TypeDecl{
			Name: name,
		}
	}

	return t
}

func makePkg() *Pkg {
	return &Pkg{}
}

func fixTypeName(name string) string {
	return name //TODO
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
