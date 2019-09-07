package syntax

type CallParam struct {
	Name  string     `json:"name"`
	Value Expression `json:"value"`
}
