package promptly

import (
	"bufio"
	"log"
	"os"
	"testing"
)

type getSimplePromptTests struct {
	input, prompt string
}

type getPromptVerifyTests struct {
	input, prompt string
	result        bool
	validCases    map[string]bool
}

var gSPTests = []getSimplePromptTests{
	{"7", "How many days of the week are there?"},
	{"One two three", "Count to three"},
	{"365 days", "How many years are there in a year?"},
}

func TestGetSimplePrompts(t *testing.T) {

	for _, test := range gSPTests {

		content := []byte(test.input)
		tmpfile, err := os.CreateTemp("", "example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name())

		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}

		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}

		rdr := bufio.NewReader(tmpfile)

		answer := GetSimplePromptText(rdr, test.prompt)

		if answer != test.input {
			t.Errorf("Answer %s does not match %s expected input", answer, test.input)
		}
	}
}
