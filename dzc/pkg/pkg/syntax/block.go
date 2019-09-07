package syntax

//easyjson:json
type Block struct {
	Statements []Statement `json:"statements"`
}

func NewBlock() *Block {
	return &Block{
		Statements: make([]Statement, 0),
	}
}
