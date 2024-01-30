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

type getPromptVerifyTests struct {
	input, prompt string
	result        bool
	validCases    map[string]bool
}

var gPVTests = []getPromptVerifyTests{
	{
		"6",
		"How many days of the week are there?",
		false,
		map[string]bool{"7": true},
	},
	{
		"7",
		"How many days of the week are there?",
		true,
		map[string]bool{"7": true},
	},
	{
		"Barbie",
		"Barbie or Oppenheimer?",
		true,
		map[string]bool{"Barbie": true, "Oppenheimer": true},
	},
	{
		"Oppenheimer",
		"Barbie or Oppenheimer?",
		true,
		map[string]bool{"Barbie": true, "Oppenheimer": true},
	},
	{
		"Super Mario Movie",
		"Barbie or Oppenheimer?",
		false,
		map[string]bool{"Barbie": true, "Oppenheimer": true},
	},
}

func TestPromptVerify(t *testing.T) {

	for _, test := range gPVTests {

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

		result, _ := GetPromptVerify(rdr, test.prompt, test.validCases)

		if result != test.result {
			t.Errorf("Test result %t does not match %t expected result", result, test.result)
		}
	}
}
