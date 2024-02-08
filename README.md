# UnitMaker - Go Unit Test Template Generator

## Introduction
UnitMaker is a Go application that automatically generates basic unit test templates for Go functions. It reads a Go source file, analyzes the functions, and creates corresponding test function skeletons in a `_test.go` file. This tool simplifies the process of starting unit tests in Go projects.

## Prerequisites
- Go 1.18 or later

## Installation
1. Clone the repository:
git clone https://github.com/chrisjchandler/unitmaker.git

2. Navigate to the project directory:
3. Build the application:

go build -o unitmaker main.go

## Usage
Run UnitMaker with a Go file as an argument to generate test templates:
./gotestgen path/to/yourfile.go


## Output

This will create a `yourfile_test.go` file in the same directory with test templates for each function found in `yourfile.go`.

### Output File
UnitMaker generates a test file named after the input file but with the suffix `_test.go`. This file includes:
- A package declaration matching the input file's package.
- Import statement for the `testing` package.
- Basic test function templates for each function in the input file.

Example of generated test function:
```go
func TestExampleFunction(t *testing.T) {
    // TODO: Add necessary parameters
    got := ExampleFunction(/* params */)
    if got != nil {
        t.Errorf("ExampleFunction() = %v, want nil")
    }
}

//Sample of what those parameters may look like for a web observability test I made


func TestextractHostname(t *testing.T) {
    testURL := "http://aigatehouse.com"
    got := extractHostname(testURL)
    want := "aigatehouse.com" // expected result

    if got != want {
        t.Errorf("extractHostname(%q) = %v, want %v", testURL, got, want)
    }
}



Limitations
The generated tests are basic templates and require manual completion.
Assumes functions with return values might return an error, which may not always be true.
May not handle complex functions with multiple parameters and return types.
Overwrites existing _test.go files without confirmation.
Contributing
Contributions to improve UnitMaker are welcome. Feel free to fork the repository, make changes, and submit a pull request.

License:
I don't care.
