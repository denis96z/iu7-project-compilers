package exec

const (
	ClassClass  = "Class"
	ClassObject = "Object"
	ClassBlock  = "Block"
	ClassMethod = "Method"

	ClassInteger = "Integer"
	ClassFloat   = "Float"
	ClassString  = "String"
)

func makeCoreObjects() map[string]Object {
	objects := make(map[string]Object)

	classClass := NewClass(ClassClass, nil, nil)
	classObject := NewClass(ClassObject, classClass, nil)

	classClass.Class = classClass
	classClass.SuperClass = classObject
	classObject.SuperClass = classObject

	classBlock := NewClass(ClassBlock, classObject, classObject)
	classMethod := NewClass(ClassMethod, classObject, classObject)

	objects[ClassClass] = classClass
	objects[ClassObject] = classObject
	objects[ClassBlock] = classBlock
	objects[ClassMethod] = classMethod

	objects[ClassInteger] = NewClass(ClassInteger, classClass, classClass)
	objects[ClassFloat] = NewClass(ClassFloat, classClass, classClass)
	objects[ClassString] = NewClass(ClassString, classClass, classClass)

	return objects
}

const (
	ClassMethodNew    = "new"
	ClassMethodNewArg = ":v"
)

//
//func (s *Script) completeObjectClass(c *Class) {
//	cmtNew := &Method{
//		Name: ClassMethodNew,
//		Class: s.objects[ClassMethod].(*Class),
//		Args: map[string]*Var{
//			ClassMethodNewArg: {
//				Name: ClassMethodNewArg,
//			},
//		},
//		Vars: map[string]*Var{},
//	}
//	cmtNew.Commands = []*Command{
//		{
//			Action: func() {
//				s.stack.Push()
//			},
//		},
//	}
//
//	c.ClassMethods[ClassMethodNew] = cmtNew
//}
