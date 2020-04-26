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
	basefile, err := inicom.LoadIni(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	// convert remainder of args to actions+files
	actionFiles, err := inicom.Parse(args[1:])
	if err != nil {
		log.Println(err)
		usage()
	}
	for _, af := range actionFiles {
		log.Printf("action: %s: %s", af.Action, af.Name)
		switch af.Action {
		case "add":
			inicom.Add(basefile, af.File)
		case "subtract":
			inicom.Subtract(basefile, af.File)
		}
	}
	basefile.WriteTo(os.Stdout)
}
