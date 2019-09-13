package syntax

//easyjson:json
type BinOperation struct {
	Operator     *Operator  `json:"operator"`
	LeftOperand  Expression `json:"left_operand"`
	RightOperand Expression `json:"left_operand"`
}
