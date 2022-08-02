# Chapter 6 - Dependency Management and Tools

## Go Modules

In order to demonstrate go modules we will create a new folder `test` and place a `main.go` file in it with the following content:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Print("Hello World!")
}
```

### Initialize

```bash
go mod init "github.com/taxibeat/test"
```

### go.mod

```bash
module github.com/taxibeat/test

go 1.12
```

The above show the initial content of the `go.mod` file.
At the top the module name and then the version of go it was created.
At this point there are no dependencies.

### Adding a module

Let's add the following code in the `main.go` file:

```go
package main

import (
    "fmt"
    "os"

    "github.com/beatlabs/patron"
)

func main() {
    err := patron.Setup("test", "1.0")
    if err != nil {
        fmt.Printf("failed to set up logging: %v\n", err)
        os.Exit(1)
    }
}
```

In order to trigger the module installation you can call:

```bash
go mod tidy
```

The result of the above action is the alteration of the `go.mod` file:

```bash
module github.com/taxibeat/test

go 1.12

require github.com/beatlabs/patron v0.26.0
```

and the creation of the `go.sum` file which contains in more details the exact dependency graph along with version, if they exist, and some hashes.

### Upgrading a module

Let's assume that we want to update to the latest version:

```bash
go get github.com/beatlabs/patron
```

will update to the latest minor version

You can even select the specific version with:

```bash
go get github.com/beatlabs/patron@v0.25.0
```

> Since Go 1.16, installing executables with `go get` in module mode is deprecated,
> and is more focused on managing requirements in `go.mod`.
> To install using requirements of the current module, use `go install`.
>
> When used with a version suffix (like @latest or @v1.4.6), `go install` builds packages in
> module-aware mode, ignoring the `go.mod` file in the current directory or any parent directory, if
> there is one. This is useful for installing executables without affecting the dependencies of
> the main module.

Every time go mod is changed we have to make sure that everything is synced by calling `go mod tidy` and `go mod vendor` to sync if vendoring is used.

### Removing a module

In order to remove a module you have to:

- Remove all relevant code first
- Call `go mod tidy` which will handle the rest and `go mod vendor` to sync if vendoring is used.

### Tidy up

It is always good to have the dependencies in a consistent state which can be accomplished with:

```bash
go mod tidy
```

which adds missing and removes unused modules.

### IDE Integration

#### Visual Studio Code

Go Modules support in Visual Studio Code is a bit limited, you can see known [issues and progress](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code).
For now whenever you update a module make sure to restart Visual Studio Code so the language server will be restarted.

#### GoLand

Follow these steps to enable to Go Modules integration in GoLand:

1. Open GoLand settings
2. Go to Go -> Go Modules (vgo)
3. Make sure "Enable Go Modules (vgo) integration" is checked

### Vendoring

```bash
go mod vendor
```

This command will create a vendor directory and put any dependency in it.

Vendoring has the following advantages:

- go build can use vendor and thus do not need to go to the internet in order to fetch modules which speeds up the build process (CI/CD)
- every module upgrade shows exactly what changes in the vendor code files
- repeatable builds, as vendored files are part of the codebase

and the following disadvantages:

- Repository size will increase due to vendored files
- Commits and reviews are getting bigger
- Vendor has to be synced every time something changes

### Other

You can always make changes to the `go.mod` file by hand but you should then use `go mod tidy` and `go mod vendor` in order to sync your repo.

As a rule of thumb we should always, after a change in dependency, call:

- `go mod tidy`
- `go mod vendor`, if we are vendoring

To clean up the mod cache and force downloading dependencies again, call:

- `go clean -cache`

Check out [Resources - Further studying material](../resources/README.md).

## Linting

### [go vet](https://golang.org/cmd/vet/)

This is a tool that comes bundled in with the go installation.

### [lint](https://github.com/golang/lint)

It is installed as part of the VS Code go extensions.

### [golangci-lint](https://github.com/golangci/golangci-lint)

This is a meta-linter, meaning that it uses various linters underneath and reports the problems in an unified format.

[go vet](https://golang.org/cmd/vet/) and [golint](https://github.com/golang/lint) are included also.

## [Makefile](src/Makefile)

The included Makefile contains a lot collection of commands that are usually needed during development. When a sub-command is not provided the default one is `test`. Every command checks that the code is properly formatted with the help of the included script. The commands are:

### test (fmtcheck will always run before this)

```bash
go test ./... -cover -race -timeout 60s
```

This command runs tests

- in all packages
- reports coverage
- tests for race conditions
- times out after 60 seconds

### testint (fmtcheck will always run before this)

```bash
go test ./... -race -cover -tags=integration -timeout 60s -count=1
```

This commands runs tests

- in all packages
- reports coverage
- tests for race conditions
- times out after 60 seconds
- includes all files with `integration` build tag
- by providing the count we bust the test cache so the tests will not cached

### cover (fmtcheck will always run before this)

```bash
go test ./... -coverpkg=./... -coverprofile=cover.out -tags=integration -covermode=atomic && \
go tool cover -func=cover.out && \
rm cover.out
```

This multi-command report the code coverage including files with the `integration` build tag.

### fmt (fmtcheck will always run before this)

```bash
go fmt ./...
```

The command formats the code in all packages.

### fmtcheck

```bash
@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"
```

The command executes a format check and exits with error if the

### lint (fmtcheck will always run before this)

```bash
golangci-lint run -E golint --exclude-use-default=false --build-tags integration
```

We are using the `golangci-lint` meta-linter with enabled `golint` including files with the `integration` build tag.

### deeplint (fmtcheck will always run before this)

```bash
golangci-lint run --enable-all --exclude-use-default=false -D dupl --build-tags integration
```

We are using the `golangci-lint` meta-linter with all linters enabled, except the `dupl` linter, including files with the `integration` build tag.

### ci

The command is used in our CI pipeline and calls the following commands:

- fmtcheck
- lint
- testint

### modsync (fmtcheck will always run before this)

```bash
go mod tidy && \
go mod vendor
```

The command uses the go modules to do the following:

- using `go mod tidy` to check and update the all modules. (see module section)
- vendor all modules

## [Dockerization](src/Dockerfile)

The included [Dockerfile](src/Dockerfile) is an example of how to create a container.
It uses a multistage build approach where in the first stage:

- copies everything to the container
- set's default version, which can be overridden by the docker CLI in order to inject the version
- build using vendor and overrides the version variable in the main package, thus injecting the aforementioned version

and the second stage:

- uses `scratch` as the base container (for more advanced scenarios you might go for alpine)
- copies only the executable from the builder stage

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Chapter 7**](../chapter7/README.md)  
[<- Previous&nbsp;: **Chapter 5**](../chapter5/README.md)
