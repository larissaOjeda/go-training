# Chapter 1 - A New Hope

## Installing go

The [installation](https://golang.org/dl/) is pretty simple and instructions are available for all major OS.
If you are a MacOS user you can use homebrew also.

## Installing an IDE

### [Visual Studio Code](https://code.visualstudio.com/)

This is a lightweight editor which is fast and support a lot of things either natively or via extensions. In order to make it speak go you have to install the following extension:

- [go extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

After installing the extension and reloading the editor you should head over to:

- View
- Command Palette
- Go: Install/Update tools
- Select all and click ok

after everything is installed restart the editor and you are ready.

Other useful extensions are:

- [gitlens](https://gitlens.amod.io/)
- [vs live share](https://marketplace.visualstudio.com/items?itemName=MS-vsliveshare.vsliveshare)

### [Goland (License needed)](https://www.jetbrains.com/go/)

This is a complete IDE which has a lot of nice features OOTB.
After installing it you have to setup some file watchers in order to trigger the formatting and import utility by doing as follows:

- File
- Settings
- Tools
- File watchers
- Click on the plus sign and add gofmt and  goimports
- Change the `Level` of each watcher to `Global` so it will stay activated for all future projects
- Click ok and restart IDE

### Vim

For the hardcore people amongst you there is also the option of vim.

You have to:

- [Install plugin manager](https://github.com/junegunn/vim-plug)
- [Plug vim-go](https://github.com/fatih/vim-go)
- [Example .vimrc file](https://github.com/fatih/dotfiles/blob/master/vimrc)
- Customize your .vimrc by binding the commands found in <https://github.com/fatih/vim-go/blob/master/doc/vim-go.txt> to the hotkeys of your preference.

### Debug support

All above options have debug support. The easiest one is Visual Studio Code and we will see in our first example the usage.

## Tools

### Linters

- [golangci-lint](https://github.com/golangci/golangci-lint)
- [golint](https://github.com/golang/lint)

## Your first program

We will create a small program in order to show how to use the go CLI and the debugger.

### Debugging

If running on macOS, make sure you have Xcode command line tools installed:

```bash
xcode-select --install
```

- Open the folder where the go training has been cloned with VS Code
- Open the `main.go` file under the following path `chapter1/src`
- Add a breakpoint on the `fmt.Printf` statement
- Hit F5

### Run from command line

```bash
go run main.go
```

The above builds and runs the code

### Build from command line

```bash
go build -o chapter1
```

The above builds the code as an executable named `chapter1`

### Install the executable in the go bin folder (only if project is in the GOPATH)

```bash
go install
```

This works only if your code resides inside the GOPATH.

## Appendix

### Add go binaries to your path

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Chapter 2**](../chapter2/README.md)
