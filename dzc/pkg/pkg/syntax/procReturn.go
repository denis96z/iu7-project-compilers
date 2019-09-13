package syntax

//easyjson:json
type ProcReturn struct {
	Type string `json:"type"`
}

func NewReturn() *ProcReturn {
	return &ProcReturn{
		Type: StatementProcReturn,
	}
}

func (v ProcReturn) GetType() string {
	return v.Type
}

func (v ProcReturn) IsCondition() bool {
	return false
}

func (v ProcReturn) IsLoop() bool {
	return false
}

func (v ProcReturn) IsTrueLoop() bool {
	return false
}

func (v ProcReturn) IsWhileLoop() bool {
	return false
}

func (v ProcReturn) IsBreak() bool {
	return false
}

func (v ProcReturn) IsContinue() bool {
	return false
}

func (v ProcReturn) IsProcCall() bool {
	return false
}

func (v ProcReturn) IsVarDecl() bool {
	return false
}

func (v ProcReturn) IsExpression() bool {
	return false
}

func (v ProcReturn) IsReturn() bool {
	return true
}

func (v ProcReturn) IsProcReturn() bool {
	return true
}

func (v ProcReturn) IsFuncReturn() bool {
	return false
}
