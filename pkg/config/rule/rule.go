package rule

import (
	"os"
	"regexp"
	"strings"

	"github.com/ivanilves/travelgrunt/pkg/config/mode"
)

// Rule is a single configuration rule entity (a list of rules is contained by a config.Config)
type Rule struct {
	// Mode is a behaviour backed by a function from the `mode` package
	Mode string `yaml:"mode"`
	// Prefix is a literal path prefix to be matched against the passed file path
	Prefix string `yaml:"prefix"`
	// NameEx is a regex to be matched against the passed file base name
	NameEx string `yaml:"name"`
	// Negate reverses the rule match effect
	Negate bool `yaml:"negate"`

	// ModeFn is an `mode` package function to invoke on passed file / dir entry
	ModeFn func(os.DirEntry) bool
}

func (r Rule) pathPrefixed(path string) bool {
	if r.Prefix != "" {
		return strings.HasPrefix(path, r.Prefix)
	}

	return true
}

func (r Rule) normalizeNameEx() string {
	rx := regexp.MustCompile(`^\*(\.[[:alnum:]]+)+$`)

	if rx.MatchString(r.NameEx) {
		nameEx := r.NameEx

		nameEx = strings.ReplaceAll(nameEx, ".", "\\.")
		nameEx = strings.ReplaceAll(nameEx, "*", ".*")
		nameEx = "^" + nameEx + "$"

		return nameEx
	}

	return r.NameEx
}

func (r Rule) nameMatched(d os.DirEntry) bool {
	if r.NameEx != "" {
		if !mode.FileOrSymlink(d) {
			return false
		}

		rx := regexp.MustCompile(r.normalizeNameEx())

		return rx.MatchString(d.Name())
	}

	return true
}

func (r Rule) match(d os.DirEntry, relPath string) bool {
	return r.pathPrefixed(relPath) && r.nameMatched(d)
}

// Admit is a local "decider" function that includes/excludes the path given [on the single rule level]
func (r Rule) Admit(d os.DirEntry, relPath string) bool {
	if r.ModeFn == nil {
		return r.match(d, relPath)
	}

	return r.ModeFn(d) && r.match(d, relPath)
}
