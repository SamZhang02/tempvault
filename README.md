# tempvault

[![asciicast](https://asciinema.org/a/667256.svg)](https://asciinema.org/a/667256)

*tempvault* is a command line tool based on [fzf](https://github.com/junegunn/fzf) to quickly access, preview and paste template files, such as your project configuration files, latex templates, or starter code files.

> [!Warning]
> This tool is a currently a work in progress with many breaking changes and have little documentation for now.

## Installation
If you would like to build tempvault from scratch, you need [Go](https://go.dev/).
You can build a binary using

```shell
go build -o tempvault
```

and put the binary in your PATH.

I also have included an installation script to run build and setups for the tool to work right away, simply run

```shell
sh install.sh
```

**Disclaimer**: I personally use MacOS, the script may not work well for all other OS.

For now, there are not ready-to-use binaries yet.

## Usage
tempvault currently support two commands:

`tempvault add [file]` adds a files to the vault.

`tempvault browse` browses and searches the vault, allowing you to select files to paste into the cwd.

With enough interest, I am open to look into refining the tool with more functionalities and customizability. 

Personally, I use the alias below
```shell
alias tv="tempvault"
alias tva="tempvault add"
alias tvb="tempvault browse"
```
for a fast workflow with the tool.

## Contributing
All contributions are welcome in the form of issues or pull request.

The build system uses simple scripts in the `justfile`, and the code mainly follows the [cobra](https://github.com/spf13/cobra) structure.
