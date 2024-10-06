package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// get reference to the local project
	src := client.Host().Directory(".")

	// define build steps
	golang := client.Container().From("golang:1.17")
	golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

	// install dependencies
	golang = golang.WithExec([]string{"go", "mod", "download"})

	// run tests
	golang = golang.WithExec([]string{"go", "test", "./tests/unit/..."})
	golang = golang.WithExec([]string{"go", "test", "./tests/integration/..."})

	// build binary
	golang = golang.WithExec([]string{"go", "build", "-o", "tracesync", "."})

	// run security scan (example using gosec)
	golang = golang.WithExec([]string{"go", "install", "github.com/securego/gosec/v2/cmd/gosec@latest"})
	golang = golang.WithExec([]string{"gosec", "./..."})

	// get reference to build output directory in container
	output := golang.Directory("/src")

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, ".")
	if err != nil {
		return err
	}

	fmt.Println("Build completed successfully")
	return nil
}
