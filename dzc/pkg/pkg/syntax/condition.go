package syntax

//easyjson:json
type Cond struct {
	Type       string     `json:"type"`
	Condition  Expression `json:"condition"`
	IfBlock    *Block     `json:"if_block"`
	ElIfBlocks []*Block   `json:"elif_blocks"`
	ElseBlock  *Block     `json:"else_block"`
}

func NewCondition(
	cond Expression, ifBlock *Block,
	elifBlocks []*Block, elseBlock *Block,
) *Cond {
	return &Cond{
		Type:       StatementCondition,
		Condition:  cond,
		IfBlock:    ifBlock,
		ElIfBlocks: elifBlocks,
		ElseBlock:  elseBlock,
	}
}

func (c Cond) GetType() string {
	return c.Type
}

func (c Cond) IsCondition() bool {
	return true
}

func (c Cond) IsLoop() bool {
	return false
}

func (c Cond) IsTrueLoop() bool {
	return false
}

func (c Cond) IsWhileLoop() bool {
	return false
}

func (c Cond) IsBreak() bool {
	return false
}

func (c Cond) IsContinue() bool {
	return false
}

func (c Cond) IsProcCall() bool {
	return false
}

func (c Cond) IsVarDecl() bool {
	return false
}

func (c Cond) IsExpression() bool {
	return false
}

func (c Cond) IsReturn() bool {
	return false
}
