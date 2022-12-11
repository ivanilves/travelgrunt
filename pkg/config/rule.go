package config

import (
	"os"
	"regexp"
	"strings"

	"github.com/ivanilves/travelgrunt/pkg/config/include"
)

// Rule is a single configuration rule entity (a list of rules is contained by a Config)
type Rule struct {
	// Mode is a behaviour backed by a function from the `include` package
	Mode string `yaml:"mode"`
	// Prefix is a literal path prefix to be matched against the passed file path
	Prefix string `yaml:"prefix"`
	// PathEx is a regex to be matched against the passed file full path
	PathEx string `yaml:"path"`
	// NameEx is a regex to be matched against the passed file base name
	NameEx string `yaml:"name"`
	// Exclude reverses the rule match effect
	Exclude bool `yaml:"exclude"`

	// IncludeFn is an `include` package function to invoke on passed file / dir entry
	IncludeFn func(os.DirEntry) bool
}

// Matched matches passed file / dir entry name against supplied expressions (if any)
func (r Rule) Matched(d os.DirEntry, rel string) bool {
	return r.pathPrefixed(rel) && r.pathMatched(rel) && r.nameMatched(d)
}

func (r Rule) pathPrefixed(path string) bool {
	if r.Prefix != "" {
		return strings.HasPrefix(path, r.Prefix)
	}

	return true
}

func (r Rule) pathMatched(path string) bool {
	if r.PathEx != "" {
		rx := regexp.MustCompile(r.PathEx)

		return rx.MatchString(path)
	}

	return true
}

func (r Rule) nameMatched(d os.DirEntry) bool {
	if r.NameEx != "" {
		if !include.FileOrSymlink(d) {
			return false
		}

		rx := regexp.MustCompile(r.NameEx)

		return rx.MatchString(d.Name())
	}

	return true
}

// Include is a local "decider" function that includes/excludes the path given [on the single rule level]
func (r Rule) Include(d os.DirEntry, rel string) bool {
	if r.IncludeFn == nil {
		return r.Matched(d, rel)
	}

	return r.IncludeFn(d) && r.Matched(d, rel)
}
