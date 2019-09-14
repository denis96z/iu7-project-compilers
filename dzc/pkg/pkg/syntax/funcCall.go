package syntax

//easyjson:json
type FuncCall struct {
	Type   string       `json:"type"`
	Func   *Func        `json:"func"`
	Params []Expression `json:"params"`
}
