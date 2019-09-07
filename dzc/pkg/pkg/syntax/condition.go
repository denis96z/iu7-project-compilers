package syntax

//easyjson:json
type Cond struct {
	Type string               `json:"type"`
	If   *IfConditionPart     `json:"if"`
	ElIf []*ElIfConditionPart `json:"elif"`
	Else *ElseConditionPart   `json:"else"`
}

func NewCondition(
	ifPart *IfConditionPart,
	elifParts []*ElIfConditionPart,
	elsePart *ElseConditionPart,
) *Cond {
	return &Cond{
		Type: StatementCondition,
		If:   ifPart,
		ElIf: elifParts,
		Else: elsePart,
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

func (c Cond) IsProcReturn() bool {
	return false
}

func (c Cond) IsFuncReturn() bool {
	return false
}

type IfConditionPart struct {
	Cond Expression `json:"condition"`
	Body *Block     `json:"body"`
}

type ElIfConditionPart struct {
	Cond Expression `json:"condition"`
	Body *Block     `json:"body"`
}

type ElseConditionPart struct {
	Body *Block `json:"body"`
}
