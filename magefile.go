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

// ProtoGo generates .pb.go files from .proto files
func ProtoGo(ctx context.Context) error {
	_, err := zmage.ProtoGo()
	return err
}

// ProtoPython generates _pb2.py and _pb2_grpc.py files from .proto files
func ProtoPython(ctx context.Context) error {
	_, err := zmage.ProtoPython()
	return err
}

// ProtoGRPCGateway generates .pb.gw.go files from .proto files
func ProtoGRPCGateway(ctx context.Context) error {
	_, err := zmage.ProtoGRPCGateway()
	return err
}

// Descriptor generates protobuf file descriptor set files from .proto files
func Descriptor(ctx context.Context) error {
	_, err := zmage.Descriptor("zvelo-api.protoset",
		"apiv1.proto",
		"health/health.proto",
	)
	return err
}

// Swagger generates .swagger.json files from .proto files
func ProtoSwagger(ctx context.Context) error {
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
	mg.CtxDeps(ctx, ProtoSwagger)

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
	mg.CtxDeps(ctx, CheckImports, ProtoGo, ProtoPython, ProtoGRPCGateway, ProtoSwagger, Descriptor, Static)
	return nil
}

// CheckImports ensures that zvelo import rules are followed
func CheckImports(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway, Static)

	pkgs, err := zmage.GoPackages(build.Default)
	if err != nil {
		return err
	}

	var failed bool
	for _, pkg := range pkgs {
		for _, imp := range pkg.Imports {
			if strings.HasPrefix(imp, "github.com/zvelo/") {
				fmt.Fprintf(os.Stderr, "package %q depends on disallowed import of %s\n", pkg.ImportPath, imp)
				failed = true
			}

			if !strings.HasPrefix(imp, "zvelo.io/") {
				continue
			}

			if imp == "zvelo.io/msg" {
				continue
			}

			if strings.HasPrefix(imp, "zvelo.io/msg/") {
				continue
			}

			fmt.Fprintf(os.Stderr, "package %q depends on disallowed import of %s\n", pkg.ImportPath, imp)
			failed = true
		}
	}

	if failed {
		return fmt.Errorf("import errors")
	}

	return nil
}

// Test runs `go vet` and `go test -race` on all packages in the repository
func Test(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.GoTest(ctx)
}

// CoverOnly calculates test coverage for all packages in the repository
func CoverOnly(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.CoverOnly()
}

// Cover runs CoverOnly and opens the results in the browser
func Cover(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.Cover()
}

// Fmt ensures that all go files are properly formatted
func Fmt(ctx context.Context) error {
	return zmage.GoFmt()
}

// Install runs `go install ./...`
func Install(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.Install(build.Default)
}

// Lint runs `go vet` and `golint`
func Lint(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.GoLint(ctx)
}

// GoVet runs `go vet`
func GoVet(ctx context.Context) error {
	mg.CtxDeps(ctx, ProtoGo, ProtoGRPCGateway)
	return zmage.GoVet()
}

// Clean up after yourself
func Clean(ctx context.Context) error {
	return zmage.Clean("zvelo-api.protoset")
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
