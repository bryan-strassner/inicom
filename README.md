[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/bryan-strassner/inicom) 

# inicom
A Go implementation of a command line ini manipulator, allowing for merging and subtracting values

Rule 1:  Never modify input files
Rule 2:  Write output to stdout

Commands:

inicom merge file1 file2 > newfile

  Using file1 as a base, add values from file2. File2 values take precedence.

inicom subtract file1 file2 > newfile

  Using file1 as a base, remove all keys indicated in file2 (ignoring values).
  TODO: empty sections pruning switch?
