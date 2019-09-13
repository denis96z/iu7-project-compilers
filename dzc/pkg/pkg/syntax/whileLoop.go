package syntax

//easyjson:json
type WhileLoop struct {
	Type      string     `json:"type"`
	Condition Expression `json:"condition"`
	Block     *Block
}

func NewWhileLoop(cond Expression, block *Block) *WhileLoop {
	return &WhileLoop{
		Type:      StatementWhileLoop,
		Condition: cond,
		Block:     block,
	}
}

func (v WhileLoop) GetType() string {
	return v.Type
}

func (v WhileLoop) IsCondition() bool {
	return false
}

func (v WhileLoop) IsLoop() bool {
	return true
}

func (v WhileLoop) IsTrueLoop() bool {
	return false
}

func (v WhileLoop) IsWhileLoop() bool {
	return true
}

func (v WhileLoop) IsBreak() bool {
	return false
}

func (v WhileLoop) IsContinue() bool {
	return false
}

func (v WhileLoop) IsProcCall() bool {
	return false
}

func (v WhileLoop) IsVarDecl() bool {
	return false
}

func (v WhileLoop) IsExpression() bool {
	return false
}

func (v WhileLoop) IsReturn() bool {
	return false
}
