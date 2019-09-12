package pkgparser

import (
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/parser"
)

type Parser struct {
	*parser.BaseDZListener

	pkg *pkginfo.PkgInfo
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) EnterPkg(ctx *parser.PkgContext) {
	p.pkg.Pkg = &pkginfo.Pkg{
		Name: ctx.GetName().GetText(),
	}
}
