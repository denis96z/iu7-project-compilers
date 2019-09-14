package blocks

import (
	"strconv"

	"dzc/pkg/parser"
	"dzc/pkg/pkg/listeners/context"
	"dzc/pkg/pkg/listeners/ctxlog"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"

	"github.com/golang-collections/collections/stack"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	curProc *syntax.Proc
	curFunc *syntax.Func

	curArgs map[string]*syntax.Arg

	//shows block (if | elif | while)
	state int

	exprStack *stack.Stack

	procCall *syntax.ProcCall

	subState  int
	loopState int

	stack *stack.Stack

	vars map[string]*syntax.Var
}

const (
	SubStateProc = iota
	SubStateFunc
)

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkg = info

	v.stack = stack.New()
	v.exprStack = stack.New()
}

func (v *Listener) Finalize() {}

//
//func (v *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
//	log.Println("enter", ctx.GetText())
//}
//
//func (v *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
//	log.Println("exit", ctx.GetText())
//}

func (v *Listener) EnterProcDecl(ctx *parser.ProcDeclContext) {
	v.subState = SubStateProc

	v.curProc = v.pkg.
		Procedures[ctx.GetName().GetText()]

	v.stack = stack.New()
	v.stack.Push(newStackFrame())

	v.vars = make(map[string]*syntax.Var)

	for _, arg := range v.curProc.Args {
		v.vars[arg.Name] = &syntax.Var{
			Name:  arg.Name,
			Type:  arg.GetValueType(),
			Value: arg,
		}
	}
}

func (v *Listener) ExitProcDecl(ctx *parser.ProcDeclContext) {
	f := v.popFrame()
	v.curProc.Body = f.block
}

func (v *Listener) EnterFuncDecl(ctx *parser.FuncDeclContext) {
	v.subState = SubStateFunc
	v.curFunc = v.pkg.
		Functions[ctx.GetName().GetText()]

	v.stack = stack.New()
	v.stack.Push(newStackFrame())

	v.vars = make(map[string]*syntax.Var)
}

func (v *Listener) ExitFuncDecl(ctx *parser.FuncDeclContext) {
	f := v.popFrame()
	v.curFunc.Body = f.block
}

func (v *Listener) EnterBlock(ctx *parser.BlockContext) {
	v.peekFrame().block = syntax.NewBlock()
}

func (v *Listener) ExitBlock(ctx *parser.BlockContext) {
	frame := v.peekFrame()
	if frame == nil {
		return
	}

	for _, item := range frame.scopeVars {
		delete(v.vars, item.Name)
	}
}

func (v *Listener) EnterCondition(ctx *parser.ConditionContext) {
	v.stack.Push(newStackFrame())
}

func (v *Listener) ExitCondition(ctx *parser.ConditionContext) {
	f := v.popFrame()
	cond := syntax.NewCondition(
		f.ifPart, f.elifParts, f.elsePart,
	)

	pf := v.peekFrame()
	pf.block.Statements = append(pf.block.Statements, cond)
}

func (v *Listener) EnterIfConditionBranch(ctx *parser.IfConditionBranchContext) {
	v.stack.Push(newStackFrame())
	v.exprStack.Push(newExprFrame())
}

func (v *Listener) ExitIfConditionBranch(ctx *parser.IfConditionBranchContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	pf.ifPart =
		&syntax.IfConditionPart{
			Cond: v.parseCondition(ctx, v.popExpr().value),
			Body: f.block,
		}
}

func (v *Listener) EnterElifConditionBranch(ctx *parser.ElifConditionBranchContext) {
	pf := v.peekFrame()
	if pf.elifParts == nil {
		pf.elifParts = make([]*syntax.ElIfConditionPart, 0)
	}

	v.stack.Push(newStackFrame())
	v.exprStack.Push(newExprFrame())
}

func (v *Listener) ExitElifConditionBranch(ctx *parser.ElifConditionBranchContext) {
	f := v.popFrame()

	p := &syntax.ElIfConditionPart{
		Cond: v.parseCondition(ctx, v.popExpr().value),
		Body: f.block,
	}

	pf := v.peekFrame()
	pf.elifParts = append(pf.elifParts, p)
}

func (v *Listener) EnterElseConditionBranch(ctx *parser.ElseConditionBranchContext) {
	v.stack.Push(newStackFrame())
}

func (v *Listener) ExitElseConditionBranch(ctx *parser.ElseConditionBranchContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	pf.elsePart =
		&syntax.ElseConditionPart{
			Body: f.block,
		}
}

func (v *Listener) EnterTrueLoop(ctx *parser.TrueLoopContext) {
	v.stack.Push(newStackFrame())
}

