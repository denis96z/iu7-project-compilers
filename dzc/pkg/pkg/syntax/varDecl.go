package syntax

//easyjson:json
type VarDecl struct {
	Type string `json:"type"`
	Var  *Var   `json:"var"`
}

func NewVarDecl(v *Var) *VarDecl {
	return &VarDecl{
		Type: StatementVarDecl,
		Var:  v,
	}
}

func (v VarDecl) GetType() string {
	return v.Type
}

func (v VarDecl) IsCondition() bool {
	return false
}

func (v VarDecl) IsLoop() bool {
	return false
}

func (v VarDecl) IsTrueLoop() bool {
	return false
}

func (v VarDecl) IsWhileLoop() bool {
	return false
}

func (v VarDecl) IsBreak() bool {
	return false
}

func (v VarDecl) IsContinue() bool {
	return false
}

func (v VarDecl) IsProcCall() bool {
	return false
}

func (v VarDecl) IsVarDecl() bool {
	return true
}

func (v VarDecl) IsExpression() bool {
	return false
}

func (v VarDecl) IsReturn() bool {
	return false
}

func (v VarDecl) IsProcReturn() bool {
	return false
}

func (v VarDecl) IsFuncReturn() bool {
	return false
}
