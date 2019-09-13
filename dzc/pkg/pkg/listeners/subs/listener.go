package subs

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	functions  map[string]*syntax.Function
	procedures map[string]*syntax.Procedure
}

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Init(info *pkg.Info) {
	v.pkg = info
	v.functions = make(map[string]*syntax.Function)
	v.procedures = make(map[string]*syntax.Procedure)
}

func (v *Listener) Finalize() {
	v.pkg.Functions = v.functions
	v.pkg.Procedures = v.procedures
}
