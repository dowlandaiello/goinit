# goinit

A simple utility that creates a new Go project in the GOPATH.

## Installation

```zsh
go install github.com/dowlandaiello/goinit
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
- Creates a new golang Dockerfile in the given directory that runs main.go (without any params).
- Creates a main.go file in the given directory.
- Creates a new golang .travis.yml in the given directory.
- Creates a new golang .gitignore in the given directory.
