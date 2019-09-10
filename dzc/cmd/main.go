package main

import (
	"dzc/pkg/ast/constparser"
	"dzc/pkg/ast/subparser"
	"dzc/pkg/ast/typeparser"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	cp := constparser.NewParser()
	p := parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(cp, p.Start())
	cp.FixDefinitions()

	fmt.Println("CONSTANTS:")
	for k, v := range cp.KnownConsts {
		fmt.Println(k, v.Type, v.Value)
	}

	ts.Seek(0)

	tp := typeparser.NewParser(cp.KnownConsts)
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(tp, p.Start())
	tp.FixIncomplete()
	tp.AddBasic()

	fmt.Println("TYPES:")
	for k, v := range tp.Types {
		fmt.Println(k, v.Text(), v.Base())
	}

	ts.Seek(0)

	sp := subparser.New(tp.Types)
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(sp, p.Start())

	fmt.Println("FUNCTIONS:")
	for k, v := range sp.Functions {
		fmt.Println(k, v.Name, v.Args, v.RetType)
	}

	fmt.Println("PROCEDURES:")
	for k, v := range sp.Procedures {
		fmt.Println(k, v.Name, v.Args)
	}

	//
	//globalParser := ast.NewGlobalParser()
	//
	//p := parser.NewDZParser(ts)
	//
	//
	//fmt.Println(
	//	globalParser.PkgInfo.Pkg,
	//	globalParser.PkgInfo.Procedures,
	//	globalParser.PkgInfo.Functions,
	//)
	//
	//fmt.Println("TYPES:")
	//for k, v := range globalParser.PkgInfo.types {
	//	base := v.Base()
	//	fmt.Println(k, base)
	//}

	//ts.Seek(0)
	//antlr.ParseTreeWalkerDefault.Walk(&ast.GlobalParser{}, p.Start())
}
