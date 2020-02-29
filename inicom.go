package main

import (
	"flag"
	"fmt"
    "os"

    "gopkg.in/ini.v1"
)

func main() {
	var AddFile string
	flag.StringVar(&AddFile, "add", "", "A file to add to the base file")
	
	var SubFile string
	flag.StringVar(&SubFile, "subtract", "", "A file to subtract from the base file")
	
	flag.Parse()

	args := flag.Args()
	
	//TODO remove the following placeholder output when the code does anything
	fmt.Fprintf(os.Stderr, "input dump: AddFile = %s, SubFile = %s, args = %v", AddFile, SubFile, args);

	//TODO use ini package to read files
	//TODO manipulate ini per command
	//TODO write output to stdout


	os.Stderr.WriteString("Goodbye")


}
