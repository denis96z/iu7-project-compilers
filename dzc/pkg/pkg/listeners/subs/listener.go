package subs

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/listeners/context"
	"dzc/pkg/pkg/listeners/ctxlog"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	functions  map[string]*syntax.Function
	procedures map[string]*syntax.Procedure

	currentName    string
	currentArgs    map[string]*syntax.Arg
	currentRetType syntax.Type
}

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkg = info
	v.functions = make(map[string]*syntax.Function)
	v.procedures = make(map[string]*syntax.Procedure)
}

func (v *Listener) Finalize() {
	v.pkg.Functions = v.functions
	v.pkg.Procedures = v.procedures
}

func (v *Listener) EnterProcDecl(ctx *parser.ProcDeclContext) {
	name := ctx.GetName().GetText()
	if v.procedures[name] != nil {
		ctxlog.Fatalf(ctx, "proc %q is redeclared", name)
	}
	v.currentName = name
	v.currentArgs = make(map[string]*syntax.Arg)
}

func (v *Listener) ExitProcDecl(ctx *parser.ProcDeclContext) {
	v.procedures[v.currentName] =
		&syntax.Procedure{
			Name: v.currentName,
			Args: v.currentArgs,
		}
}

func (v *Listener) EnterFuncDecl(ctx *parser.FuncDeclContext) {
	name := ctx.GetName().GetText()
	if v.functions[name] != nil {
		ctxlog.Fatalf(ctx, "func %q is redeclared", name)
	}

	v.currentName = name
	v.currentArgs = make(map[string]*syntax.Arg)
	v.currentRetType = v.parseRetType(ctx, ctx.GetTName().GetText())
}

func (v *Listener) EnterProcArg(ctx *parser.ProcArgContext) {
	v.addArg(ctx, ctx.GetName().GetText(), ctx.GetTName().GetText())
}

func (v *Listener) addArg(ctx context.Context, name string, tName string) {
	if v.currentArgs[name] != nil {
		ctxlog.Fatalf(ctx, "arg %q of %q is redeclared", name, v.currentName)
	}
	v.currentArgs[name] =
		&syntax.Arg{
			Name: name,
			Type: v.parseArgType(ctx, tName),
		}
}

func (v *Listener) parseArgType(ctx context.Context, tName string) syntax.Type {
	t := v.pkg.Types[tName]
	if t != nil {
		if !t.IsBasic() && !t.IsEnum() && !t.IsStruct() && !t.IsSlice() {
			failWithTypeNotAllowed(ctx, tName)
		}
	}

	if syntax.IsRefType(tName) {
		vtName := syntax.ParseValueTypeNameFromRefTypeName(tName)

		vt := v.pkg.Types[vtName]
		if vt == nil {
			failWithUnknownType(ctx, vtName)
		}

		syntax.CheckRefValueTypeIsAllowed(ctx, vt)
		t = syntax.NewRef(vt)
	} else if syntax.IsSliceType(tName) {
		itName := syntax.ParseItemTypeNameFromSliceTypeName(tName)

		it := v.pkg.Types[itName]
		if it == nil {
			failWithUnknownType(ctx, itName)
		}

		syntax.CheckSliceValueTypeIsAllowed(ctx, it)
		t = syntax.NewSlice(it)
	} else if t == nil {
		failWithTypeNotAllowed(ctx, tName)
	}

	v.pkg.Types[tName] = t
	return t
}

func (v *Listener) parseRetType(ctx context.Context, tName string) syntax.Type {
	t := v.pkg.Types[tName]
	if t != nil {
		if !t.IsBasic() && !t.IsEnum() && !t.IsStruct() {
			failWithTypeNotAllowed(ctx, tName)
		}
	}
	return t
}

func failWithUnknownType(ctx context.Context, tName string) {
	ctxlog.Fatalf(ctx, "type %q is unknown", tName)
}

func failWithTypeNotAllowed(ctx context.Context, tName string) {
	ctxlog.Fatalf(ctx, "type %q not allowed", tName)
}
