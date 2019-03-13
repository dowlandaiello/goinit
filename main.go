package main

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
)

// main is the main goinit function.
func main() {
	modulePath := os.Args[1] // Get module path

	goPath := os.Getenv("GOPATH") // Get gopath

	if goPath == "" { // Check could not fetch gopath
		goPath = build.Default.GOPATH // Set gopath
	}

	err := os.MkdirAll(filepath.FromSlash(fmt.Sprintf("%s/src/%s", goPath, modulePath)), os.ModePerm) // Make project dir

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	moduleFile, err := os.Create(filepath.FromSlash(fmt.Sprintf("%s/src/%s/go.mod", goPath, modulePath))) // Create module file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	_, err = moduleFile.WriteString(fmt.Sprintf("module %s\n", modulePath)) // Write module file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	mainFile, err := os.Create(filepath.FromSlash(fmt.Sprintf("%s/src/%s/main.go", goPath, modulePath))) // Create main.go

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	_, err = mainFile.WriteString("// package main is the main package.\npackage main\n\n// main is the main function.\nfunc main() {\n}") // Write module file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	travisFile, err := os.Create(filepath.FromSlash(fmt.Sprintf("%s/src/%s/.travis.yml", goPath, modulePath))) // Create travis

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	_, err = travisFile.WriteString("language: go\n\ngo:\n  - 1.12\n  - master\n\ninstall: true\n\nsudo: false\n\nmatrix:\n  allow_failures:\n    - go: master\n  fast_finish: true\n\nnotifications:\n  email: false\n\nbefore_install:\n  - export GO111MODULE=on # Enable go mod\n  - go mod vendor # Download deps\n\nbefore_script:\n  - GO_FILES=$(find . -iname '*.go' -type f | grep -v -r ./vendor/)\n  - go get github.com/mattn/goveralls\n  - go get golang.org/x/tools/cmd/cover\n  - go get -u golang.org/x/lint/golint\n  - go get github.com/fzip/gocyclo\n\nscript:\n  - test -z $(gofmt -s -l $GO_FILES)\n  - go vet ./...\n  - go test ./...\n  - go build") // Write module file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	dockerFile, err := os.Create(filepath.FromSlash(fmt.Sprintf("%s/src/%s/Dockerfile", goPath, modulePath))) // Create dockerfile

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	_, err = dockerFile.WriteString(fmt.Sprintf("FROM golang:1.12\n\nWORKDIR /go/src/%s\nCOPY . .\n\nRUN go get -d -v ./...\n\nCMD go run main.go", modulePath)) // Write docker file

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	cmd := exec.Command("git", "init") // Initialize git repository

	cmd.Dir = filepath.FromSlash(fmt.Sprintf("%s/src/%s", goPath, modulePath)) // Set CMD dir

	_, err = cmd.Output() // Get output

	if err != nil { // Check for errors
		panic(err) // Panic
	}

	cmd = exec.Command("git", "remote", "add", "origin", fmt.Sprintf("https://github.com/%s.git", modulePath)) // Set git remote

	cmd.Dir = filepath.FromSlash(fmt.Sprintf("%s/src/%s", goPath, modulePath)) // Set CMD dir

	cmd.Output() // Run
}
