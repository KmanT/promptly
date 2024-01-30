# Promptly
A library for creating Command Line Applications

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

## Includes:
- A simple get input
- Getting input with validation with expected results
- Getting input with validation with regex
- Getting input and validation until the input is correct

## Todo:
- Numeric bounds checking
