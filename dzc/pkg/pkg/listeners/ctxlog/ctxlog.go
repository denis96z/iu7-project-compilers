package ctxlog

import (
	"fmt"
	"log"

	"dzc/pkg/pkg/listeners/context"
)

func Fatalf(ctx context.Context, format string, v ...interface{}) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	log.Fatalf(fmt.Sprintf("[%d:%d] %s", line, column, format), v...)
}
