package domain_test

import (
	"testing"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/loan"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/book"
)

func TestUserCanBorrowWhenAllConditionsMet(t *testing.T) {
	service := loan.NewLoanEligibilityService()

	// Create Value Objects
	userId, _ := user.NewUserId("user-123")
	email, _ := user.NewEmail("john@example.com")
	u, _ := user.NewUser(userId, "John Doe", email, 0, false)
	
	bookId, _ := book.NewBookId("book-456")
	isbn, _ := book.NewISBN("9780134494166")
	b, _ := book.NewBook(bookId, "Clean Architecture", "Robert Martin", isbn, 3)

	if !service.CanBorrow(u, b) {
		t.Error("User should be able to borrow")
	}
}

func TestUserCannotBorrowWhenMaxLoansReached(t *testing.T) {
	service := loan.NewLoanEligibilityService()

	// Create Value Objects
	userId, _ := user.NewUserId("user-123")
	email, _ := user.NewEmail("john@example.com")
	u, _ := user.NewUser(userId, "John Doe", email, 5, false)
	
	bookId, _ := book.NewBookId("book-456")
	isbn, _ := book.NewISBN("9780134494166")
	b, _ := book.NewBook(bookId, "Clean Architecture", "Robert Martin", isbn, 3)

	if service.CanBorrow(u, b) {
		t.Error("User should not be able to borrow (max loans reached)")
	}
}