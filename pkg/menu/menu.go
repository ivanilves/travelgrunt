package menu

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const label = "Select path to travel"

// Overhead shows how many lines are occupied by menu control elements
const Overhead = 3

// MinSize stands for minimal menu size
const MinSize = 5

func getSize(itemCount int, size int) int {
	if itemCount <= size {
		return itemCount
	}

	if size <= MinSize {
		return MinSize
	}

	return size
}

// Build creates an interactive menu to chose destination directory from
func Build(items []string, size int, parentID string) (selected string, err error) {
	var extension string

	if len(items) == 0 {
		return "", fmt.Errorf("no items")
	}

	if len(parentID) > 0 {
		extension = fmt.Sprintf(" [from \"%s\"]", parentID)
	}

	if len(items) == 1 {
		return items[0], nil
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.Replace(strings.ToLower(item), "/", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:    label + extension,
		Items:    items,
		Size:     getSize(len(items), size-Overhead),
		Searcher: searcher,
	}

	_, selected, err = prompt.Run()

	return selected, err
}
