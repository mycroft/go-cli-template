package main

import (
	"context"
)

type App struct{}

const (
	GO_VERSION = "1.22"
	LINT_IMAGE = "golangci/golangci-lint:v1.58.2"
)

// Simply build the project with go build
func (m *App) Build(ctx context.Context, source *Directory) (string, error) {
	return dag.Golang().
		GetContainer(GO_VERSION, source).
		WithExec([]string{"go", "build", "-o", ".", "./cmd/..."}).
		Stdout(ctx)
}

// Build & run tests
func (m *App) Test(ctx context.Context, source *Directory) (string, error) {
	return dag.Golang().
		GetContainer(GO_VERSION, source).
		WithExec([]string{"go", "build", "-o", ".", "./cmd/..."}).
		WithExec([]string{"go", "test", "./..."}).
		Stdout(ctx)
}

// Lint
func (m *App) Lint(ctx context.Context, source *Directory) (string, error) {
	return dag.Container().From(LINT_IMAGE).
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"golangci-lint", "run", "-v", "--timeout", "5m"}).
		Stdout(ctx)
}
