package pkgparser

import (
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/parser"
)

type Parser struct {
	*parser.BaseDZListener

	Pkg *pkginfo.Pkg
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) EnterPkg(ctx *parser.PkgContext) {
	p.Pkg = &pkginfo.Pkg{
		Name: ctx.GetName().GetText(),
	}
}
