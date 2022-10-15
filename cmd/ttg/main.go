package main

import (
	"flag"
	"log"
	"os"

	"github.com/ivanilves/ttg/pkg/directory"
	"github.com/ivanilves/ttg/pkg/file"
	"github.com/ivanilves/ttg/pkg/filter"
	"github.com/ivanilves/ttg/pkg/menu"
	"github.com/ivanilves/ttg/pkg/scm"
	"github.com/ivanilves/ttg/pkg/shell"
	"github.com/ivanilves/ttg/pkg/terminal"
)

var appVersion = "default"

var version bool
var outFile string

func init() {
	flag.BoolVar(&version, "version", false, "print application version and exit")
	flag.StringVar(&outFile, "outFile", "", "output project path into the file specified instead of spawning a shell")
}

func usage() {
	println("Usage: " + shell.Name() + " [<match> <match2> ... <matchN>]")
	println("")
	println("Options:")
	flag.PrintDefaults()
}

func main() {
	if shell.IsRunningInside() {
		log.Fatalf("%s already running (pid: %d), please type \"exit\" to return to the parent shell first", shell.Name(), shell.Getppid())
	}

	flag.Usage = usage
	flag.Parse()

	if version {
		println(appVersion)

		os.Exit(0)
	}

	matches := flag.Args()

	rootPath, err := scm.RootPath()

	if err != nil {
		log.Fatalf("failed to extract top level filesystem path from SCM: %s", err.Error())
	}

	entries, err := directory.Collect(rootPath)

	if err != nil {
		log.Fatalf("failed to collect Terragrunt project directories: %s", err.Error())
	}

	if err := filter.Validate(matches); err != nil {
		log.Fatalf("invalid filter: %s", err.Error())
	}

	selected, err := menu.Build(filter.Apply(entries, matches), terminal.Height())

	if err != nil {
		log.Fatalf("failed to build menu: %s", err.Error())
	}

	if outFile != "" {
		if err := file.Write(outFile, entries[selected]); err != nil {
			log.Fatalf("failed to write output file: %s", err.Error())
		}

		os.Exit(0)
	}

	shell.Spawn(entries[selected])
}
