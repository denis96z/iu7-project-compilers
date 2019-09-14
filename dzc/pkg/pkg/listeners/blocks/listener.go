package blocks

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/pkg"
	"dzc/pkg/pkg/syntax"

	"github.com/golang-collections/collections/stack"
)

type Listener struct {
	*parser.BaseDZListener

	pkg *pkg.Info

	curProc *syntax.Proc
	curFunc *syntax.Func

	//shows block (if | elif | while)
	state int

	subState  int
	loopState int

	stack *stack.Stack

	vars map[string]*syntax.Var
}

const (
	SubStateProc = iota
	SubStateFunc

	StateNone = iota
	StateIf
	StateElif
	StateWhile
)

func New() *Listener {
	return &Listener{}
}

func (v *Listener) Initialize(info *pkg.Info) {
	v.pkg = info
}

func (v *Listener) Finalize() {}

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
	v.stack.Push(newFrame())

	v.vars = make(map[string]*syntax.Var)
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
	v.stack.Push(newFrame())

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
	v.stack.Push(newFrame())
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
	v.state = StateIf
	v.stack.Push(newFrame())
}

func (v *Listener) ExitIfConditionBranch(ctx *parser.IfConditionBranchContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	pf.ifPart =
		&syntax.IfConditionPart{
			Cond: f.cond,
			Body: f.block,
		}
}

func (v *Listener) EnterElifConditionBranch(ctx *parser.ElifConditionBranchContext) {
	pf := v.peekFrame()
	if pf.elifParts == nil {
		pf.elifParts = make([]*syntax.ElIfConditionPart, 0)
	}

	v.state = StateElif
	v.stack.Push(newFrame())
}

func (v *Listener) ExitElifConditionBranch(ctx *parser.ElifConditionBranchContext) {
	f := v.popFrame()

	p := &syntax.ElIfConditionPart{
		Cond: f.cond,
		Body: f.block,
	}

	pf := v.peekFrame()
	pf.elifParts = append(pf.elifParts, p)
}

func (v *Listener) EnterElseConditionBranch(ctx *parser.ElseConditionBranchContext) {
	v.stack.Push(newFrame())
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
	v.stack.Push(newFrame())
}

func (v *Listener) ExitTrueLoop(ctx *parser.TrueLoopContext) {
	f := v.popFrame()

	pf := v.peekFrame()
	loop := syntax.NewTrueLoop(f.block)

	pf.block.Statements =
		append(pf.block.Statements, loop)
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

type frame struct {
	//proc || func
	block *syntax.Block

	//cond || while
	cond syntax.Expression

	//cond
	ifPart    *syntax.IfConditionPart
	elifParts []*syntax.ElIfConditionPart
	elsePart  *syntax.ElseConditionPart

	scopeVars []*syntax.Var
}

func newFrame() *frame {
	return &frame{
		scopeVars: make([]*syntax.Var, 0),
	}
}

func (v *Listener) peekFrame() *frame {
	f := v.stack.Peek()
	if v == nil {
		return nil
	}
	return f.(*frame)
}

func (v *Listener) popFrame() *frame {
	return v.stack.Pop().(*frame)
}
