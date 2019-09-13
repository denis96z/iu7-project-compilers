package syntax

type Operator struct {
	Name string `json:"operator"`
}

const (
	OperatorAdd = "+"
	OperatorSub = "-"
	OperatorMul = "*"
	OperatorDev = "/"
	OperatorMod = "%"
)
