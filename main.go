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
