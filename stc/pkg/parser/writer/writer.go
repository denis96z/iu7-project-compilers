package writer

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Writer struct {
	output string
	offset string

	blockStack *blockStack
	blockQueue *blockStack
}

func NewWriter() *Writer {
	return &Writer{
		blockStack: newBlockStack(),
		blockQueue: newBlockStack(),
	}
}

func (w *Writer) GetOutput() string {
	return w.output
}

const (
	offsetDiff = "    "
)

func (w *Writer) doIncOffset(offset *string) {
	*offset += offsetDiff
}

func (w *Writer) incOffset() {
	if w.blockStack.isEmpty() {
		w.doIncOffset(&w.offset)
	} else {
		b := w.blockStack.peek()
		w.doIncOffset(&b.offset)
	}
}

func (w *Writer) doDecOffset(offset *string) {
	*offset = (*offset)[:len(*offset)-(len(offsetDiff))]
}

func (w *Writer) decOffset() {
	if w.blockStack.isEmpty() {
		w.doDecOffset(&w.offset)
	} else {
		b := w.blockStack.peek()
		w.doDecOffset(&b.offset)
	}
}

func (w *Writer) doWriteLine(buf *string, offset string, line string) {
	*buf += offset + line + "\n"
}

func (w *Writer) writeLine(line string) {
	if w.blockStack.isEmpty() {
		w.doWriteLine(&w.output, w.offset, line)
	} else {
		b := w.blockStack.peek()
		w.doWriteLine(&b.output, b.offset, line)
	}
}

func (w *Writer) writeArg(name string) {
	w.writeLine(fmt.Sprintf(`.param pmc %s`, name))
	w.writeLine(fmt.Sprintf(`.lex '%s', %s`, name, name))
	if !w.blockStack.isEmpty() { //TODO fix for main
		w.blockStack.peek().vars[name] = struct{}{}
	}
}

func (w *Writer) writeVar(name string) {
	w.writeLine(fmt.Sprintf(`.local pmc %s`, name))
	w.writeLine(fmt.Sprintf(`.lex '%s', %s`, name, name))
	if !w.blockStack.isEmpty() { //TODO fix for main
		w.blockStack.peek().vars[name] = struct{}{}
	}
}

func (w *Writer) WriteStartClassDef(name string) {
	w.writeLine(fmt.Sprintf(`.namespace ['%s']`, name))

	w.WriteStartVirtualMethodDef("get_string")
	w.WriteReturn(fmt.Sprintf(`'%s'`, name))
	w.WriteEndMethodDef("")
}

func (w *Writer) WriteStartMethodDef(name string) {
	b := &block{
		name: name,
		vars: make(map[string]struct{}),
	}
	w.blockStack.push(b)

	w.writeLine(fmt.Sprintf(`.sub '%s' :method`, name))
	w.incOffset()
}

func (w *Writer) WriteStartVirtualMethodDef(name string) {
	b := &block{
		name: name,
		vars: make(map[string]struct{}),
	}
	w.blockStack.push(b)

	w.writeLine(fmt.Sprintf(`.sub '%s' :method :vtable`, name))
	w.incOffset()
}

func (w *Writer) WriteEndMethodDef(ret string) {
	if ret != "" {
		w.WriteReturn(ret)
	}
	w.decOffset()
	w.writeLine(`.end`)

	b := w.blockStack.pop()
	w.blockQueue.push(b)

	w.WriteAllBlocks()
}

func (w *Writer) WriteStartBlock(name, parent string) {
	if !w.blockStack.isEmpty() {
		parent = w.blockStack.peek().name
	}

	b := &block{
		name: name,
		vars: make(map[string]struct{}),
	}
	w.blockStack.push(b)

	w.writeLine(fmt.Sprintf(`.sub '%s' :outer('%s')`, name, parent))
	w.incOffset()
}

func (w *Writer) WriteEndBlock(ret string) {
	if ret != "" {
		w.WriteReturn(ret)
	}

	w.decOffset()
	w.writeLine(`.end`)

	b := w.blockStack.pop()
	w.blockQueue.push(b)
}

func (w *Writer) WriteBlockArgDecl(name string) {
	w.writeArg(name)
}

