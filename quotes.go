package quotes

import (
	"errors"
	"fmt"
)

func Sailor() string {
	return "A smooth sea never made a skilled sailor."
}

func LifeChanges() string {
	return "Changes are hard at first, messy in the middle, and beautiful at the end."
}

func BuildQuote(name string) (string, error) {
	if name == "" {
		return "", errors.New("please pass your name as the argument to quotes.BuildQuote")
	}

	quote := fmt.Sprintf("%v, never settle.", name)
	return quote, nil
}
