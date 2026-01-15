package user_test

import (
	"testing"
	"time"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
)

func TestUser(t *testing.T) {
	t.Run("CanBorrow", func(t *testing.T) {
		u := user.NewUser("John Doe", "john@example.com")

		if !u.CanBorrow() {
			t.Error("New user should be able to borrow")
		}

		if u.Name() != "John Doe" {
			t.Errorf("Expected name 'John Doe', got '%s'", u.Name())
		}
		if u.Email() != "john@example.com" {
			t.Errorf("Expected email 'john@example.com', got '%s'", u.Email())
		}
		if u.Status() != user.UserStatusActive {
			t.Errorf("Expected status 'active', got '%s'", u.Status())
		}
		if u.CurrentBorrowCount() != 0 {
			t.Errorf("Expected borrow count 0, got %d", u.CurrentBorrowCount())
		}
	})

	t.Run("BorrowBook", func(t *testing.T) {
		u := user.NewUser("John Doe", "john@example.com")

		u2, err := u.BorrowBook()
		if err != nil {
			t.Fatalf("Failed to borrow book: %v", err)
		}

		if u2.CurrentBorrowCount() != 1 {
			t.Errorf("Expected borrow count 1, got %d", u2.CurrentBorrowCount())
		}

		// Original user should be unchanged (immutability)
		if u.CurrentBorrowCount() != 0 {
			t.Error("Original user should remain unchanged")
		}
	})

	t.Run("CannotBorrowWhenSuspended", func(t *testing.T) {
		userId, err := user.NewUserId("12345678")
		if err != nil {
			t.Fatal(err)
		}

		// Reconstruct suspended user
		u := user.ReconstructUser(
			userId,
			"John Doe",
			"john@example.com",
			user.UserStatusSuspended,
			0,
			0,
			time.Now(),
		)

		if u.CanBorrow() {
			t.Error("Suspended user should not be able to borrow")
		}
	})

	t.Run("CannotBorrowAtMaxLimit", func(t *testing.T) {
		userId, err := user.NewUserId("12345678")
		if err != nil {
			t.Fatal(err)
		}

		// Reconstruct user with max borrow count
		u := user.ReconstructUser(
			userId,
			"John Doe",
			"john@example.com",
			user.UserStatusActive,
			user.MaxBorrowLimit, // At limit
			0,
			time.Now(),
		)

		if u.CanBorrow() {
			t.Error("User should not be able to borrow (max loans reached)")
		}
	})

	t.Run("CannotBorrowWithOverdueFees", func(t *testing.T) {
		userId, err := user.NewUserId("12345678")
		if err != nil {
			t.Fatal(err)
		}

		// Reconstruct user with overdue fees
		u := user.ReconstructUser(
			userId,
			"John Doe",
			"john@example.com",
			user.UserStatusActive,
			0,
			10.50, // Has overdue fees
			time.Now(),
		)

		if u.CanBorrow() {
			t.Error("User with overdue fees should not be able to borrow")
		}
	})

	t.Run("FeeManagement", func(t *testing.T) {
		u := user.NewUser("John Doe", "john@example.com")

		u2, err := u.AddOverdueFee(5.00)
		if err != nil {
			t.Fatalf("Failed to add overdue fee: %v", err)
		}

		if u2.OverdueFees() != 5.00 {
			t.Errorf("Expected overdue fees 5.00, got %.2f", u2.OverdueFees())
		}

		// Test negative fee
		_, err = u.AddOverdueFee(-1.00)
		if err == nil {
			t.Error("Expected error for negative fee")
		}
	})

	t.Run("Immutability", func(t *testing.T) {
		userId, err := user.NewUserId("12345678")
		if err != nil {
			t.Fatal(err)
		}
		u := user.ReconstructUser(
			userId,
			"John Doe",
			"john@example.com",
			user.UserStatusActive,
			2,
			0,
			time.Now(),
		)

		// Test ReturnBook immutability
		u2, err := u.ReturnBook()
		if err != nil {
			t.Fatalf("Failed to return book: %v", err)
		}

		if u2.CurrentBorrowCount() != 1 {
			t.Errorf("Expected borrow count 1, got %d", u2.CurrentBorrowCount())
		}

		// Original unchanged
		if u.CurrentBorrowCount() != 2 {
			t.Error("Original user should remain unchanged")
		}

		// Test Suspend immutability
		u3 := u.Suspend()
		if u3.Status() != user.UserStatusSuspended {
			t.Errorf("Expected status 'suspended', got '%s'", u3.Status())
		}

		// Original unchanged
		if u.Status() != user.UserStatusActive {
			t.Error("Original user should remain active")
		}
	})
}