func (v *Listener) ExitTrueLoop(ctx *parser.TrueLoopContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	loop := syntax.NewTrueLoop(f.block)

	pf.block.Statements =
		append(pf.block.Statements, loop)
}

func (v *Listener) EnterWhileLoop(ctx *parser.WhileLoopContext) {
	v.stack.Push(newStackFrame())
	v.exprStack.Push(newExprFrame())
}

func (v *Listener) ExitWhileLoop(ctx *parser.WhileLoopContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	loop := syntax.NewWhileLoop(
		v.parseCondition(ctx, v.popExpr().value), f.block,
	)

	pf.block.Statements =
		append(pf.block.Statements, loop)
}

func (v *Listener) EnterVarDecl(ctx *parser.VarDeclContext) {
	name := ctx.GetName().GetText()
	if v.vars[name] != nil {
		ctxlog.Fatalf(ctx, "var %q is redeclared", name)
	}

	v.peekFrame().variable =
		&syntax.Var{
			Name: name,
			Type: v.parseVarType(ctx, ctx.GetTName().GetText()),
		}

	v.exprStack.Push(newExprFrame())
}

func (v *Listener) ExitVarDecl(ctx *parser.VarDeclContext) {
	vr := v.peekFrame().variable
	val := v.popExpr().value

	if vr.Type != val.GetValueType() {
		ctxlog.Fatalf(ctx, "value type mismatch, type is %q", val.GetValueType())
	}

	vr.Value = val
	v.vars[vr.Name] = vr

	varDecl := syntax.NewVarDecl(vr)

	v.peekFrame().block.Statements =
		append(v.peekFrame().block.Statements, varDecl)
}

func (v *Listener) EnterProcCall(ctx *parser.ProcCallContext) {
	name := ctx.GetProcName().GetText()

	v.procCall = &syntax.ProcCall{
		Type:   syntax.StatementProcCall,
		Proc:   v.pkg.Procedures[name],
		Params: make(map[string]*syntax.CallParam),
	}
	if v.procCall.Proc == nil {
		ctxlog.Fatalf(ctx, "unknown procedure %q", name)
	}
}

func (v *Listener) ExitProcCall(ctx *parser.ProcCallContext) {
	f := v.peekFrame()

	if len(v.procCall.Params) != len(v.procCall.Proc.Args) {
		ctxlog.Fatalf(ctx, "wrong number of params in proc call")
	}

	f.block.Statements =
		append(f.block.Statements, v.procCall)
}

func (v *Listener) EnterProcParam(ctx *parser.ProcParamContext) {
	name := ctx.GetName().GetText()
	if v.procCall.Proc.Args[name] == nil {
		ctxlog.Fatalf(ctx, "unknown proc param %q", name)
	}
	if v.procCall.Params[name] != nil {
		ctxlog.Fatalf(ctx, "param %q is specified twice", name)
	}

	expr := newExprFrame()
	expr.param = &syntax.CallParam{
		Name: name,
	}

	v.exprStack.Push(expr)
}

func (v *Listener) ExitProcParam(ctx *parser.ProcParamContext) {
	expr := v.popExpr()
	expr.param.Value = expr.value
	v.procCall.Params[expr.param.Name] = expr.param
}

func (v *Listener) EnterExpression(ctx *parser.ExpressionContext) {
	if ctx.GetVarName() != nil {
		name := ctx.GetVarName().GetText()

		vr := v.vars[name]
		if vr == nil {
			ctxlog.Fatalf(ctx, "unknown var %q", name)
		}

		v.peekExpr().value = vr
		return
	}

	if ctx.GetConstVal() != nil {
		v.peekExpr().value =
			&syntax.ConstVal{
				Type:  nil,
				Value: ctx.GetConstVal().GetText(),
			}
		return
	}

	if ctx.GetConstName() != nil {
		name := ctx.GetConstName().GetText()

		c := v.pkg.Consts[name]
		if c == nil {
			ctxlog.Fatalf(ctx, "const %q is undefined", name)
		}

		v.peekExpr().value = c
		return
	}

	if ctx.GetFuncCallValue() != nil {
		//XXX: separate handler
		return
	}

}

func (v *Listener) EnterFuncCall(ctx *parser.FuncCallContext) {
	name := ctx.GetFuncName().GetText()

	f := v.pkg.Functions[name]
	if f == nil {
		ctxlog.Fatalf(ctx, "func %q is undefined", name)
	}

	v.peekExpr().funcCall = &syntax.FuncCall{
		Func:   f,
		Params: make(map[string]*syntax.CallParam),
	}
}

func (v *Listener) ExitFuncCall(ctx *parser.FuncCallContext) {
	funcCall := v.peekExpr().funcCall

	if len(funcCall.Params) != len(funcCall.Func.Args) {
		ctxlog.Fatalf(ctx, "wrong number of params in func call")
	}

	v.peekExpr().value = funcCall
}

