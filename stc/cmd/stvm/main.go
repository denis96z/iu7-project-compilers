package main

import (
	"io/ioutil"
	"log"
	"os"

	"stc/pkg/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("source filename is required")
	}

	nameSrc := os.Args[1]

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

	script := exec.NewScript()
	script.Execute(string(src))
}
