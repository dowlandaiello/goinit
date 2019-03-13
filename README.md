# goinit

A simple utility that creates a new Go project in the GOPATH.

## Installation

```zsh
go get github.com/dowlandaiello/goinit
```

## Usage

```zsh
goinit github.com/username/package
```

## What it Does

- Creates a new folder with the given package name in the gopath (under the respective username).
- Creates a go.mod file in the given package.
- Initializes a git repository in the directory.
- Sets the remote to "https://github.com/user/package.git".