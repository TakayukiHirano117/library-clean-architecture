package domain_test

import (
	"testing"
	"time"

	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/book"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/loan"
	"github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
)

func TestDueDateIsCalculatedCorrectly(t *testing.T) {
	borrowedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	loanId, _ := loan.NewLoanId("loan-123")
	userId, _ := user.NewUserId("user-456")
	bookId, _ := book.NewBookId("book-789")

	loan := loan.NewLoan(
		loanId,
		userId,
		bookId,
		borrowedAt,
		nil, // returnedAt
	)

	expectedDueDate := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	if !loan.GetDueDate().Equal(expectedDueDate) {
		t.Errorf("Expected due date %v, got %v", expectedDueDate, loan.GetDueDate())
	}
}

func TestLoanIsOverdueAfterDueDate(t *testing.T) {
	borrowedAt := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	currentDate := time.Date(2025, 1, 20, 0, 0, 0, 0, time.UTC)

	loanId, _ := loan.NewLoanId("loan-123")
	userId, _ := user.NewUserId("user-456")
	bookId, _ := book.NewBookId("book-789")

	loan := loan.NewLoan(loanId, userId, bookId, borrowedAt, nil)

	if !loan.IsOverdue(&currentDate) {
		t.Error("Expected loan to be overdue")
	}
}

func TestCanMarkLoanAsReturned(t *testing.T) {
	loanId, _ := loan.NewLoanId("loan-123")
	userId, _ := user.NewUserId("user-456")
	bookId, _ := book.NewBookId("book-789")

	loan := loan.NewLoan(loanId, userId, bookId, time.Now(), nil)

	if loan.IsReturned() {
		t.Error("Loan should not be returned initially")
	}

	now := time.Now()
	loan.MarkAsReturned(&now)

	if !loan.IsReturned() {
		t.Error("Loan should be marked as returned")
	}
	if loan.GetReturnedAt().IsZero() {
		t.Error("ReturnedAt should be set")
	}
}
