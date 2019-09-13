package syntax

//go:generate easyjson

//easyjson:json
type Pkg struct {
	Name string `json:"name"`
}
