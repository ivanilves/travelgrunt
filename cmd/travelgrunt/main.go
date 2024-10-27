package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/ivanilves/travelgrunt/pkg/config"
	"github.com/ivanilves/travelgrunt/pkg/directory"
	"github.com/ivanilves/travelgrunt/pkg/directory/tree"
	"github.com/ivanilves/travelgrunt/pkg/file"
	"github.com/ivanilves/travelgrunt/pkg/filter"
	"github.com/ivanilves/travelgrunt/pkg/menu"
	"github.com/ivanilves/travelgrunt/pkg/scm"
	"github.com/ivanilves/travelgrunt/pkg/terminal"
)

var appVersion = "default"

var outFile string
var expression string
var editFile bool
var top bool
var version bool

func init() {
	flag.StringVar(&outFile, "out-file", "", "output selected path into the file specified")
	flag.StringVar(&expression, "x", "", "use arbitrary expression passed instead of configured rules")
	flag.BoolVar(&editFile, "e", false, "edit file selected instead of changing working directory")
	flag.BoolVar(&top, "top", false, "get to the repository top level (root) path")
	flag.BoolVar(&version, "version", false, "print application version and exit")
}

func usage() {
	println("Usage: " + os.Args[0] + " [<match> <match2> ... <matchN>]")
	println("")
	println("Options:")
	flag.PrintDefaults()
}

func writeFileAndExit(fileName string, data string) {
	if err := file.Write(fileName, data); err != nil {
		log.Fatalf("failed to write file (%s): %s", fileName, err.Error())
	}

	os.Exit(0)
}

func buildMenuFromTree(t tree.Tree) string {
	var selected string
	var parentID string

	for idx := -1; idx < t.LevelCount(); idx++ {
		if len(parentID) > 0 {
			if !t.GetNode(parentID).HasChildren() {
				return parentID
			}
		}

		selected, err := menu.Build(t.LevelChildNames(idx, parentID), terminal.Height(), parentID)

		if err != nil {
			if err.Error() == "^C" {
				os.Exit(1)
			}

			if err.Error() == "^D" {
				return parentID
			}

			log.Fatalf("failed to build menu: %s", err.Error())
		}

		if selected == "." {
			return parentID
		}

		parentID = t.LevelChildItems(idx, parentID)[selected]
	}

	return selected
}

func getEntryPath(entries map[string]string, selected, rootPath string) (path string) {
	path = entries[selected]

	if path == "" {
		path = rootPath + "/" + selected
	}

	return path
}

func tag(s string) {
	sep := strings.Repeat("-", len(s))

	color.Blue(sep)
	color.HiYellow(s)
	color.Blue(sep)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if version {
		println(appVersion)

		os.Exit(0)
	}

	rootPath, err := scm.RootPath()

	if err != nil {
		log.Fatalf("failed to extract top level filesystem path from SCM: %s", err.Error())
	}

	if top {
		tag(rootPath)

		writeFileAndExit(outFile, rootPath)
	}

	cfg, err := config.NewConfig(rootPath)

	if err != nil {
		log.Fatalf("failed to load travelgrunt config: %s", err.Error())
	}

	if expression != "" {
		cfg = cfg.WithNameEx(expression)
	}

	cfg.UseFiles = editFile

	entries, paths, err := directory.Collect(rootPath, cfg)

	if err != nil {
		log.Fatalf("failed to collect directories: %s", err.Error())
	}

	if err := filter.Validate(flag.Args()); err != nil {
		log.Fatalf("invalid filter: %s", err.Error())
	}

	selected := buildMenuFromTree(
		tree.NewTree(filter.Apply(paths, flag.Args())),
	)

	if outFile != "" {
		path := getEntryPath(entries, selected, rootPath)

		if !editFile {
			tag(path)
		}

		writeFileAndExit(outFile, path)
	}

	log.Fatal("Please configure shell aliases as described: https://github.com/ivanilves/travelgrunt#shell-aliases")
}
