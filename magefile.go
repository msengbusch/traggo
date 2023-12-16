//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"strings"
	"time"
)

const Module = "github.com/traggo/server"

const BuildDir = "build"

const MainFileCli = "cmd/cli/main.go"
const MainFileServer = "cmd/server/main.go"
const ExecFileCli = BuildDir + "/traggo-cli"
const ExecFileServer = BuildDir + "/traggo-server"

type Build mg.Namespace
type Run mg.Namespace

var Shell = sh.RunCmd("sh", "-c")
var GoBuild = sh.RunCmd("go", "build")

func Clean() error {
	err := sh.Rm(BuildDir)
	if err != nil {
		return fmt.Errorf("failed to clean build directory: %w", err)
	}
	return nil
}

func Install() error {
	err := sh.Run("bun", "install")
	if err != nil {
		return fmt.Errorf("failed to install js dependencies via bun: %w", err)
	}
	return nil
}

func (Build) Go(mode string) error {
	err := GoBuild(goFlags(mode)...)
	if err != nil {
		return fmt.Errorf("failed to build go: %s %w", mode, err)
	}
	return nil
}

func (Build) Web() error {
	return nil
}

func (Run) Go() error {
	return nil
}

func (Run) Web() error {
	return nil
}

func Docker() error {
	return nil
}

func goFlags(mode string) []string {
	var flags []string

	var target = "server"
	var tags = ""
	var ldflags = ""

	for _, m := range strings.Split(mode, ",") {
		switch m {
		case "dev":
			tags = tags + "dev "
		case "prod":
			tags = tags + "prod "

		case "static":
			ldflags = ldflags + " -linkmode external -extldflags '-static'"
		case "info":
			ldflags = ldflags +
				" -X " + Module + "/environment.BuildDate=" + date() +
				" -X " + Module + "/environment.BuildCommit=" + commit() +
				" -X " + Module + "/environment.BuildVersion=" + version()

		case "cli":
			target = "cli"
		case "server":
			target = "server"
		}
	}

	if tags != "" {
		flags = append(flags, "-tags", tags)
	}

	if ldflags != "" {
		flags = append(flags, "-ldflags", ldflags)
	}

	switch target {
	case "cli":
		flags = append(flags, "-o", ExecFileCli, MainFileCli)
	case "server":
		flags = append(flags, "-o", ExecFileServer, MainFileServer)
	}

	return flags
}

func commit() string {
	output, err := sh.Output("git", "rev-parse", "--verify", "HEAD")
	if err != nil {
		fmt.Println("failed to get current git commit: %w", err)
		return ""
	}
	return output
}

func date() string {
	return time.Now().Format("%Y-%m-%dT%H:%M:%SZ")
}

func version() string {
	output, err := sh.Output("git", "describe", "--tags", "--abbrev=0")
	if err != nil {
		fmt.Println("failed to get current git tag: %w", err)
		return ""
	}
	return output
}
