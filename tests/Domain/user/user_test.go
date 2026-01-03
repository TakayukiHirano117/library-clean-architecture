package domain_test

import (
	"testing"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
)

func TestCanCreateValidUser(t *testing.T) {
	// Create Value Object
	userId, _ := user.NewUserId("user-123")
	email, _ := user.NewEmail("john@example.com")

	u, err := user.NewUser(userId, "John Doe", email, 0, false)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if u.GetId().GetValue() != "user-123" {
		t.Errorf("Expected ID 'user-123', got '%s'", u.GetId().GetValue())
	}
	if u.GetName() != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", u.GetName())
	}
	if u.GetEmail().GetValue() != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", u.GetEmail())
	}
}

// TODO: voのテストで検証する
// func TestInvalidEmailReturnsError(t *testing.T) {
// 	userId, _ := user.NewUserId("user-123")

// 	_, err := user.NewUser(userId, "John Doe", "invalid-email", 0, false)
// 	if err == nil {
// 		t.Error("Expected error for invalid email")
// 	}
// }

func TestCannotBorrowMoreWhenMaxLoansReached(t *testing.T) {
	userId, _ := user.NewUserId("user-123")
	email, _ := user.NewEmail("john@example.com")

	u, _ := user.NewUser(userId, "John Doe", email, 5, false)

	if u.CanBorrowMore() {
		t.Error("User should not be able to borrow more (max loans reached)")
	}
}

func TestCanBorrowMoreWhenUnderLimit(t *testing.T) {
	userId, _ := user.NewUserId("user-123")
	email, _ := user.NewEmail("john@example.com")
	u, _ := user.NewUser(userId, "John Doe", email, 2, false)

	if !u.CanBorrowMore() {
		t.Error("User should be able to borrow more")
	}
}