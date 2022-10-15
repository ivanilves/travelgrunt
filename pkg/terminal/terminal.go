package terminal

import (
	"golang.org/x/term"
)

const defautHeight = 20

// Height reveals a number of lines we have in our terminal
func Height() int {
	_, h, err := term.GetSize(0)

	if err != nil {
		return defautHeight
	}

	return h
}
