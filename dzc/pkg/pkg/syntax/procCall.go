package syntax

//easyjson:json
type ProcCall struct {
	Type   string       `json:"type"`
	Proc   *Procedure   `json:"proc"`
	Params []Expression `json:"params"`
}

func NewProcCall(proc *Procedure, params []Expression) *ProcCall {
	return &ProcCall{
		Type:   StatementProcCall,
		Proc:   proc,
		Params: params,
	}
}

func (v ProcCall) GetType() string {
	return v.Type
}

func (v ProcCall) IsCondition() bool {
	return false
}

func (v ProcCall) IsLoop() bool {
	return false
}

func (v ProcCall) IsTrueLoop() bool {
	return false
}

func (v ProcCall) IsWhileLoop() bool {
	return false
}

func (v ProcCall) IsBreak() bool {
	return false
}

func (v ProcCall) IsContinue() bool {
	return false
}

func (v ProcCall) IsProcCall() bool {
	return true
}

func (v ProcCall) IsVarDecl() bool {
	return false
}

func (v ProcCall) IsExpression() bool {
	return false
}

func (v ProcCall) IsReturn() bool {
	return false
}
