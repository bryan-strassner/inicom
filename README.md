# inicom
A Go implementation of a command line ini manipulator, allowing for adding and subtracting values from other ini files.

Rules:
- Never modify input files
- Write output to stdout

Commands:
```
$ inicom [--add={add_file}] [--subtract={sub_file}] {base_file}

  flags:
  {base_file} indicates the file to use as the anchor of this process.
  --add={add_file} indicates a file to layer over the top of the base file
  --subtract={sub_file} indicates a file of matching secitons+keys to remove from the base file.
            Values in this file are inconsequential.

  If both add and subtract are specified, subtract is evaluated last.  
```
Example:
```
./mydir
  |- basefile.ini
  |- addfile.ini

$ inicom --add=mydir/addfile.ini mydir/basefile.ini > newfile.ini
```

Backlog:
- TODO: initial version
- TODO: test samples
- TODO: empty sections pruning switch
- TODO: move to cobra
    - TODO: support multiple adds and substracts in one command
    - TODO: evaluate adds and subtracts in left-to-right order
    - TODO: allow for base_file to be specified out of order with flags
