package main

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func validCommand(lookup string) bool {
	switch lookup {
	case
		"add",
		"subtract":
		return true
	}
	return false
}

type actionFile struct {
	action string
	name   string
	file   *ini.File
}

func usage() {
	log.Fatal("Usage: inicom {basefile} [{add|subtract} file]...")
}

func loadIni(filename string) (*ini.File, error) {
	// wrapper to set common LoadOptions for loading the files
	return ini.LoadSources(ini.LoadOptions{AllowPythonMultilineValues: true}, filename)
}

func parse(args []string) []actionFile {
	var actionFiles []actionFile
	// use ini package to read in all files and actions with their related files
	// pattern should be: "action" "file", repeating
	if len(args)%2 != 0 {
		usage()
	}
	for i := 0; i < len(args); i += 2 {
		log.Printf("parsing args: %s, %s", args[i], args[i+1])
		if !validCommand(args[i]) {
			usage()
		}
		inifile, err := loadIni(args[i+1])
		if err != nil {
			log.Fatal(err.Error())
		}
		actionFiles = append(actionFiles, actionFile{args[i], args[i+1], inifile})
	}
	return actionFiles
}

func main() {
	args := os.Args[1:]
	log.Printf("\nArgs are %v\n", args)
	if len(args) < 1 {
		usage()
	}
	// acquire basefile
	basefile, err := loadIni(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	// convert remainder of args to actions+files
	actionFiles := parse(args[1:])
	log.Printf("basefile: %T", basefile)
	for _, af := range actionFiles {
		log.Printf("action: %s: %s", af.action, af.name)
	}

	//TODO manipulate ini per command in order. Apply changes to basefile

	basefile.WriteTo(os.Stdout)
	log.Println("Goodbye")
}
