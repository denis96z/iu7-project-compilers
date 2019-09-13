package pkgs

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"
)

type Listener struct {
	*parser.BaseDZListener

	pkgInfo *pkg.Info

	pkg *syntax.Pkg
}

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkgInfo = info
}

func (v *Listener) EnterPkg(ctx *parser.PkgContext) {
	v.pkg = &syntax.Pkg{
		Name: ctx.GetName().GetText(),
	}
}

func (v *Listener) Finalize() {
	v.pkgInfo.Pkg = v.pkg
}
