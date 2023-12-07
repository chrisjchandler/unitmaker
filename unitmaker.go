package main

import (
    "flag"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "os"
    "strings"
)

func main() {
    // Parsing command line argument
    flag.Parse()
    args := flag.Args()

    if len(args) != 1 {
        fmt.Println("Please provide a Go file as an argument")
        os.Exit(1)
    }

    filename := args[0]
    fset := token.NewFileSet()

    // Parsing the Go file
    node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
    if err != nil {
        fmt.Println("Error parsing file:", err)
        os.Exit(1)
    }

    // Create the test file
    testFilename := strings.TrimSuffix(filename, ".go") + "_test.go"
    testFile, err := os.Create(testFilename)
    if err != nil {
        fmt.Println("Error creating test file:", err)
        os.Exit(1)
    }
    defer testFile.Close()

    // Writing package declaration file
    fmt.Fprintf(testFile, "package %s\n\nimport \"testing\"\n\n", node.Name.Name)

    // Iterating through the AST and finding functions
    for _, f := range node.Decls {
        fn, ok := f.(*ast.FuncDecl)
        if !ok {
            continue
        }
        generateTest(fn, testFile)
    }
}

func generateTest(fn *ast.FuncDecl, testFile *os.File) {
    // Check if the function has parameters or return values
    hasParams := fn.Type.Params != nil && len(fn.Type.Params.List) > 0
    hasReturns := fn.Type.Results != nil && len(fn.Type.Results.List) > 0

    // Function test name
    testName := "Test" + fn.Name.Name

    // Start building the test function
    var testFunctionBuilder strings.Builder
    testFunctionBuilder.WriteString(fmt.Sprintf("func %s(t *testing.T) {\n", testName))

    // If the function has parameters, we need to handle them (here we just pass zero values)
    if hasParams {
        testFunctionBuilder.WriteString("    // TODO: Add necessary parameters\n")
    }

    // Call the function
    call := fmt.Sprintf("    %s(", fn.Name.Name)
    if hasParams {
        call += "/* params */"
    }
    call += ")"
    if hasReturns {
        call = "got := " + call
    }
    testFunctionBuilder.WriteString(call + "\n")

    // If the function has return values, add a basic assertion
    if hasReturns {
        testFunctionBuilder.WriteString(fmt.Sprintf("    if got != nil {\n"))
        testFunctionBuilder.WriteString(fmt.Sprintf("        t.Errorf(\"%s() = %%v, want nil\")\n", fn.Name.Name))
        testFunctionBuilder.WriteString("    }\n")
    }

    testFunctionBuilder.WriteString("}\n\n")

    // Write the test function to the test file
    fmt.Fprint(testFile, testFunctionBuilder.String())
}
