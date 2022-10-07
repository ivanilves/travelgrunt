package menu

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const label = "Select Terragrunt project to travel"

// Build creates an interactive menu to chose Terragrunt project from
func Build(items []string) (selected string, err error) {
	if len(items) == 0 {
		return "", fmt.Errorf("no items")
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
		Size:     len(items),
		Searcher: searcher,
	}

	_, selected, err = prompt.Run()

	if err != nil {
		return "", err
	}

	return selected, nil
}
