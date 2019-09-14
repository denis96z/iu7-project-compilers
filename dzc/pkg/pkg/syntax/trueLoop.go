package syntax

//easyjson:json
type TrueLoop struct {
	Type string `json:"type"`
	Body *Block `json:"body"`
}

func NewTrueLoop(block *Block) *TrueLoop {
	return &TrueLoop{
		Type: StatementTrueLoop,
		Body: block,
	}
}

func (v TrueLoop) GetType() string {
	return v.Type
}

func (v TrueLoop) IsCondition() bool {
	return false
}

func (v TrueLoop) IsLoop() bool {
	return true
}

func (v TrueLoop) IsTrueLoop() bool {
	return true
}

func (v TrueLoop) IsWhileLoop() bool {
	return false
}

func (v TrueLoop) IsBreak() bool {
	return false
}

func (v TrueLoop) IsContinue() bool {
	return false
}

func (v TrueLoop) IsProcCall() bool {
	return false
}

func (v TrueLoop) IsVarDecl() bool {
	return false
}

func (v TrueLoop) IsExpression() bool {
	return false
}

func (v TrueLoop) IsReturn() bool {
	return false
}

func (v TrueLoop) IsProcReturn() bool {
	return false
}

func (v TrueLoop) IsFuncReturn() bool {
	return false
}
