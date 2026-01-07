package user

import (
	"errors"
	"regexp"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if len(value) == 0 {
		return nil, errors.New("Email cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !emailRegex.MatchString(value) {
		return nil, errors.New("Invalid email format")
	}

	return &Email{value: value}, nil
}

func (e *Email) GetValue() string {
	return e.value
}

func (e *Email) GetDomain() string {
  parts := strings.Split(e.value, "@")
  return parts[1]  // @以降の部分
}

func (e *Email) Equals(other *Email) bool {
	return e.value == other.value
}

func (e *Email) String() string {
	return e.value
}
