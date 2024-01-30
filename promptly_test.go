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
	result, caseS bool
	validCases    []string
}

var gPVTests = []getPromptVerifyTests{
	{
		"6",
		"How many days of the week are there?",
		false,
		true,
		[]string{"7"},
	},
	{
		"7",
		"How many days of the week are there?",
		true,
		true,
		[]string{"7"},
	},
	{
		"Barbie",
		"Barbie or Oppenheimer?",
		true,
		true,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"Oppenheimer",
		"Barbie or Oppenheimer?",
		true,
		true,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"Super Mario Movie",
		"Barbie or Oppenheimer?",
		false,
		true,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"barbie",
		"Barbie or Oppenheimer?",
		true,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"oppenheimer",
		"Barbie or Oppenheimer?",
		true,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"super mario movie",
		"Barbie or Oppenheimer?",
		false,
		false,
		[]string{"Barbie", "Oppenheimer"},
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

		result, _ := GetPromptVerify(rdr, test.prompt, test.validCases, test.caseS)

		if result != test.result {
			t.Errorf("Test result %t does not match %t expected result", result, test.result)
		}
	}
}

type getPromptVerifyRegexTests struct {
	input, prompt, r string
	result           bool
}

var rThreeDigit string = `^\d{1,3}$`
var rAlphabetOnly string = `^[a-zA-Z]+$`

var gPVRTests = []getPromptVerifyRegexTests{
	{"28", "How old are you?", rThreeDigit, true},
	{"twenty-eight", "How old are you?", rThreeDigit, false},
	{"20 + eight", "How old are you?", rThreeDigit, false},
	{"", "How old are you?", rThreeDigit, false},
	{"KmanT", "What is your name?", rAlphabetOnly, true},
	{"28", "What is your name?", rAlphabetOnly, false},
	{"Bob Smith", "What is your name?", rAlphabetOnly, false},
	{"", "What is your name?", rAlphabetOnly, false},
}

func TestPromptVerifyRegex(t *testing.T) {
	for _, test := range gPVRTests {

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

		result, _, _ := GetPromptVerifyRegex(rdr, test.prompt, test.r)

		if result != test.result {
			t.Errorf(
				"Input %s resulting in %t does not match %t expected result",
				test.input,
				result,
				test.result,
			)
		}
	}
}
