package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"dzc/pkg/ast/complexparser"
	"dzc/pkg/ast/constparser"
	"dzc/pkg/ast/pkginfo"
	"dzc/pkg/ast/pkgparser"
	"dzc/pkg/ast/subparser"
	"dzc/pkg/ast/typeparser"
	"dzc/pkg/codegen"
	"dzc/pkg/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("source filename is required")
	}

	nameSrc := os.Args[1]
	pathSrc := filepath.Dir(nameSrc)

	fSrc, err := os.Open(nameSrc)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := fSrc.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	src, err := ioutil.ReadAll(fSrc)
	if err != nil {
		log.Fatal(err)
	}

	//for {
	//	t := lx.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n",
	//		lx.SymbolicNames[t.GetTokenType()], t.GetText())
	//}



	pPkg := pkgparser.New()
	p := parser.NewDZParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(pPkg, p.Start())

	fmt.Println("PKG:", pPkg.Pkg.Name)



	pConst := constparser.New()


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

	pkg := &pkginfo.PkgInfo{
		Pkg:     pPkg.Pkg,
		Consts:  pConst.Consts,
		Types:   pTypes.Types,
		Enums:   pComplex.Enums,
		Structs: pComplex.Structs,
	}

	dst, err := codegen.GenPkgCode(pkg)
	if err != nil {
		log.Fatal(err)
	}

	nameDst := fmt.Sprintf(
		"%s/%s.c", pathSrc, pkg.Pkg.Name,
	)

	dSrc, err := os.Create(nameDst)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dSrc.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	_, err = dSrc.WriteString(dst)
	if err != nil {
		log.Fatal(err)
	}
}