func (v *Listener) EnterFuncParam(ctx *parser.FuncParamContext) {
	funcCall := v.peekExpr().funcCall

	name := ctx.GetName().GetText()
	if funcCall.Func.Args[name] == nil {
		ctxlog.Fatalf(ctx, "func param %q is undefined", name)
	}
	if funcCall.Params[name] != nil {
		ctxlog.Fatalf(ctx, "param %q is specified twice", name)
	}

	expr := newExprFrame()
	expr.param = &syntax.CallParam{
		Name: name,
	}

	v.exprStack.Push(expr)
}

func (v *Listener) ExitFuncParam(ctx *parser.FuncParamContext) {
	expr := v.popExpr()
	expr.param.Value = expr.value
	v.peekExpr().funcCall.Params[expr.param.Name] = expr.param
}

//
//func (v *Listener) EnterReturnStatement(ctx *parser.ReturnStatementContext) {
//	if v.subState == SubStateProc {
//		if ctx.GetValue() != nil {
//			ctxlog.Fatalf(ctx, "return statement is empty for function")
//		}
//		v.curBlock.Statements = append(v.curBlock.Statements, syntax.NewProcReturn())
//	} else if v.subState == SubStateFunc {
//		if ctx.GetValue() == nil {
//			ctxlog.Fatalf(ctx, "return statement is empty for function")
//		}
//	} else {
//		panic("not implemented")
//	}
//}

type exprFrame struct {
	value    syntax.Expression
	param    *syntax.CallParam
	funcCall *syntax.FuncCall
}

func newExprFrame() *exprFrame {
	return &exprFrame{
		param: &syntax.CallParam{},
	}
}

type stackFrame struct {
	//proc || func || loop || while
	block *syntax.Block

	variable *syntax.Var

	//cond
	ifPart    *syntax.IfConditionPart
	elifParts []*syntax.ElIfConditionPart
	elsePart  *syntax.ElseConditionPart

	scopeVars []*syntax.Var
}

func newStackFrame() *stackFrame {
	return &stackFrame{
		scopeVars: make([]*syntax.Var, 0),
	}
}

func (v *Listener) peekExpr() *exprFrame {
	f := v.exprStack.Peek()
	if f == nil {
		return nil
	}
	return f.(*exprFrame)
}

func (v *Listener) popExpr() *exprFrame {
	return v.exprStack.Pop().(*exprFrame)
}

func (v *Listener) peekFrame() *stackFrame {
	f := v.stack.Peek()
	if f == nil {
		return nil
	}
	return f.(*stackFrame)
}

func (v *Listener) popFrame() *stackFrame {
	return v.stack.Pop().(*stackFrame)
}

func (v *Listener) parseCondition(
	ctx context.Context, val syntax.Expression,
) syntax.Expression {
	boolType := syntax.GetBasicType(syntax.BasicTypeBool)

	if val, ok := val.(*syntax.ConstVal); ok {
		boolVal, err := strconv.ParseBool(val.Value.(string))
		if err != nil {
			ctxlog.Fatalf(ctx, "not a bool condition")
		}

		val.Type = boolType
		val.Value = boolVal

		return val
	}

	if val == nil || val.GetValueType() != boolType {
		ctxlog.Fatalf(ctx, "not a bool condition")
	}

	return val
}

func (v *Listener) parseVarType(ctx context.Context, tName string) syntax.Type {
	t := v.pkg.Types[tName]
	if t != nil {
		if t.IsRef() || t.IsSlice() {
			ctxlog.Fatalf(ctx, "type %q is not allowed here", tName)
		}
	}

	if syntax.IsArrayType(tName) {
		itName := syntax.ParseItemTypeNameFromSliceTypeName(tName)

		it := v.pkg.Types[itName]
		if it == nil {
			ctxlog.Fatalf(ctx, "type %q is undefined", itName)
		}

		var size int
		switch vSize := syntax.ParseSizeFromArrayTypeName(tName).(type) {
		case int:
			size = vSize
		case string:
			if c := v.pkg.Consts[vSize]; c != nil {
				if c.Type.GetName() != syntax.BasicTypeSize {
					ctxlog.Fatalf(ctx, "const %q is not of type %q", vSize, syntax.BasicTypeSize)
				}
				size = int(c.Value.(uint64))
			} else {
				ctxlog.Fatalf(ctx, "const %q is undefined", vSize)
			}
		}

		t = syntax.NewArray(it, size)
		v.pkg.Types[tName] = t
	} else if t == nil {
		ctxlog.Fatalf(ctx, "type %q is not allowed here", tName)
	}

	return t
}
