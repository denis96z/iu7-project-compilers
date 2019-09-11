package subparser

import (
	"log"

	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/ast/utils"
	"dzc/pkg/parser"
)

type Parser struct {
	*parser.BaseDZListener

	Functions  map[string]*pkginfo.Function
	Procedures map[string]*pkginfo.Procedure

	types map[string]pkginfo.Type

	currentName    string
	currentArgs    map[string]*pkginfo.Arg
	currentRetType pkginfo.Type
}

func New(types map[string]pkginfo.Type) *Parser {
	return &Parser{
		types:      types,
		Functions:  make(map[string]*pkginfo.Function),
		Procedures: make(map[string]*pkginfo.Procedure),
	}
}

func (p *Parser) EnterProcheader(ctx *parser.ProcheaderContext) {
	name := ctx.GetId().GetText()
	if p.Procedures[name] != nil {
		log.Fatalf("proc %q is redeclared", name)
	}
	p.currentName = name
}

func (p *Parser) ExitProcdecl(ctx *parser.ProcdeclContext) {
	p.Procedures[p.currentName] = &pkginfo.Procedure{
		Name: p.currentName,
		Args: p.currentArgs,
	}
}

func (p *Parser) EnterFuncheader(ctx *parser.FuncheaderContext) {
	name := ctx.GetId().GetText()
	if p.Functions[name] != nil {
		log.Fatalf("func %q is redeclared", name)
	}
	p.currentName = name
}

func (p *Parser) ExitFuncdecl(ctx *parser.FuncdeclContext) {
	p.Functions[p.currentName] = &pkginfo.Function{
		Name:    p.currentName,
		Args:    p.currentArgs,
		RetType: p.currentRetType,
	}
}

func (p *Parser) EnterArgs(ctx *parser.ArgsContext) {
	p.currentArgs = make(map[string]*pkginfo.Arg)
}

func (p *Parser) EnterArgdecl(ctx *parser.ArgdeclContext) {
	name := ctx.GetId().GetText()
	if p.currentArgs[name] != nil {
		log.Fatalf("arg %q in %q declaration is redeclared", name, p.currentName)
	}

	arg := &pkginfo.Arg{
		Name: name,
	}

	tName := utils.FixTypeName(ctx.GetT().GetText())
	if t := p.types[tName]; t != nil {
		arg.Type = t
	} else {
		//TODO make type if base is known
		log.Fatalf("arg %q in %q declaration has unknown type %q", name, p.currentName, tName)
	}

	p.currentArgs[name] = arg
}

func (p *Parser) EnterFuncret(ctx *parser.FuncretContext) {
	tName := utils.FixTypeName(ctx.GetT().GetText())
	if t := p.types[tName]; t != nil {
		p.currentRetType = t
	} else {
		//TODO make type if base is known
		log.Fatalf("return type %q in %q declaration is unknown", tName, p.currentName)
	}
}
