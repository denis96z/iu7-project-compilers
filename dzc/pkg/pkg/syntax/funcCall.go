package syntax

//easyjson:json
type FuncCall struct {
	Func   *Func                 `json:"func"`
	Params map[string]*CallParam `json:"params"`
}

func (v FuncCall) GetValueType() Type {
	return v.Func.RetType
}
