package context

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Context interface {
	GetStart() antlr.Token
}

func CtxLogErrf(ctx Context, format string, args ...interface{}) {
	log.Fatalf(fmt.Sprintf("%s %s", ctxParseInfo(ctx), format), args...)
}

func ctxParseInfo(ctx Context) string {
	return fmt.Sprintf("[%d:%d]",
		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}