func (w *Writer) EnsureVars() {
	if w.blockStack.count() > 1 {
		b := w.blockStack.pop()

		for k := range w.blockStack.peek().vars {
			if _, ok := b.vars[k]; ok {
				continue
			}
			w.doWriteLine(&b.output, b.offset, fmt.Sprintf(`.local pmc %s`, k))
			w.doWriteLine(&b.output, b.offset, fmt.Sprintf(`find_lex %s, '%s'`, k, k))
			b.vars[k] = struct{}{}
		}

		w.blockStack.push(b)
	}
}

func (w *Writer) WriteAllBlocks() {
	for !w.blockQueue.isEmpty() {
		v := w.blockQueue.pop().output
		if strings.HasSuffix(v, "\n") {
			v = v[:len(v)-1]
		}
		w.writeLine(v)
	}
}

func (w *Writer) WriteMethodArgDecl(name string) {
	w.writeArg(name)
}

func (w *Writer) WriteLocalVarDecl(name string) {
	w.writeVar(name)
}

func (w *Writer) WriteAssignment(name, value string) {
	w.writeLine(fmt.Sprintf(`%s = %s`, name, value))
}

func (w *Writer) WriteBinOperation(r, o, o1, o2 string) {
	w.writeLine(fmt.Sprintf(`%s = %s %s %s`, r, o1, o, o2))
}

func (w *Writer) WriteUnMethodCall(r, obj, mt string) {
	w.writeLine(fmt.Sprintf(`%s = %s.'%s'()`, r, obj, mt))
}

func (w *Writer) WriteBinMethodCall(r, obj, mt, arg string) {
	w.writeLine(fmt.Sprintf(`%s = %s.'%s'(%s)`, r, obj, mt, arg))
}

func (w *Writer) WriteKeywordMethodCall(r, obj, mt string, args []string) {
	w.WriteBinMethodCall(r, obj, mt, strings.Join(args, ","))
}

func (w *Writer) WriteNew(r, c string) {
	w.writeLine(fmt.Sprintf(`%s = new '%s'`, r, c))
}

func (w *Writer) WriteSay(v string) {
	w.writeLine(fmt.Sprintf(`say %s`, v))
}

func (w *Writer) WriteReturn(value string) {
	w.writeLine(fmt.Sprintf(`.return(%s)`, value))
}

func (w *Writer) WriteStartMain() {
	b := &block{
		name: "main",
		vars: make(map[string]struct{}),
	}
	w.blockStack.push(b)

	w.writeLine(`.namespace []`)
	w.writeLine(`.sub 'main' :main`)
	w.incOffset()
}

func (w *Writer) WriteEndMain() {
	w.decOffset()
	w.writeLine(`.end`)

	b := w.blockStack.pop()
	w.blockQueue.push(b)

	w.WriteAllBlocks()
}

func (w *Writer) WriteClassVarAssignment(name, varName, superName string) {
	if superName == "" {
		w.writeLine(fmt.Sprintf(`%s = newclass '%s'`, varName, name))
	} else {
		w.writeLine(fmt.Sprintf(`%s = subclass '%s','%s'`, varName, superName, name))
	}
}

func (w *Writer) WriteConstBoxing(varName, value string) {
	w.writeLine(fmt.Sprintf(`%s = box %s`, varName, value))
}

type block struct {
	name   string
	output string
	offset string
	vars   map[string]struct{}
}

type blockStack struct {
	stack *stack.Stack
}

func newBlockStack() *blockStack {
	return &blockStack{
		stack: stack.New(),
	}
}

func (s *blockStack) push(block *block) {
	s.stack.Push(block)
}

func (s *blockStack) peek() *block {
	v := s.stack.Peek()
	if v == nil {
		panic("empty block stack")
	}
	return v.(*block)
}

func (s *blockStack) pop() *block {
	_ = s.peek()
	return s.stack.Pop().(*block)
}

func (s *blockStack) count() int {
	return s.stack.Len()
}

func (s *blockStack) isEmpty() bool {
	return s.count() == 0
}

type blockQueue struct {
	queue *queue.Queue
}

func newBlockQueue() *blockQueue {
	return &blockQueue{
		queue: queue.New(),
	}
}

func (q *blockQueue) enqueue(b *block) {
	q.queue.Enqueue(b)
}

func (q *blockQueue) peek() *block {
	if q.isEmpty() {
		panic("empty block queue")
	}
	return q.queue.Peek().(*block)
}

func (q *blockQueue) dequeue() *block {
	_ = q.peek()
	return q.queue.Dequeue().(*block)
}

func (q *blockQueue) count() int {
	return q.queue.Len()
}

func (q *blockQueue) isEmpty() bool {
	return q.count() == 0
}
