# inicom
A Go implementation of a command line ini manipulator, allowing for adding and subtracting values from other ini files.

Rules:
- Never modify input files
- Write output to stdout
- Try to conform to https://github.com/golang-standards/project-layout

Commands:
```
$ inicom {base_file} [{add | subtract} {file} ...]

  Arguments:
  {base_file} indicates the file to use as the anchor of this process.
  [{add | subtract} {file}] As many repititions as desired of the verb 'add' or 'subtract' to indicate the action taken with the subsequent file
    add indicates to layer over the top of the base file
    subtract indicates to remove matching sections+keys from the base file. Values in this file are inconsequential.
    repeated adds an subtracts are evaluated in order and operate on the result produced by any prior actions
  In all cases of valid input, the output of this command will be a normalized version.
```
Example:
```
./mydir
  |- addfile.ini
  |- add2file.ini
  |- basefile.ini
  |- subfile.ini

$ inicom basefile.ini add addfile.ini subtract subfile.ini add add2file.ini > newfile.ini
```



Backlog:
- TODO: initial version
- TODO: unit tests
- TODO: Dockerfile build
- TODO: publish to a docker repository
- TODO: automation of builds
