# Intro to Go(lang) Notes

Introduction

Why Go?

Big names involved (Google, etc.)
V.1 release in 2012
Static programming language
Natively takes advantage of multiple core processors
Efficient Compilation (compiled programs, garbage collection, no VM)
Efficient execution
Ease of Programming
Good at creating web services at scale
Good at networking (http, tcp, udp), concurrency/parallelism, systems, automation, command-line tools, image processing(?)
Clean syntax, powerful standard library, portable

General Tips

Good Go-appropriate IDEs include VSC and Goland.
"package main" is the entry point of all go programs.

Go Commands:

go version = prints out the currently installed version of go
go env = prints out all of the environment variables
go help = prints out all available go commands
go fmt = automatically formats code.
Running "go fmt ./..." formats automatically formats everything in a repo.
go run your_file_here = builds and runs your code.
go build your_executable_here = builds the file, reports any errors, and puts an executable into the current folder.
go build your_package_here = builds the file, reports any errors, and throws away any binary.
go install your_executable_here = compiles the program, names it, and puts the executable in workspace/bin
go install your_package_here = compiles the package, puts the executable in workspace/pkg, and makes it an archive file.

Packages

Packages and package managers allow you to "re-use" other people's code
NPM is a great example.
It is important to know that packages are safe to use (no packages, etc.)
godoc.org has official and 3rd-party packages and documentation available for use.
Using other packages makes your software "dependent" on the packages, hence "dependency".
It is possible to both update dependencies AND specify which version to use.

Modules

Modules help with go package management.
Go Version 1.13 requires modules
"go mod init" creates a new module, initializing the go.mod file that describes it.
"go build", "go test", and other package-building commands add new dependencies to go.mod
"go -list -m all" prints the current module's dependencies
"go get" changes the required version of a dependency, or adds a dependency
"go mod tidy" removes unused dependencies

Overall need to understand modules more, but I think the main point is to understand the importance of thinking about how direct and indirect dependencies affect YOUR code.

Both go.mod and go.sum should be checked into version control.

blog.golang.org = the go blog
blog.golang.org/using-go-modules = good link to visit to learn about go modules.
