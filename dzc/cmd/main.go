package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"dzc/pkg/ast/complexparser"
	"dzc/pkg/ast/constparser"
	"dzc/pkg/ast/pkgparser"
	"dzc/pkg/ast/subparser"
	"dzc/pkg/ast/typeparser"
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

	pPkg := pkgparser.New()
	p := parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pPkg, p.Start())

	fmt.Println("PKG:", pPkg.Pkg.Name)

	ts.Seek(0)

	pConst := constparser.New()
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pConst, p.Start())
	pConst.FixIncomplete()

	fmt.Println("CONSTANTS:")
	for k, v := range pConst.Consts {
		fmt.Println(k, v.Type, v.Value)
	}

	ts.Seek(0)

	pTypes := typeparser.New(pConst.Consts)
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pTypes, p.Start())
	pTypes.FixIncomplete()
	pTypes.AddBasic()

	fmt.Println("TYPES:")
	for k, v := range pTypes.Types {
		fmt.Println(k, v.Text(), v.Base())
	}

	ts.Seek(0)

	pComplex := complexparser.New(pTypes.Types)
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pComplex, p.Start())

	fmt.Println("ENUMS:")
	for k, v := range pComplex.Enums {
		fmt.Println(k, v.Name, v.Options)
	}

	fmt.Println("STRUCTS:")
	for k, v := range pComplex.Structs {
		fmt.Println(k, v.Name, v.Attrs)
	}

	ts.Seek(0)

	pSub := subparser.New(pTypes.Types)
	p = parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pSub, p.Start())

	fmt.Println("FUNCTIONS:")
	for k, v := range pSub.Functions {
		fmt.Println(k, v.Name, v.Args, v.RetType)
	}

	fmt.Println("PROCEDURES:")
	for k, v := range pSub.Procedures {
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
