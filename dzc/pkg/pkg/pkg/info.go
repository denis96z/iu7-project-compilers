package pkg

import (
	"dzc/pkg/pkg/syntax"
)

//easyjson:json
type Info struct {
	Pkg        *syntax.Pkg                  `json:"pkg"`
	Consts     map[string]*syntax.Const     `json:"consts"`
	Types      map[string]syntax.Type       `json:"types"`
	Procedures map[string]*syntax.Procedure `json:"procedures"`
	Functions  map[string]*syntax.Function  `json:"functions"`
}
