package exec

type Class struct {
	Name          string
	Class         *Class
	SuperClass    *Class
	ClassAttrs    map[string]*Var
	ClassMethods  map[string]*Method
	ObjectAttrs   map[string]*Var
	ObjectMethods map[string]*Method
}

func NewClass(name string, class *Class, super *Class) *Class {
	c := &Class{
		Name:          name,
		ClassAttrs:    make(map[string]*Var),
		ClassMethods:  make(map[string]*Method),
		ObjectAttrs:   make(map[string]*Var),
		ObjectMethods: make(map[string]*Method),
	}
	if class != nil {
		c.Class = class
	}
	if super != nil {
		c.SuperClass = super
	}
	return c
}

func (c *Class) GetName() string {
	return c.Name
}

func (c *Class) GetClass() *Class {
	return c.Class
}

func (c *Class) GetValue() interface{} {
	return c
}

type Var struct {
	Name  string
	Class *Class
	Value Object
}

func (a *Var) GetName() string {
	return a.Name
}

func (a *Var) GetClass() *Class {
	return a.Class
}

func (a *Var) GetValue() interface{} {
	return a.Value
}

type Method struct {
	Name     string
	Class    *Class
	Args     map[string]*Var
	Vars     map[string]*Var
	Commands []*Command
}

func (m *Method) GetName() string {
	return m.Name
}

func (m *Method) GetClass() *Class {
	return m.Class
}

func (m *Method) GetValue() interface{} {
	return m
}

type Command struct {
	Action CommandFunc
}

type CommandFunc func()
