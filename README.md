# github.com/rwg/golang-dl

This repository is a fork of the golang.org/dl Go package, which contains
wrapper programs that download and run specific versions of Go. This allows
multiple versions of the Go toolchain to be installed side-by-side and run
using version-specific `go<version>` commands.

The upstream repository is available at:

- https://cs.opensource.google/go/dl
- https://github.com/golang/dl

## Changes from upstream

- Exit statuses from `go` commands that the wrapper programs run are now used
  as the wrapper programs' exit statuses. (The upstream version only used exit
  statuses 0 and 1.)

## Usage

Use `go install` to install the wrapper program from this repository for the
version of Go you want to install and use. In this example, we'll install the
wrapper for Go version 1.21.0:

```sh
go install github.com/rwg/golang-dl/go1.21.0@latest
```

This creates a `go1.21.0` wrapper program in your `$GOPATH/bin` directory.

Now use the wrapper program to download the corresponding version of the Go
toolchain:

```sh
go1.21.0 download
```

Now that the Go 1.21.0 toolchain has been downloaded, the `go1.21.0` wrapper
program can be used anywhere you'd normally use the `go` command when you want
to force the use of Go 1.21.0. For example, this command will print the path
where the Go 1.21.0 toolchain installed by the wrapper was installed:

```sh
go1.21.0 env GOROOT
```

To remove a specific Go toolchain version, delete the directory returned by
`go<version> env GOROOT` and the `$GOPATH/bin/go<version>` wrapper program.
