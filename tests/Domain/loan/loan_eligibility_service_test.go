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
	u := user.NewUser("John Doe", "john@example.com")
	
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
	u := user.NewUser("John Doe", "john@example.com")
	
	for i := 0; i < user.MaxBorrowLimit; i++ {
		u, _ = u.BorrowBook()
	}

	bookId, _ := book.NewBookId("book-456")
	isbn, _ := book.NewISBN("9780134494166")
	b, _ := book.NewBook(bookId, "Clean Architecture", "Robert Martin", isbn, 3)

	if service.CanBorrow(u, b) {
		t.Error("User should not be able to borrow (max loans reached)")
	}
}