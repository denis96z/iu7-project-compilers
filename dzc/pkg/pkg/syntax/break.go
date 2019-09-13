package syntax

//easyjson:json
type Break struct {
	Type string `json:"break"`
}

func NewBreak() *Break {
	return &Break{
		Type: StatementBreak,
	}
}

func (v Break) GetType() string {
	return v.Type
}

func (v Break) IsCondition() bool {
	return false
}

func (v Break) IsLoop() bool {
	return false
}

func (v Break) IsTrueLoop() bool {
	return false
}

func (v Break) IsWhileLoop() bool {
	return false
}

func (v Break) IsBreak() bool {
	return true
}

func (v Break) IsContinue() bool {
	return false
}

func (v Break) IsProcCall() bool {
	return false
}

func (v Break) IsVarDecl() bool {
	return false
}

func (v Break) IsExpression() bool {
	return false
}

func (v Break) IsReturn() bool {
	return false
}
