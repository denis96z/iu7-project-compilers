package context

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Context interface {
	GetStart() antlr.Token
}
