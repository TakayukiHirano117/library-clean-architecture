package user

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type UserId struct {
	value string
}

// NewUserId は検証付きでUserIdを作成します（8桁）
func NewUserId(value string) (*UserId, error) {
	matched, err := regexp.MatchString(`^\d{8}$`, value)
	if err != nil {
		return nil, fmt.Errorf("failed to validate UserId format: %w", err)
	}
	if !matched {
		return nil, errors.New("UserId must be exactly 8 digits. Got: " + value)
	}
	return &UserId{value: value}, nil
}

// GenerateUserId はランダムな8桁のUserIdを作成します
func GenerateUserId() *UserId {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomId := rand.Intn(90000000) + 10000000
	return &UserId{value: strconv.Itoa(randomId)}
}

func (u *UserId) Value() string {
	return u.value
}

func (u *UserId) Equals(other *UserId) bool {
	return u.value == other.value
}

func (u *UserId) String() string {
	return u.value
}
