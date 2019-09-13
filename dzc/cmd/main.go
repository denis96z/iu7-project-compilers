package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"dzc/pkg/pkg"
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

	pkgInfo := pkg.NewParser().ParseSource(string(src))

	//dst, err := codegen.GenPkgCode(pkgInfo)
	//if err != nil {
	//	log.Fatal(err)
	//}

	dst, err := json.Marshal(pkgInfo)
	if err != nil {
		log.Fatal(err)
	}

	nameDst := fmt.Sprintf(
		"%s/%s.json", pathSrc, pkgInfo.Pkg.Name,
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

	_, err = dSrc.Write(dst)
	if err != nil {
		log.Fatal(err)
	}
}
