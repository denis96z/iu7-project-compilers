package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"dzc/pkg/ast"
	"dzc/pkg/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("source filename is required")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	src, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lx := parser.NewDZLexer(
		antlr.NewInputStream(string(src)),
	)

	//for {
	//	t := lx.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lx.SymbolicNames[t.GetTokenType()], t.GetText())
	//}

	ts := antlr.NewCommonTokenStream(lx, antlr.TokenDefaultChannel)

	globalParser := ast.NewGlobalParser()

	p := parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(globalParser, p.Start())

	fmt.Println(globalParser.Declarations)

	//ts.Seek(0)
	//antlr.ParseTreeWalkerDefault.Walk(&ast.GlobalParser{}, p.Start())
}
