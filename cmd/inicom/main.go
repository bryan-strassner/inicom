package main

import (
	"log"
	"os"

	"github.com/bryan-strassner/inicom/internal/inicom"
)

func usage() {
	log.Fatal("Usage: inicom {basefile} [{add|subtract} file]...")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}
	// acquire basefile
	basefile, err := inicom.Basefile(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	// convert remainder of args to actions+files
	actionFiles, err := inicom.Parse(args[1:])
	if err != nil {
		log.Fatal(err.Error())
	}

	inicom.Process(basefile, actionFiles)
	if err != nil {
		log.Fatal(err.Error())
	}

	basefile.WriteTo(os.Stdout)
}
