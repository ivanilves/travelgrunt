package filter

import (
	"fmt"
	"sort"
	"strings"
)

// MinMatchLen defines a minimum length of the first match string/pattern
const MinMatchLen = 3

// Validate checks match strings for the minimum length requirements
func Validate(matches []string) error {
	if len(matches[0]) < MinMatchLen {
		return fmt.Errorf("first match pattern \"%s\" is shorter than required minimum of %d characters", matches[0], MinMatchLen)
	}

	return nil
}

// DoesMatch checks if string passed matches ALL patterns specified
func DoesMatch(s string, matches []string) bool {
	for _, match := range matches {
		if !strings.Contains(s, match) {
			return false
		}
	}

	return true
}

// Apply selects strings matching ALL patterns specified and sorts the result (for the sake of determinism)
func Apply(entries map[string]string, matches []string) (list []string) {
	for name := range entries {
		if DoesMatch(name, matches) {
			list = append(list, name)
		}
	}

	sort.Strings(list)

	return list
}
