package main

import (
	"log"
	"os"

	"github.com/bryan-strassner/inicom"
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
	actionFiles := inicom.Parse(args[1:])
	for _, af := range actionFiles {
		log.Printf("action: %s: %s", af.action, af.name)
		switch af.action {
		case "add":
			inicom.Add(basefile, af.file)
		case "subtract":
			inicom.Subtract(basefile, af.file)
		}
	}
	basefile.WriteTo(os.Stdout)
}
