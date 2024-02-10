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
	input, prompt, safeW string
	result, caseS, safeE bool
	validCases           []string
}

var gPVTests = []getPromptVerifyTests{
	{
		"6",
		"How many days of the week are there?",
		"q",
		false,
		true,
		false,
		[]string{"7"},
	},
	{
		"7",
		"How many days of the week are there?",
		"q",
		true,
		true,
		false,
		[]string{"7"},
	},
	{
		"q",
		"How many days of the week are there?",
		"q",
		false,
		false,
		true,
		[]string{"7"},
	},
	{
		"Barbie",
		"Barbie or Oppenheimer?",
		"q",
		true,
		true,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"Oppenheimer",
		"Barbie or Oppenheimer?",
		"q",
		true,
		true,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"Super Mario Movie",
		"Barbie or Oppenheimer?",
		"q",
		false,
		true,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"barbie",
		"Barbie or Oppenheimer?",
		"q",
		true,
		false,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"oppenheimer",
		"Barbie or Oppenheimer?",
		"q",
		true,
		false,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"super mario movie",
		"Barbie or Oppenheimer?",
		"q",
		false,
		false,
		false,
		[]string{"Barbie", "Oppenheimer"},
	},
	{
		"q",
		"Barbie or Oppenheimer?",
		"q",
		false,
		false,
		true,
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

		result, safeExit, _ := GetPromptVerify(rdr, test.prompt, test.safeW, test.validCases, test.caseS)

		if safeExit != test.safeE {
			t.Errorf(
				"Test safeExit status %t does not match the expected safeExit status %t for input %s",
				safeExit,
				test.safeE,
				test.input,
			)
		}

		if result != test.result {
			t.Errorf("result %t safeExit %t", result, safeExit)
			t.Errorf("test.result %t test.safeE %t", test.result, test.safeE)
			t.Errorf("Test result %t does not match %t expected result for input %s",
				result,
				test.result,
				test.input,
			)
		}

	}
}

type getPromptVerifyRegexTests struct {
	input, prompt, safeW, r string
	result, safeE           bool
}

var rThreeDigit string = `^\d{1,3}$`
var rAlphabetOnly string = `^[a-zA-Z]+$`

var gPVRTests = []getPromptVerifyRegexTests{
	{"28", "How old are you?", "q", rThreeDigit, true, false},
	{"twenty-eight", "How old are you?", "q", rThreeDigit, false, false},
	{"20 + eight", "How old are you?", "q", rThreeDigit, false, false},
	{"q", "How old are you?", "q", rThreeDigit, false, true},
	{"", "How old are you?", "q", rThreeDigit, false, false},
	{"KmanT", "What is your name?", "q", rAlphabetOnly, true, false},
	{"28", "What is your name?", "q", rAlphabetOnly, false, false},
	{"Bob Smith", "What is your name?", "q", rAlphabetOnly, false, false},
	{"", "What is your name?", "q", rAlphabetOnly, false, false},
	{"q", "What is your name?", "q", rAlphabetOnly, false, true},
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

		result, safeExit, _, _ := GetPromptVerifyRegex(rdr, test.prompt, test.safeW, test.r)

		if safeExit != test.safeE {
			t.Errorf("Test safeExit status %t does not match the expected safeExit status %t", safeExit, test.safeE)
		}

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

// verify int range

// verify float32 range

// verify float64 range
