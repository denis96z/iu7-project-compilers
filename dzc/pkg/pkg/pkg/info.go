package pkg

import (
	"dzc/pkg/pkg/syntax"
)

type Info struct {
	Pkg        *syntax.Pkg
	Consts     map[string]*syntax.Const
	Types      map[string]syntax.Type
	Procedures map[string]*syntax.Procedure
	Functions  map[string]*syntax.Function
}
