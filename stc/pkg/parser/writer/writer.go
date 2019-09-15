package writer

import (
	"fmt"
	"strings"
)

type Writer struct {
	output string
	offset string
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) GetOutput() string {
	return w.output
}

const (
	offsetDiff = "    "

	varPrefix = "$"
)

func (w *Writer) incOffset() {
	w.offset += offsetDiff
}

func (w *Writer) decOffset() {
	w.offset = w.offset[:len(w.offset)-(len(offsetDiff))]
}

func (w *Writer) addCmd(format string, args ...interface{}) {
	w.output += w.offset + fmt.Sprintf(format+"\n", args...)
}

func (w *Writer) addEnd() {
	w.addCmd("end")
}

func (w *Writer) DefineClass(name string, superName *string) {
	if superName == nil {
		w.addCmd("defc %s", name)
	} else {
		w.addCmd("defc %s, %s", name, *superName)
	}
	w.incOffset()
}

func (w *Writer) EndClassDefinition() {
	w.decOffset()
	w.addEnd()
}

func (w *Writer) DefineInstanceVar(name string) {
	w.addCmd("deff %s", name)
}

func (w *Writer) DefineMethod(name string) {
	w.addCmd("defm %s", name)
	w.incOffset()
}

func (w *Writer) DefineKeywordMethod(name string) {
	w.addCmd("defkwm %s", name)
	w.incOffset()
}

func (w *Writer) DefineArg(name string) {
	w.addCmd("defa %s", name)
}

func (w *Writer) DefineBlock(name string) {
	w.addCmd("defb %s", name)
	w.incOffset()
}

func (w *Writer) EndBlock() {
	w.decOffset()
	w.addEnd()
}

func (w *Writer) EndMethodDefinition() {
	w.decOffset()
	w.addEnd()
}

func (w *Writer) DefineVar(name string) {
	w.addCmd("defv %s", name)
}

func (w *Writer) Move(dst, src string) {
	w.addCmd("mov %s, %s", dst, src)
}

func (w *Writer) binOp(name, dst, src1, src2 string) {
	w.addCmd("%s %s, %s, %s", name, dst, src1, src2)
}

func (w *Writer) Add(dst, src1, src2 string) {
	w.binOp("add", dst, src1, src2)
}

func (w *Writer) Sub(dst, src1, src2 string) {
	w.binOp("sub", dst, src1, src2)
}

func (w *Writer) Mul(dst, src1, src2 string) {
	w.binOp("mul", dst, src1, src2)
}

func (w *Writer) Div(dst, src1, src2 string) {
	w.binOp("div", dst, src1, src2)
}

func (w *Writer) Mod(dst, src1, src2 string) {
	w.binOp("mod", dst, src1, src2)
}

func (w *Writer) Call(dst, recv, name string, args []string) {
	var ar string
	if len(args) > 0 {
		ar = ", " + strings.Join(args, ", ")
	}
	w.addCmd("call %s, %s, %s%s", dst, recv, name, ar)
}

func (w *Writer) Return(result string) {
	w.addCmd("ret %s", result)
}
