package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

// TranspilationRule represents a rule for transpiling code.
type TranspilationRule struct {
	Pattern *regexp.Regexp
	Replace func([]string) string
}

// transpile transpiles code according to a set of rules.
func transpile(input string, rules []TranspilationRule) string {
	for _, rule := range rules {
		input = rule.Pattern.ReplaceAllStringFunc(input, func(match string) string {
			parts := rule.Pattern.FindStringSubmatch(match)
			return rule.Replace(parts)
		})
	}
	return input
}

// readFile reads the content of a file
func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// writeFile writes content to a file
func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// executeFile executes a Go file
func executeFile(filename string) error {
	cmd := exec.Command("go", "run", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

//func main() {
//	inputFile := "input/input.go"
//	outputFile := "transpiled/transpiled.go"
//
//	// Step 1: Read the input file
//	inputContent, err := readFile(inputFile)
//	if err != nil {
//		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
//		return
//	}
//
//	// Step 2: Transpile the content
//	transpiledContent := transpile(inputContent)
//
//	// Step 3: Write the transpiled content to the output file
//	err = writeFile(outputFile, transpiledContent)
//	if err != nil {
//		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
//		return
//	}
//
//	// Step 4: Execute the transpiled Go file
//	err = executeFile(outputFile)
//	if err != nil {
//		fmt.Printf("Error executing file %s: %v\n", outputFile, err)
//	}
//}

func main() {
	// Define the transpilation rules
	rules := []TranspilationRule{
		{
			Pattern: regexp.MustCompile(`(\w+)\s*\+\s*(\w+)`),
			Replace: func(parts []string) string {
				if len(parts) == 3 {
					num1 := parts[1]
					num2 := parts[2]
					return fmt.Sprintf("%s*%s + %s*%s", num1, num1, num2, num2)
				}
				return parts[0]
			},
		},
		// Add more rules here
	}

	// Example input code
	inputCode := `
	package main

	import "fmt"

	func main() {
		num1 := 2
		num2 := 3
		result := num1 + num2
		fmt.Println(result)
	}
	`

	// Transpile the input code
	transpiledCode := transpile(inputCode, rules)

	fmt.Println("Original Code:")
	fmt.Println(inputCode)
	fmt.Println("Transpiled Code:")
	fmt.Println(transpiledCode)
}
