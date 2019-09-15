package script

import (
	"stc/pkg/parser/context"
)

type Script struct {
	Classes map[string]*Class
}

func NewScript() *Script {
	return &Script{
		Classes: make(map[string]*Class),
	}
}

func (s *Script) AddClass(
	ctx context.Context,
	name, varName, superName string,
) {
	if s.Classes[name] != nil {
		context.CtxLogFatalf(ctx, "class %q is redefined", name)
	}
	s.Classes[varName] = &Class{
		Name:      name,
		VarName:   varName,
		SuperName: superName,
	}
}

type Class struct {
	Name       string
	VarName    string
	SuperName  string
	IsDeclared bool
}
