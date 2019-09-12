package ast

import (
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Parser struct {
	listeners []Listener
}

func New() *Parser {
	return &Parser{
		listeners: nil,
	}
}

type Listener interface {
	antlr.ParseTreeListener
	FixIncomplete()
}

func (p *Parser) ParseSource(src string) *pkginfo.PkgInfo {
	lx := parser.NewDZLexer(
		antlr.NewInputStream(string(src)),
	)

	_ = &pkginfo.PkgInfo{}
	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)

	for _, listener := range p.listeners {
		ts.Seek(0)

		ps := parser.NewDZParser(ts)
		antlr.ParseTreeWalkerDefault.Walk(listener, ps.Start())
		listener.FixIncomplete()
	}

	return nil
}
