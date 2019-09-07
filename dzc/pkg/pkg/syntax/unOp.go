package syntax

//easyjson:json
type UnaryOperation struct {
	Operand  Expression `json:"operand"`
	Operator *Operator  `json:"operator"`
}
