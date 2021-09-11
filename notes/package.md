## Package

### $GOPATH
The $GOPATH environment variable lists places for Go to look for Go Workspaces.
By default, Go assumes our GOPATH location is at $HOME/go

### $GOROOT
The $GOROOT is where Go’s code, compiler, and tooling lives — this is not our source code.
The $GOROOT is usually something like /usr/local/go. Our $GOPATH is usually something like $HOME/go.

### Naming
Generally, the package name should match the name of the directory it’s kept in, but the main package 
is an exception to that rule. Because main package used as an entry point to the program.

go files in the same package AKA folder must have the same package name. 
The only exception are tests that can have a _test extension. The rule is 1 folder, 
1 package therefore the same package name for all go files in the same folder.

* The Gotool use the name in the import path as the name of the directory to load the pkg source code from.

```go
import "foo"
```
If it could not find the `foo` directory in the $GOPATH, then the program won't compile.

### Resources
* https://www.digitalocean.com/community/tutorials/importing-packages-in-go
* https://stackoverflow.com/a/40868503/2967670