package filter

import (
	"fmt"
	"sort"
	"strings"
)

// MinMatchLen defines a minimum length of the match string/pattern
const MinMatchLen = 2

// Validate checks match strings for the minimum length requirements
func Validate(matches []string) error {
	if len(matches) == 0 {
		return nil
	}

	for _, match := range matches {
		if len(match) < MinMatchLen {
			return fmt.Errorf("match pattern \"%s\" is shorter than required minimum of %d characters", match, MinMatchLen)
		}
	}

	return nil
}

// DoesMatch checks if string passed matches ALL patterns specified
func DoesMatch(s string, matches []string) bool {
	if len(matches) == 0 {
		return true
	}

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
