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

// Build creates an interactive menu to chose Terragrunt project from
func Build(items []string, maxSize int, previous string) (selected string, err error) {
	if len(items) == 0 {
		return "", fmt.Errorf("no items")
	}

	if len(previous) > 0 {
		fmt.Printf("=> %s\n", previous)
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
		Label:    label,
		Items:    items,
		Size:     getSize(len(items), maxSize-Overhead),
		Searcher: searcher,
	}

	_, selected, err = prompt.Run()

	if err != nil {
		return "", err
	}

	return selected, nil
}
