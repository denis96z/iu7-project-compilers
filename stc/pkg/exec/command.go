package exec

const (
	CommandDefineClass       = "defcls"
	CommandDefineClassAttr   = "defcat"
	CommandDefineClassMethod = "defcmt"
)

func (s *Script) executeLine(pLine *ParsedLine) {
	switch pLine.Command {
	case CommandDefineClass:
		s.defineClass(pLine.Params)
	case CommandDefineClassAttr:
		s.defineClassAttr(pLine.Params)
	case CommandDefineClassMethod:
		s.defineClassMethod(pLine.Params)
	}
}

func (s *Script) defineClass(params []string) {
	s.checkNumParams(params, 2)

	name := params[0]
	superName := params[1]

	s.checkObjectDefined(superName)
	superClass := s.objects[superName]

	s.objects[name] = NewClass(name,
		s.objects[ClassClass].(*Class), superClass.(*Class))
}

func (s *Script) defineClassAttr(params []string) {
	s.checkNumParams(params, 2)

	className := params[0]
	attrName := params[1]

	s.checkObjectDefined(className)
	class := s.objects[className].(*Class)

	class.ClassAttrs[attrName] =
		&Var{Name: attrName}
}

func (s *Script) defineClassMethod(params []string) {
	s.checkNumParams(params, 2)

	className := params[0]
	methodName := params[1]

	s.checkObjectDefined(className)
	class := s.objects[className].(*Class)

	class.ClassMethods[methodName] =
		&Method{
			Name:     methodName,
			Class:    s.objects[ClassMethod].(*Class),
			Args:     make(map[string]*Var),
			Vars:     make(map[string]*Var),
			Commands: make([]*Command, 0),
		}
}

func (s *Script) checkNumParams(params []string, n int) {
	if len(params) != n {
		s.logWrongFormat("%d params expected", n)
	}
}

func (s *Script) checkObjectDefined(name string) {
	if s.objects[name] == nil {
		s.logUndefined(name)
	}
}

func (s *Script) checkObjectUndefined(name string) {
	if s.objects[name] != nil {
		s.logRedefined(name)
	}
}

//
//func (s *Script) defineClass(
//	_ Object, label string,
//	name string, superName string,
//) {
//	s.checkObjectUndefined(name)
//
//	super := s.objects[superName]
//	if super == nil {
//		s.logUndefined(superName)
//	}
//
//	s.objects[name] = NewClass(label,
//		s.objects[ClassClass].(*Class), super.(*Class))
//}
//
//func (s *Script) defineClassAttr(
//	block Object, label string, name string,
//) {
//	s.checkObjectUndefined(name)
//
//	attr := &Var{
//		Name:  label,
//		Class: s.objects[ClassObject].(*Class),
//	}
//
//	s.objects[name] = attr
//
//	class := block.(*Class)
//	class.ClassAttrs[label] = attr
//}
//
//func (s *Script) defineClassMethod(
//	block Object, label string, name string,
//) {
//	if s.objects[name] != nil {
//		s.logRedefined(name)
//	}
//
//	method := &Method{
//		Name:     label,
//		Class:    s.objects[ClassMethod].(*Class),
//		Commands: make([]*Command, 0),
//	}
//
//	s.objects[name] = method
//
//	class := block.(*Class)
//	class.ClassMethods[label] = method
//}
//
//func (s *Script) defineArg(
//	block Object, label string, name string,
//) {
//	s.checkObjectUndefined(name)
//
//	arg := &Var{
//		Name:  label,
//		Class: s.objects[ClassObject].(*Class),
//	}
//
//	s.objects[name] = arg
//
//	switch b := block.(type) {
//	case *Method:
//		b.Args[label] = arg
//	default:
//		panic("not implemented")
//	}
//}
//
//func (s *Script) defineVar(
//	block Object, label string, name string,
//) {
//	s.checkObjectUndefined(name)
//
//	v := &Var{
//		Name:  label,
//		Class: s.objects[ClassObject].(*Class),
//	}
//
//	s.objects[name] = v
//
//	switch b := block.(type) {
//	case *Method:
//		b.Vars[label] = v
//	default:
//		panic("not implemented")
//	}
//}
//
//func (s *Script) pop(block Object, name string) {
//	obj := s.objects[name]
//	s.checkObjectDefined(name)
//
//	switch b := obj.(type) {
//	case *Var:
//		b.Value = s.stack.Pop()
//	default:
//		panic("not implemented")
//	}
//}
