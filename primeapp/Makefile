.PHONY: all test test-cover cover cover-html cover-func clean

# Default: run everything important
all: test test-cover cover-func

# Run all tests in current directory
test:
	go test .

# Run tests with coverage summary
test-cover:
	go test -cover .

# Generate coverage profile
cover:
	go test -coverprofile=coverage.out .

# Show coverage per function in terminal
cover-func: cover
	go tool cover -func=coverage.out

# Open coverage in browser
cover-html: cover
	go tool cover -html=coverage.out

# Clean coverage files
clean:
	rm -f coverage.out
