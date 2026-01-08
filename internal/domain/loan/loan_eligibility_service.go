package loan

import (
  "fmt"
  "github.com/TakayukiHirano117/library-clean-architecture/internal/domain/user"
  "github.com/TakayukiHirano117/library-clean-architecture/internal/domain/book"
)

// LoanEligibilityService - Domain Service for loan eligibility checks
type LoanEligibilityService struct{}

func NewLoanEligibilityService() *LoanEligibilityService {
  return &LoanEligibilityService{}
}

func (s *LoanEligibilityService) CanBorrow(u *user.User, b *book.Book) bool {
  // Rule 1: User must not have reached max loan limit
  if !u.CanBorrow() {
    return false
  }

  // Rule 2: User must not have overdue books
  if u.OverdueFees() > 0 {
    return false
  }

  // Rule 3: Book must have available copies
  if !b.IsAvailable() {
    return false
  }

  return true
}

func (s *LoanEligibilityService) GetIneligibilityReason(u *user.User, b *book.Book) *string {
  if !u.CanBorrow() {
    reason := fmt.Sprintf("User has reached maximum loan limit (%d books)", user.MaxBorrowLimit)
    return &reason
  }

  if u.OverdueFees() > 0 {
    reason := "User has overdue books"
    return &reason
  }

  if !b.IsAvailable() {
    reason := fmt.Sprintf("No copies available for book %s", b.GetTitle())
    return &reason
  }

  return nil
}