package inicom

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
)

type ActionFile struct {
	Action string
	Name   string
	File   *ini.File
}

func validCommand(lookup string) bool {
	switch lookup {
	case
		"add",
		"subtract":
		return true
	}
	return false
}

func loadIni(filename string) (*ini.File, error) {
	// wrapper to set common LoadOptions for loading the files
	return ini.LoadSources(ini.LoadOptions{AllowPythonMultilineValues: true, Insensitive: true}, filename)
}

// Basefile : acquire the basefile as an ini File
func Basefile(filename string) (*ini.File, error) {
	return loadIni(filename)
}

// Parse : parse the inicom inputs and return the "ActionFile" structs
func Parse(args []string) ([]ActionFile, error) {
	var actionFiles []ActionFile
	// use ini package to read in all files and actions with their related files
	// pattern should be: "action" "file", repeating
	if len(args)%2 != 0 {
		return actionFiles, fmt.Errorf("the number of arguments following the base file must be an even number in the pattern <action> <object-of-action>")
	}
	for i := 0; i < len(args); i += 2 {
		log.Printf("parsing args: %s, %s", args[i], args[i+1])
		if !validCommand(args[i]) {
			return actionFiles, fmt.Errorf("invalid args: %s, %s", args[i], args[i+1])
		}
		inifile, err := loadIni(args[i+1])
		if err != nil {
			log.Fatal(err.Error())
		}
		actionFiles = append(actionFiles, ActionFile{args[i], args[i+1], inifile})
	}
	return actionFiles, nil
}

// Process : modify basefile with actions specified by actions
func Process(basefile *ini.File, actionFiles []ActionFile) {
	for _, af := range actionFiles {
		log.Printf("action: %s: %s", af.Action, af.Name)
		switch af.Action {
		case "add":
			add(basefile, af.File)
		case "subtract":
			subtract(basefile, af.File)
		}
	}
}

// add : modify basefile by adding the contents of addfile
func add(basefile *ini.File, addfile *ini.File) {
	for _, section := range addfile.Sections() {
		basesection, err := basefile.GetSection(section.Name())
		if err != nil {
			log.Printf("Adding section [%s]", section.Name())
			basesection, _ = basefile.NewSection(section.Name())
		}
		for _, key := range section.Keys() {
			if basesection.HasKey(key.Name()) {
				basekey, _ := basesection.GetKey(key.Name())
				log.Printf("Overriding key '%s' in section [%s]", key.Name(), section.Name())
				basekey.SetValue(key.String())
			} else {
				log.Printf("Adding key '%s' to section [%s]", key.Name(), section.Name())
				basesection.NewKey(key.Name(), key.String())
			}
		}
	}
}

// subtract : modify basefile by removing the keys specified in subfile
func subtract(basefile *ini.File, subfile *ini.File) {
	for _, section := range subfile.Sections() {
		basesection, err := basefile.GetSection(section.Name())
		if err != nil {
			log.Printf("Section [%s] not present, nothing to remove", section.Name())
		} else {
			// section exists, deal with keys
			for _, key := range section.Keys() {
				if basesection.HasKey(key.Name()) {
					log.Printf("Deleting key '%s' in section [%s]", key.Name(), section.Name())
					basesection.DeleteKey(key.Name())
				} else {
					log.Printf("Key '%s' doesn't exist in section [%s]", key.Name(), section.Name())
				}
			}
			if section.Name() != "default" && len(basesection.Keys()) == 0 {
				log.Printf("Section [%s] is now empty and is being removed.", section.Name())
				basefile.DeleteSection(section.Name())
			}
		}
	}
}
