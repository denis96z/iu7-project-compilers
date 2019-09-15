package parser

import (
	"fmt"
	"runtime"

	"stc/pkg/parser/context"
	"stc/pkg/parser/script"
	"stc/pkg/parser/writer"

	"github.com/golang-collections/collections/stack"
)

func NewListener() *listener {
	return &listener{
		idx:     -1,
		s:       script.NewScript(),
		wr:      writer.NewWriter(),
		evStack: newEvalStack(),
	}
}

func (v *listener) GetOutput() string {
	return v.wr.GetOutput()
}

type listener struct {
	*BaseSmallTalkListener

	s  *script.Script
	wr *writer.Writer

	idx     int
	evStack *evalStack

	isClassDef          bool
	isInstanceMethodDef bool
	isBinaryMethodDef   bool
	isKeywordMethodDef  bool

	isMainDef     bool
	isMethodDef   bool
	curMethodName string

	isBlockDef bool
}

func (v *listener) EnterClassDef(ctx *ClassDefContext) {
	name := ctx.GetName().GetText()
	varName := v.makeVar()

	superName := ""
	if ctx.GetSup() != nil {
		superName = ctx.GetSup().GetText()
	}

	v.isClassDef = true
	v.s.AddClass(ctx, name, varName, superName)

	v.wr.WriteStartClassDef(name)
}

func (v *listener) ExitClassDef(ctx *ClassDefContext) {
	v.isClassDef = false
}

func (v *listener) EnterInstanceMethods(ctx *InstanceMethodsContext) {
	v.isInstanceMethodDef = true
}

func (v *listener) ExitInstanceMethods(ctx *InstanceMethodsContext) {
	v.isInstanceMethodDef = false
}

func (v *listener) enterMethod(name string) {
	v.isMethodDef = true
	v.curMethodName = name
}

func (v *listener) exitMethod() {
	ret := ""
	if !v.evStack.isEmpty() {
		ret = v.evStack.pop()
	}
	if ret == stackItemBlock {
		ret = ""
	}

	v.wr.WriteEndMethodDef(ret)
	v.isMethodDef = false
}

func (v *listener) EnterUnaryMethod(ctx *UnaryMethodContext) {
	name := ctx.GetName().GetText()
	v.enterMethod(name)
	v.wr.WriteStartMethodDef(name)
}

func (v *listener) ExitUnaryMethod(ctx *UnaryMethodContext) {
	v.exitMethod()
}

func (v *listener) EnterBinaryMethod(ctx *BinaryMethodContext) {
	name := ctx.GetName().GetText()

	v.enterMethod(name)
	v.isBinaryMethodDef = true

	v.wr.WriteStartMethodDef(name)
}

func (v *listener) ExitBinaryMethod(ctx *BinaryMethodContext) {
	v.exitMethod()
	v.isBinaryMethodDef = false
}

func (v *listener) EnterKeywordMethod(ctx *KeywordMethodContext) {
	name := ctx.GetName().GetText()

	v.enterMethod(name)
	v.isKeywordMethodDef = true

	v.wr.WriteStartMethodDef(name)
}

func (v *listener) ExitKeywordMethod(ctx *KeywordMethodContext) {
	v.exitMethod()
	v.isKeywordMethodDef = false
}

func (v *listener) EnterMethodArg(ctx *MethodArgContext) {
	name := ctx.GetName().GetText()
	v.wr.WriteMethodArgDecl(name)
}

func (v *listener) EnterLocalVar(ctx *LocalVarContext) {
	name := ctx.GetName().GetText()
	v.wr.WriteLocalVarDecl(name)
}

func (v *listener) EnterStatements(ctx *StatementsContext) {
	v.wr.EnsureVars()
}

