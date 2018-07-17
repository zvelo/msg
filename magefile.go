// +build mage

package main

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	"zvelo.io/msg/internal/swagger"
	"zvelo.io/zmage"
)

// Default is the default mage target
var Default = Generate

// Prototool generates protobuf files
func Prototool(ctx context.Context) error {
	destinations := []string{
		"msgpb/msg.pb.go",
		"msgpb/msg.pb.gw.go",
		"msg.swagger.json",
	}

	var modified bool
	for _, dest := range destinations {
		var err error
		modified, err = zmage.Modified(dest, "msg.proto", "prototool.yaml")
		if err != nil {
			return err
		}
		if modified {
			break
		}
	}

	if !modified {
		return nil
	}

	if err := sh.Run("prototool", "gen"); err != nil {
		return err
	}

	return swagger.Patch("msg.swagger.json")
}

// ProtoPython generates _pb2.py and _pb2_grpc.py files from .proto files
func ProtoPython(ctx context.Context) error {
	python := "python"

	if _, err := exec.LookPath("python3"); err == nil {
		python = "python3"
	}

	if exe := os.Getenv("PYTHON"); exe != "" {
		python = exe
	}

	matches, err := filepath.Glob("*.proto")
	if err != nil {
		return err
	}

	args := []string{
		"-m", "grpc_tools.protoc",
		"-Iinclude",
		"-I.",
		"--python_out=./python/",
		"--grpc_python_out=./python/",
	}

	args = append(args, matches...)

	if err = os.MkdirAll("python", 0755); err != nil {
		return err
	}

	return sh.Run(python, args...)
}

// Descriptor generates protobuf file descriptor set files from .proto files
func Descriptor(ctx context.Context) error {
	modified, err := zmage.Modified("msg.protoset", "msg.proto", "include/health/health.proto")
	if err != nil {
		return err
	}

	if !modified {
		return nil
	}

	return sh.Run(
		"protoc",
		"-Iinclude",
		"-I.",
		"--descriptor_set_out=msg.protoset",
		"--include_imports",
		"msg.proto",
		"include/health/health.proto",
	)
}

// Static embeds static files into the internal/static package
func Static(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)

	const out = "internal/static/static.go"

	files := []string{
		"schema.graphql",
		"msg.swagger.json",
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
	mg.CtxDeps(ctx, CheckImports, Prototool, Descriptor, Static)
	return nil
}

// CheckImports ensures that zvelo import rules are followed
func CheckImports(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool, Static)

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
	mg.CtxDeps(ctx, Prototool)
	return zmage.GoTest(ctx)
}

// CoverOnly calculates test coverage for all packages in the repository
func CoverOnly(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)
	return zmage.CoverOnly()
}

// Cover runs CoverOnly and opens the results in the browser
func Cover(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)
	return zmage.Cover()
}

// Fmt ensures that all go files are properly formatted
func Fmt(ctx context.Context) error {
	return zmage.GoFmt()
}

// Install runs `go install ./...`
func Install(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)
	return zmage.Install(build.Default)
}

// Lint runs `go vet` and `golint`
func Lint(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)
	return zmage.GoLint(ctx)
}

// GoVet runs `go vet`
func GoVet(ctx context.Context) error {
	mg.CtxDeps(ctx, Prototool)
	return zmage.GoVet()
}

// Clean up after yourself
func Clean(ctx context.Context) error {
	return zmage.Clean("msg.protoset")
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
