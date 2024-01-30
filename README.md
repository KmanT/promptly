# Promptly
A library for creating Command Line Applications. All you need is bufio.Reader
and your imagination.

## Example
```go
package main

import (
	"KmanT/promptly"
	"bufio"
	"fmt"
	"os"
)

func main() {
	rdr := bufio.NewReader(os.Stdin)

	name := promptly.GetPromptVerifyRegexLoop(rdr, "What is your name?", `^[a-zA-Z]+$`)

	answers := []string{"7"}
	valid, _ := promptly.GetPromptVerify(rdr, "How many days are in the week?", answers, true)

	var status string
	if valid {
		status = "smarty"
	} else {
		status = "dummy"
	}

	fmt.Printf("Hello, %s! You are a %s", name, status)
}
```

## What's included

### GetSimplePromptText
Just get text from user input after a prompt

### GetPromptVerify
Get user input and verify the input against a list of valid answers

### GetPromptVerifyRegex
Get user input and verify that the input matches a given pattern

### GetPromptVerifyLoop
Get user input, verify the input against a list of valid answers, and loops
until the user has input a valid response

### GetPromptVerifyRegexLoop
Get user input, verify that the input matches a given pattern, and loops until
the user has input that is a valid response

## To do:
- Numeric bounds checking
- Simple prompt with custom middleware functions
- Safewords to exit / reset?

## In question:
- Multi-answer prompts
