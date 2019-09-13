package syntax

//easyjson:json
type FuncCall struct {
	Type   string       `json:"type"`
	Func   *Function    `json:"func"`
	Params []Expression `json:"params"`
}
