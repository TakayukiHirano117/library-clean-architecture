// Go uses built-in testing package
// Test files should be named *_test.go
// Example: BookId_test.go, ISBN_test.go

// Run tests with:
// go test ./...

// Run with coverage:
// go test -cover ./...

# Testing in Go

Go has a built-in testing framework.
No configuration files needed!

## Test File Naming
- Test files must end with _test.go
- Test functions must start with Test
- Example: func TestBookId(t *testing.T)

## Running Tests
```bash
go test ./...              # Run all tests
go test -v ./...           # Verbose output
go test -cover ./...       # With coverage
go test ./src/domain/book  # Specific package
```