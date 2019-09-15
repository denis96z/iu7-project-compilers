package exec

type Block struct {
	Name     string
	Class    *Class
	Commands []*Command
}

func (s *Script) newBlock(name string) *Block {
	return &Block{
		Name:     name,
		Class:    s.objects[ClassBlock].(*Class),
		Commands: make([]*Command, 0),
	}
}

func (b *Block) GetName() string {
	return b.Name
}

func (b *Block) GetClass() *Class {
	return b.Class
}

func (b *Block) GetValue() interface{} {
	return b
}
