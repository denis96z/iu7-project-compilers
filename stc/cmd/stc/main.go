package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"stc/pkg/parser"

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

	lx := parser.NewSmallTalkLexer(
		antlr.NewInputStream(string(src)),
	)
	ts := antlr.NewCommonTokenStream(
		lx, antlr.TokenDefaultChannel,
	)

	lst := parser.NewListener()
	defer func() {
		//if r := recover(); r != nil {
		//	log.Println(r)
		//}

		dst := lst.GetOutput()
		nameDst := fmt.Sprintf("%s/out.pir", pathSrc)

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
	}()

	ps := parser.NewSmallTalkParser(ts)
	antlr.ParseTreeWalkerDefault.Walk(lst, ps.Script())
}
