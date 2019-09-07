package syntax

//easyjson:json
type FuncReturn struct {
	Type  string     `json:"type"`
	Value Expression `json:"value"`
}

func NewFuncReturn(value Expression) *FuncReturn {
	return &FuncReturn{
		Type:  StatementFuncReturn,
		Value: value,
	}
}

func (v FuncReturn) GetType() string {
	return v.Type
}

func (v FuncReturn) IsCondition() bool {
	return false
}

func (v FuncReturn) IsLoop() bool {
	return false
}

func (v FuncReturn) IsTrueLoop() bool {
	return false
}

func (v FuncReturn) IsWhileLoop() bool {
	return false
}

func (v FuncReturn) IsBreak() bool {
	return false
}

func (v FuncReturn) IsContinue() bool {
	return false
}

func (v FuncReturn) IsProcCall() bool {
	return false
}

func (v FuncReturn) IsVarDecl() bool {
	return false
}

func (v FuncReturn) IsExpression() bool {
	return false
}

func (v FuncReturn) IsReturn() bool {
	return true
}

func (v FuncReturn) IsProcReturn() bool {
	return false
}

func (v FuncReturn) IsFuncReturn() bool {
	return true
}