func (v *listener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {
	value := v.evStack.pop()
	for i := 1; i < ctx.GetChildCount(); i++ {
		name := v.evStack.pop()
		v.wr.WriteAssignment(name, value)
	}
	v.evStack.push(value)
}

func (v *listener) EnterAssignmentItem(ctx *AssignmentItemContext) {
	name := ctx.GetName().GetText()
	v.evStack.push(name)
}

func (v *listener) ExitUnaryMessage(ctx *UnaryMessageContext) {
	method := ctx.GetMethodName().GetText()

	obj := v.evStack.pop()
	res := v.makeVar()

	if method == "new" {
		v.wr.WriteNew(res, obj)
	} else {
		v.wr.WriteUnMethodCall(res, obj, method)
	}
	v.evStack.push(res)
}

func (v *listener) ExitBinaryMessage(ctx *BinaryMessageContext) {
	var isOp bool
	var method string

	if ctx.GetIdName() != nil {
		isOp = false
		method = ctx.GetIdName().GetText()
	} else {
		isOp = true
		method = ctx.GetOpName().GetText()
	}

	o2 := v.evStack.pop()
	o1 := v.evStack.pop()

	r := v.makeVar()
	if isOp {
		v.wr.WriteBinOperation(r, method, o1, o2)
	} else {
		if o1 == "Console" {
			if method == "write" {
				r = o2
				v.wr.WriteSay(o2)
			}
		} else {
			v.wr.WriteBinMethodCall(r, o1, method, o2)
		}
	}

	v.evStack.push(r)
}

func (v *listener) ExitKeywordMessage(ctx *KeywordMessageContext) {
	method := ctx.GetMethodName().GetText()

	args := make([]string, 0)
	for i := 1; i <= ctx.GetChildCount(); i++ {
		args = append(args, v.evStack.pop())
	}

	obj := v.evStack.pop()
	r := v.makeVar()

	v.wr.WriteKeywordMethodCall(r, obj, method, args)
}

const (
	stackItemBlock = "%block%"
)

func (v *listener) EnterValueBlock(ctx *ValueBlockContext) {
	var parent string
	if v.isMethodDef {
		parent = v.curMethodName
	} else if v.isMainDef {
		parent = "main"
	}

	//block item is used to indicate
	//whether the block is empty
	v.evStack.push(stackItemBlock)

	block := v.makeBlock()
	v.wr.WriteStartBlock(block, parent)
}

func (v *listener) ExitValueBlock(ctx *ValueBlockContext) {
	ret := v.evStack.pop()
	if ret == stackItemBlock {
		ret = ""
	}

	v.wr.WriteEndBlock(ret)
}

func (v *listener) EnterBlockArg(ctx *BlockArgContext) {
	name := ctx.GetName().GetText()
	v.wr.WriteBlockArgDecl(name)
}

func (v *listener) EnterMain(ctx *MainContext) {
	v.wr.WriteStartMain()

	declStack := stack.New()
	for _, c := range v.s.Classes {
		if c.IsDeclared {
			break
		}

		for {
			declStack.Push(c)
			if c.SuperName == "" {
				break
			}

			var sc *script.Class
			for _, item := range v.s.Classes {
				if item.Name == c.SuperName {
					sc = item
					break
				}
			}

			if sc == nil {
				context.CtxLogFatalf(ctx, "class %q is undefined", c.SuperName)
			} else {
				if sc.IsDeclared {
					break
				}
				c = sc
			}
		}
		for declStack.Len() != 0 {
			c = declStack.Pop().(*script.Class)
			c.IsDeclared = true
			v.wr.WriteClassVarAssignment(c.Name, c.VarName, c.SuperName)
		}
	}
}

func (v *listener) ExitMain(ctx *MainContext) {
	v.wr.WriteEndMain()
}

const (
	stBinIntPrefix = "2r"
	stOctIntPrefix = "8r"
	stHexIntPrefix = "16r"

	vmBinIntPrefix = "0b"
	vmOctIntPrefix = "0o"
	vmHexIntPrefix = "0x"
)

func (v *listener) EnterValueConstBinInt(ctx *ValueConstBinIntContext) {
	value := vmBinIntPrefix + ctx.GetText()[len(stBinIntPrefix):]
	v.enterConstDecl(value)
}

func (v *listener) EnterValueConstOctInt(ctx *ValueConstOctIntContext) {
	value := vmOctIntPrefix + ctx.GetText()[len(stOctIntPrefix):]
	v.enterConstDecl(value)
}

func (v *listener) EnterValueConstDecInt(ctx *ValueConstDecIntContext) {
	value := ctx.GetText()
	v.enterConstDecl(value)
}

func (v *listener) EnterValueConstHexInt(ctx *ValueConstHexIntContext) {
	value := vmHexIntPrefix + ctx.GetText()[len(stHexIntPrefix):]
	v.enterConstDecl(value)
}

func (v *listener) EnterValueConstFloat(ctx *ValueConstFloatContext) {
	value := ctx.GetText()
	v.enterConstDecl(value)
}

func (v *listener) EnterValueConstString(ctx *ValueConstStringContext) {
	value := ctx.GetText()
	v.enterConstDecl(value)
}

func (v *listener) EnterValueVar(ctx *ValueVarContext) {
	value := ctx.GetText()
	v.evStack.push(value)
}

func (v *listener) EnterValueSelf(ctx *ValueSelfContext) {
	self := ctx.GetText()
	v.evStack.push(self)
}

func (v *listener) enterConstDecl(value string) {
	varName := v.makeVar()
	v.evStack.push(varName)
	v.wr.WriteConstBoxing(varName, value)
}

func (v *listener) makeVar() string {
	v.idx++
	return fmt.Sprintf("$P%d", v.idx)
}

func (v *listener) makeBlock() string {
	v.idx++
	return fmt.Sprintf("block_%d", v.idx)
}

type evalStack struct {
	stack *stack.Stack
}

func newEvalStack() *evalStack {
	return &evalStack{
		stack: stack.New(),
	}
}

func (s *evalStack) push(item string) {
	s.stack.Push(item)
	s.logState("push")
}

func (s *evalStack) isEmpty() bool {
	return s.count() == 0
}

func (s *evalStack) count() int {
	return s.stack.Len()
}

func (s *evalStack) peek() string {
	if s.isEmpty() {
		panic("stack is empty")
	}
	return s.stack.Peek().(string)
}

func (s *evalStack) pop() string {
	if s.isEmpty() {
		panic("stack is empty")
	}
	return s.stack.Pop().(string)
}

func (s *evalStack) clear() {
	for !s.isEmpty() {
		s.pop()
	}
}

func (s *evalStack) logState(op string) {
	//if Debug {
	//	f := getFrame(2)
	//	fName := f.Function[strings.LastIndex(f.Function, ".")+1:]
	//	log.Println(s.name, fmt.Sprintf("%s:%d %s", fName, f.Line, op), s.names)
	//}
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
