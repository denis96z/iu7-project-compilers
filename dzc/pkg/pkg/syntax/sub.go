package syntax

type Procedure struct {
	Name string
	Args map[string]*Arg
}

type Function struct {
	Name    string
	Args    map[string]*Arg
	RetType Type
}

type Arg struct {
	Name string
	Type Type
}
