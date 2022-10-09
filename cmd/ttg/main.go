package main

import (
	"flag"
	"log"

	"github.com/ivanilves/ttg/pkg/directory"
	"github.com/ivanilves/ttg/pkg/filter"
	"github.com/ivanilves/ttg/pkg/menu"
	"github.com/ivanilves/ttg/pkg/scm"
	"github.com/ivanilves/ttg/pkg/shell"
)

func usage() {
	log.Fatalf("Usage: %s <match> [<match2> ... <matchN>]\n", shell.Name())
}

func main() {
	if shell.IsRunningInside() {
		log.Fatalf("%s already running (pid: %d), please type \"exit\" to return to the parent shell first", shell.Name(), shell.Getppid())
	}

	flag.Usage = usage
	flag.Parse()

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

	selected, err := menu.Build(filter.Apply(entries, matches))

	if err != nil {
		log.Fatalf("failed to build menu: %s", err.Error())
	}

	shell.Spawn(entries[selected])
}
