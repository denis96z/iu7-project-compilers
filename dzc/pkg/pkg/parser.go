package pkg

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/listeners/blocks"
	"dzc/pkg/pkg/listeners/consts"
	"dzc/pkg/pkg/listeners/pkgs"
	"dzc/pkg/pkg/listeners/subs"
	"dzc/pkg/pkg/listeners/types"
	"dzc/pkg/pkg/pkg"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Parser struct {
	listeners []Listener
}

func NewParser() *Parser {
	return &Parser{
		listeners: []Listener{
			pkgs.New(),
			consts.New(),
			types.New(),
			subs.New(),
			blocks.New(),
		},
	}
}

type Listener interface {
	antlr.ParseTreeListener

	Initialize(info *pkg.Info)
	Finalize()
}

func (p *Parser) ParseSource(src string) *pkg.Info {
	lx := parser.NewDZLexer(
		antlr.NewInputStream(src),
	)

	//for {
	//	t := lx.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lx.SymbolicNames[t.GetTokenType()], t.GetText())
	//}
	//os.Exit(1)

	info := &pkg.Info{}
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)

	for _, listener := range p.listeners {
		ts.Seek(0)

		listener.Initialize(info)

		ps := parser.NewDZParser(ts)
		antlr.ParseTreeWalkerDefault.Walk(listener, ps.Start())

		listener.Finalize()
	}

	return info
}
