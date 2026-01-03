package book

import "errors"

type BookId struct {
  value string
}

func NewBookId(value string) (*BookId, error) {
  if len(value) == 0 {
    return nil, errors.New("BookId cannot be empty")
  }

  return &BookId{value: value}, nil
}

func (b *BookId) GetValue() string {
  return b.value
}

func (b *BookId) Equals(other *BookId) bool {
  return b.value == other.value
}

func (b *BookId) String() string {
  return b.value
}