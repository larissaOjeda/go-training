# Chapter 2 - The looks

## Visibility

Go has a simple model for access control on function, methods, fields, etc.

If the name start with a capital letter then it is accessible from the outside (`exported`) else it is not (`unexported`). There is no other access level modifier available e.g. `internal`, `protected`.

Proper design is essential in order to work with this simple model.

## Documentation

Documentation is 1st class citizen in go. The default installation up to go version 1.12 contains a tool that shows the documentation of your current project via CLI or Web.

```bash
godoc -http=:6060
```

After the upgrade to go 1.13 you have to manually download the utility via:

```bash
go get golang.org/x/tools/cmd/godoc
godoc
```

The primary convention here is to comment anything that is exported because this is the public api of the package. You can see some details following the [link](https://golang.org/doc/effective_go.html#commentary).

The format of this comment should be as follows:

```go
// Sum calculates the summary of two integers.
// This is a second line comment.
func Sum(a, b int) int {
    return a + b
}
```

The comment always starts with a space and the name of the function and continues like reading a sentence. It always ends with a punctuation.

You can see more details following the [link](https://golang.org/doc/effective_go.html#commentary).

## Naming conventions

### Packages

- short, prefer transport over transportmechanism
- clear, like logging, postgres
- singular, e.g. `user` and not the plural `users`
- avoid catchall packages like utils, helpers, models
- since the package name is part of the declaration of a type we should use it e.g. a package named `cache` with a constructor `NewCache()` can be rename to just `New()` since the usage will always contain the package name `cache.New()`

### Variables

- Use `MixedCaps` or `mixedCaps` rather than underscores to write multi-word names.
- Abbreviations should always be capitalized e.g. `ServerHTTP`.
- single letter for indices (i, j, k in for loops)
- short names like "cust" for Customer or even "c" are perfectly valid as long as the declaration and its usages is very close. `The greater the distance between a name's declaration and its uses, the longer the name should be.`
- use repeated letters to represent a slice/array and use single letter in loops

```go
var tt []Thing

for i, t := range tt {
    ...
}
```

### Functions and methods

- Avoid repeating the package name in name of function and methods e.g. prefer `log.Info()` than `log.LogInfo()`.
- Go does not have getters and setters so the convention given a unexported field named `age` is `Age()` for the getter and `SetAge(age int)` for the setter.

### Interfaces

The name of the interface should be whatever the function is plus a "er" at the end

```go
type Reader interface {
    Read() ([]byte,error)
}
```

The above does not always make sense so try, but do not force it

```go
type Repository interface {
    Customer() (*Customer, error)
    Save(c *Customer) error
    Delete(id int) error
}
```

When embedding interfaces we should concatenate the name e.g. given a `Reader` and a `Closer` interface, the composition would be `ReadCloser`.

## Package structure

There is unfortunately no standard about structuring the packages. There are some guidelines like:

- [Standard Package Layout by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1), example [hasura](https://github.com/taxibeat/hasura)
- [Code like the Go team](https://www.youtube.com/watch?v=MzTcsI6tn-0)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout), example [REMS](https://github.com/taxibeat/rems), following the Ports and Adapters Architecture

Some of the above guidelines might be outdated or are not endorsed anymore.

## Project Structure

All project contain a lot of files for different purposes:

- Go files for our code
- Dockerfile
- docker-compose file
- CI files (Travis/CircleCI/Jenkins)
- Deployment files (Jenkins)
- Makefile
- Readme.md (root directory)
- .gitignore file
- go.mod, go.sum
- various scripts
- vendor
- Helm charts
- Infrastructure (terraform)
- Monitoring files (Grafana, Prometheus alerts)
- etc.

As you might see the list is big and we need to organize them a lot better in order to make it more maintainable.

We might organize our projects following the below structure (project named `myservice`):

- [myservice](#root-myservice)
  - [api [OPT]](#api)
  - [cmd [OPT]](#cmd)
  - [config [OPT]](#config)
  - [internal [OPT]](#internal)
  - [infra [REQ]](#infra)
    - [build [REQ]](#build)
      - [Dockerfile [REQ]](#dockerfile)
      - [Jenkinsfile.ci [OPT]](#jenkinsfile-ci)
    - [deploy [REQ]](#deploy)
      - [local [OPT]](#local)
      - [helm [OPT]](#helm)
      - [terraform [OPT]](#terraform)
      - [Jenkinsfile [REQ]](#Jenkinsfile-cd)
    - [observe [OPT]](#observe)
      - [dashboard [OPT]](#dashboard)
      - [alerting [OPT]](#alerting)  
  - [doc [REQ]](#doc)
  - [example [OPT]](#example)
  - [script [OPT]](#script)
  - [test [OPT]](#test)
  - [tool [OPT]](#tool)
  - [vendor [REQ]](#vendor)
  - [.travis.yml [OPT]](#travis)
  - [go.mod and go.sum [REQ]](#go-modules)
  - [Makefile [OPT]](#makefile)
  - [README.md [REQ]](#readme-file)
  - [.gitignore [REQ]](#gitignore)

Some of the items are optional [OPT] or required [REQ]. The above layout is heavily influenced by [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Conventions here are the following:

- folder/package names are in singular format
- folder/package names should be short
- all go packages in the root folder can be imported by others projects
- all packages in the `internal` folder are private to the repository and cannot be imported by other projects

### root (myservice)

The following files are contained, along with the other folders:

### api

OpenAPI/Swagger specs, JSON schema files, protocol definition files.

### cmd

Main entrypoints for this project.  
The directory name for each entrypoint should match the name of the executable you want to have e.g. `/cmd/myservice/main.go`.

### config

Configuration file templates or default configs e.g. `env` files.

### internal

Private application and library code that should not be imported by other projects.  

### infra

This folder groups everything infra related together.

#### build

Packaging and Continuous Integration. This folder should contains

##### Dockerfile

The standard dockerfile to create the deployment artifact of the service.

##### Jenkinsfile CI

If the project uses Jenkins as a CI server the ci file should be in here.

#### deploy

Group deployment related files and folders.  
The root folder should contain:

- `Jenkinsfile`

##### local

Local deployment setup e.g. `docker-compose`. Local Kubernetes will deprecate this.

##### helm

Helm packages for the deployment. Each deployable unit should have a sub-folder.

##### terraform

Terraform files for setting up infrastructure.

##### Jenkinsfile CD

The Jenkinsfile responsible for deployment of the service.

#### observe

Artifacts needed to observe our service.

##### dashboard

Dashboards for Grafana.

##### alerting

Alerts for Prometheus AlertManager.

### doc

Design and user documents (in addition to your godoc generated documentation).  
Things that could be in there are:

- OpenAPI documentation used by [Hypatia](https://github.com/taxibeat/hypatia)
- Architecture diagrams
- Run-books
- etc.

### example

Examples for your applications and/or public libraries.

### script

Contains all scripts of the project.

### test

Additional external test apps and test data. Functional tests should live here.

### tool

Supporting tools for this project.

### vendor

Application dependencies.

#### travis

Unfortunately, it has to be in the root.

#### go modules

These files are generated by the go modules tool.

#### Makefile

Makefile is explained in [chapter 5](../chapter5/README.md#makefile).

#### README file

Readme contains entry-point of the service documentation, which will be displayed in Github.

#### gitignore

The standard .gitignore file.

## Formatting

Code formatting in go is actually a pretty simple thing. There is a tool that does almost all the formatting.
No need for policies, no need for holy wars.

One of the Go proverbs is:

**Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.**

The only thing you have to do is to call:

```bash
go fmt ./...
```

and every go file in every package will be formated with the only style that actually matters.
Visual Studio Code will setup this formatting automatically on each save and Goland has to be setup with file watchers.

The only thing that *gofmt* does not handle is line wrapping. Go has no line length limit. If a line feels too long, wrap it and indent with an extra tab.

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Chapter 3**](../chapter3/README.md)  
[<- Previous&nbsp;: **Chapter 1**](../chapter1/README.md)
