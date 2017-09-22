package service

import (
	"strings"
	"errors"
)

// StringService provides operations on strings.
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringService struct{}

func (StringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
