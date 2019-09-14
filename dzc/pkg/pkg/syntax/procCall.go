package syntax

//easyjson:json
type ProcCall struct {
	Type   string                `json:"type"`
	Proc   *Proc                 `json:"proc"`
	Params map[string]*CallParam `json:"params"`
}

func NewProcCall(proc *Proc, params map[string]*CallParam) *ProcCall {
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

func (v ProcCall) IsProcReturn() bool {
	return false
}

func (v ProcCall) IsFuncReturn() bool {
	return false
}
