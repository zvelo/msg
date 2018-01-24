// +build mage

package main

import (
	"context"
	fmt "fmt"
	"go/build"
	"os"
	"strings"

	"github.com/magefile/mage/mg"

	"zvelo.io/msg/internal/swagger"
	"zvelo.io/zmage"
)

// Default is the default mage target
var Default = Generate

// Go generates .pb.go files from .proto files
func Go(ctx context.Context) error {
	_, err := zmage.ProtoGo()
	return err
}

// Python generates _pb2.py and _pb2_grpc.py files from .proto files
func Python(ctx context.Context) error {
	_, err := zmage.ProtoPython()
	return err
}

// GRPCGateway generates .pb.gw.go files from .proto files
func GRPCGateway(ctx context.Context) error {
	_, err := zmage.ProtoGRPCGateway()
	return err
}

// Swagger generates .swagger.json files from .proto files
func Swagger(ctx context.Context) error {
	files, err := zmage.ProtoSwagger()
	if err != nil {
		return err
	}

	for _, file := range files {
		file = strings.Replace(file, ".proto", ".swagger.json", -1)
		if err := swagger.Patch(file); err != nil {
			_ = os.RemoveAll(file)
			return err
		}
	}

	return nil
}

// Static embeds static files into the internal/static package
func Static(ctx context.Context) error {
	mg.CtxDeps(ctx, Swagger)

	const out = "internal/static/static.go"

	files := []string{
		"schema.graphql",
		"apiv1.swagger.json",
	}

	modified, err := zmage.Modified(out, files...)
	if !modified || err != nil {
		return err
	}

	return zmage.Embed(zmage.EmbedConfig{
		OutputFile: out,
		Package:    "static",
		Files:      files,
	})
}

// Generate all necessary files
func Generate(ctx context.Context) error {
	mg.CtxDeps(ctx, Go, Python, GRPCGateway, Swagger, Static)
	return nil
}

// Test runs `go vet` and `go test -race` on all packages in the repository
func Test(ctx context.Context) error {
	return zmage.GoTest(ctx)
}

// CoverOnly calculates test coverage for all packages in the repository
func CoverOnly(ctx context.Context) error {
	return zmage.CoverOnly()
}

// Cover runs CoverOnly and opens the results in the browser
func Cover(ctx context.Context) error {
	return zmage.Cover()
}

// Fmt ensures that all go files are properly formatted
func Fmt(ctx context.Context) error {
	return zmage.GoFmt()
}

// Install runs `go install ./...`
func Install(ctx context.Context) error {
	mg.CtxDeps(ctx, Generate)
	return zmage.Install(build.Default)
}

// Lint runs `go vet` and `golint`
func Lint(ctx context.Context) error {
	return zmage.GoLint(ctx)
}

// GoVet runs `go vet`
func GoVet(ctx context.Context) error {
	return zmage.GoVet()
}

// Clean up after yourself
func Clean(ctx context.Context) error {
	return zmage.Clean()
}

// GoPackages lists all the non-vendor packages in the repository
func GoPackages(ctx context.Context) error {
	pkgs, err := zmage.GoPackages(build.Default)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		fmt.Println(pkg.ImportPath)
	}

	return nil
}
