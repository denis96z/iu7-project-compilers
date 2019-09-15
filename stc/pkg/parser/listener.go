package parser

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"stc/pkg/parser/writer"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewListener() *listener {
	return &listener{
		wr:      writer.NewWriter(),
		bkStack: newEvalStack("blocks"),
		evStack: newEvalStack("flow"),
	}
}

func (v *listener) GetOutput() string {
	return v.wr.GetOutput()
}

type listener struct {
	*BaseSmallTalkListener

	wr *writer.Writer

	idx int

	binOp string

	bkStack *evalStack
	evStack *evalStack
}

func (v *listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("-> ", ctx.GetText())
}

func (v *listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("<- ", ctx.GetText())
}

func (v *listener) EnterClassDef(ctx *ClassDefContext) {
	name := ctx.GetName().GetText()

	super := ctx.GetSup()
	if super == nil {
		v.wr.DefineClass(name, nil)
	} else {
		superName := super.GetText()
		v.wr.DefineClass(name, &superName)
	}
}

func (v *listener) ExitClassDef(ctx *ClassDefContext) {
	v.wr.EndClassDefinition()
}

func (v *listener) EnterInstanceVar(ctx *InstanceVarContext) {
	name := ctx.GetName().GetText()
	v.wr.DefineInstanceVar(name)
}

func (v *listener) ExitMethod(ctx *MethodContext) {
	v.wr.EndMethodDefinition()
}

func (v *listener) EnterNamedMethod(ctx *NamedMethodContext) {
	name := ctx.GetName().GetText()
	v.wr.DefineMethod(name)
}

func (v *listener) EnterKeywordMethod(ctx *KeywordMethodContext) {
	name := ctx.GetName().GetText()
	v.wr.DefineKeywordMethod(name[:len(name)-1])
}

func (v *listener) EnterBlock(ctx *BlockContext) {
	tmp := v.makeVar()
	v.bkStack.push(tmp)

	v.wr.DefineBlock(tmp)
}

func (v *listener) ExitBlock(ctx *BlockContext) {
	v.wr.EndBlock()

	tmp := v.makeVar()
	value := v.bkStack.pop()

	v.wr.Move(tmp, value)

	v.evStack.push(tmp)
}

func (v *listener) EnterBlockArg(ctx *BlockArgContext) {
	v.wr.DefineArg(ctx.GetName().GetText())
}

func (v *listener) EnterLocalVar(ctx *LocalVarContext) {
	v.wr.DefineVar(ctx.GetName().GetText())
}

func (v *listener) ExitStatements(ctx *StatementsContext) {
	if v.evStack.isEmpty() {
		v.wr.Return("nil")
	} else {
		v.wr.Return(v.evStack.pop())
	}
}

func (v *listener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {
	for {
		if v.evStack.count() <= 1 {
			break
		}

		value := v.evStack.pop()
		v.wr.Move(v.evStack.peek(), value)
	}
}

func (v *listener) EnterAssignmentItem(ctx *AssignmentItemContext) {
	v.evStack.push(ctx.GetName().GetText())
}

func (v *listener) ExitBinExprTailItem(ctx *BinExprTailItemContext) {
	tmp := v.makeVar()

	v2 := v.evStack.pop()
	op := v.evStack.pop()
	v1 := v.evStack.pop()

	switch op {
	case "+":
		v.wr.Add(tmp, v1, v2)
	case "-":
		v.wr.Sub(tmp, v1, v2)
	case "*":
		v.wr.Mul(tmp, v1, v2)
	case "/":
		v.wr.Div(tmp, v1, v2)
	case "%":
		v.wr.Mod(tmp, v1, v2)
	}

	v.evStack.push(tmp)
}

func (v *listener) ExitBinaryOperator(ctx *BinaryOperatorContext) {
	v.evStack.push(ctx.GetText())
}

func (v *listener) ExitUnaryExpression(ctx *UnaryExpressionContext) {
	tmp := v.makeVar()
	value := v.evStack.pop()
	v.wr.Move(tmp, value)
	v.evStack.push(tmp)
}

func (v *listener) ExitLiteral(ctx *LiteralContext) {
	v.evStack.push(ctx.GetText())
}

func (v *listener) ExitPrimaryIdentifier(ctx *PrimaryIdentifierContext) {
	v.evStack.push(ctx.GetText())
}

//func (v *listener) ExitMethodSendRange(ctx *MethodSendRangeContext) {
//	tmp := v.makeVar()
//
//
//}

func (v *listener) ExitMethodSendItem(ctx *MethodSendItemContext) {
	if ctx.GetName() == nil {
		return
	}

	tmp := v.makeVar()

	args := make([]string, 0)
	for i := 1; i < len(ctx.GetChildren()); i++ {
		args = append(args, v.evStack.pop())
	}

	m := ctx.GetName().GetText()
	recv := v.evStack.peek()

	v.wr.Call(tmp, recv, m, args)
}

func (v *listener) makeVar() string {
	v.idx++
	return fmt.Sprintf("%%r%d", v.idx)
}

type evalStack struct {
	name  string
	names []string
}

func newEvalStack(name string) *evalStack {
	return &evalStack{
		name:  name,
		names: make([]string, 0),
	}
}

var (
	Debug = true
)

func (s *evalStack) push(item string) {
	s.names = append(s.names, item)
	s.logState("push")
}

func (s *evalStack) isEmpty() bool {
	return len(s.names) == 0
}

func (s *evalStack) count() int {
	return len(s.names)
}

func (s *evalStack) peek() string {
	if len(s.names) >= 0 {
		return s.names[len(s.names)-1]
	}
	panic("stack is empty")
}

func (s *evalStack) pop() string {
	if len(s.names) >= 0 {
		v := s.names[len(s.names)-1]
		s.names = s.names[:len(s.names)-1]
		s.logState("pop")
		return v
	}
	panic("stack is empty")
}

func (s *evalStack) clear() {
	s.names = make([]string, 0)
	s.logState("clear")
}

func (s *evalStack) logState(op string) {
	if Debug {
		f := getFrame(2)
		fName := f.Function[strings.LastIndex(f.Function, ".")+1:]
		log.Println(s.name, fmt.Sprintf("%s:%d %s", fName, f.Line, op), s.names)
	}
}

func getFrame(skipFrames int) runtime.Frame {
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
