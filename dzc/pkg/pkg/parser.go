package pkg

import (
	"dzc/pkg/parser"
	"dzc/pkg/pkg/listeners/consts"
	"dzc/pkg/pkg/listeners/pkgs"
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
		},
	}
}

type Listener interface {
	antlr.ParseTreeListener

	SetPkg(info *pkg.Info)
	FixIncomplete()
	UpdatePkg()
}

func (p *Parser) ParseSource(src string) *pkg.Info {
	lx := parser.NewDZLexer(
		antlr.NewInputStream(src),
	)

	info := &pkg.Info{}
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)

	for _, listener := range p.listeners {
		ts.Seek(0)

		listener.SetPkg(info)

		ps := parser.NewDZParser(ts)
		antlr.ParseTreeWalkerDefault.Walk(listener, ps.Start())

		listener.FixIncomplete()
		listener.UpdatePkg()
	}

	return info
}
