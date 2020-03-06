package main

import (
	"fmt"
    "os"

    "gopkg.in/ini.v1"
)

const (
	add = "add" // add file indicator - file will be added to the current state
	subtract = "subtract" // subtract file indicator - file will be subtracted from the current state
)

type actionFile struct {
	action string
    file *ini.File 
}

func errExit(msg string) {
	// print the message to stderr and exit(1)
	os.Stderr.WriteString(fmt.Sprintf("\n%s\n", msg))
	os.Exit(1)
}

func usage() {
	errExit("Usage: inicom {basefile} [{add|subtract} file]...")
}

func parse(args []string) ([]actionFile) {
	var actionFiles []actionFile
	// use ini package to read in all files and actions with their related files
	// pattern should be: "action" "file", repeating
	if len(args) % 2 != 0 {
		usage()
	}
	for i := 0; i < len(args); i += 2 {
		os.Stderr.WriteString(fmt.Sprintf("\nparsing args: %s, %s", args[i], args[i+1]))
		inifile, err := ini.Load(args[i+1])
		if err != nil {
			errExit(err.Error())
		}
		actionFiles = append(actionFiles, actionFile{args[i], inifile})
	}
	return actionFiles
}

func main() {
	args := os.Args[1:]
	// debug. remove?
	os.Stderr.WriteString(fmt.Sprintf("Args are %v", args))
	if len(args) < 1 {
		usage()
	}
	// acquire basefile
	basefile, err := ini.Load(args[0])
	if err != nil {
		errExit(err.Error())
	}

	// convert remainder of args to actions+files
	actionFiles := parse(args[1:])

	//TODO manipulate ini per command
	//TODO write output to stdout
	os.Stderr.WriteString(fmt.Sprintf("\nbasefile: %T, actionFiles: %d", basefile, len(actionFiles)))

	os.Stderr.WriteString("\nGoodbye\n")


}
