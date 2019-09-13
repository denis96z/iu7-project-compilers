package syntax

//easyjson:json
type Continue struct {
	Type string `json:"type"`
}

func NewContinue() *Continue {
	return &Continue{
		Type: StatementContinue,
	}
}

func (v Continue) GetType() string {
	return v.Type
}

func (v Continue) IsCondition() bool {
	return false
}

func (v Continue) IsLoop() bool {
	return false
}

func (v Continue) IsTrueLoop() bool {
	return false
}

func (v Continue) IsWhileLoop() bool {
	return false
}

func (v Continue) IsBreak() bool {
	return false
}

func (v Continue) IsContinue() bool {
	return true
}

func (v Continue) IsProcCall() bool {
	return false
}

func (v Continue) IsVarDecl() bool {
	return false
}

func (v Continue) IsExpression() bool {
	return false
}

func (v Continue) IsReturn() bool {
	return false
}
