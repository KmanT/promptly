package promptly

import (
	"bufio"
	"log"
	"os"
	"testing"
)

type getSimplePromptTest struct {
	input, prompt string
}

var gSPTests = []getSimplePromptTest{
	{"7", "How many days of the week are there?"},
	{"One two three", "Count to three"},
	{"365 days", "How many years are there in a year?"},
}

func initTest(input *string) *os.File {
	content := []byte(*input)
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

	return tmpfile
}

func TestGetSimplePrompts(t *testing.T) {
	for _, test := range gSPTests {

		tmpfile := initTest(&test.input)

		rdr := bufio.NewReader(tmpfile)

		answer := GetSimplePromptText(rdr, test.prompt)

		if answer != test.input {
			t.Errorf("Answer %s does not match %s expected input", answer, test.input)
		}
	}
}

type getPromptVerifyTest struct {
	input, prompt, safeW string
	result, caseS, safeE bool
	validCases           []string
}

var gPVTests = []getPromptVerifyTest{
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

		tmpfile := initTest(&test.input)

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

type getPromptVerifyRegexTest struct {
	input, prompt, safeW, r string
	result, safeE           bool
}

var rThreeDigit string = `^\d{1,3}$`
var rAlphabetOnly string = `^[a-zA-Z]+$`

var gPVRTests = []getPromptVerifyRegexTest{
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

		tmpfile := initTest(&test.input)

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

type getPromptVerifyIntRangeTest struct {
	input, prompt, safeW    string
	inclusive, valid, safeE bool
	min, max, output        int
}

var gPVIRTests = []getPromptVerifyIntRangeTest{
	{"1", "Pick 1 - 10", "q", true, true, false, 1, 10, 1},
	{"+1", "Pick -1 - 10", "q", true, true, false, 1, 10, 1},
	{"-1", "Pick -1 - 10", "q", true, true, false, -1, 10, -1},
	{"5", "Pick 1 - 10", "q", true, true, false, 1, 10, 5},
	{"5", "Pick 1 - 10", "q", false, true, false, 1, 10, 5},
	{"10", "Pick 1 - 10", "q", true, true, false, 1, 10, 10},
	{"11", "Pick 1 - 10", "q", false, false, false, 1, 10, 11},
	{"0", "Pick 1 - 10", "q", false, false, false, 1, 10, 0},
	{"q", "Pick 1 - 10", "q", false, false, true, 1, 10, 0},
	{"Uh yeah", "Pick 1 - 10", "q", false, false, false, 1, 10, 0},
	{"Uh yeah 5", "Pick 1 - 10", "q", false, false, false, 1, 10, 0},
}

func TestPromptVerifyIntRange(t *testing.T) {
	for _, test := range gPVIRTests {

		tmpfile := initTest(&test.input)

		rdr := bufio.NewReader(tmpfile)

		result, safeExit, o, _ := GetPromptVerifyIntRange(rdr, test.prompt, test.safeW, test.min, test.max, test.inclusive)

		if safeExit != test.safeE {
			t.Errorf("Test safeExit status %t does not match the expected safeExit status %t", safeExit, test.safeE)
		}

		if result != test.valid {
			t.Errorf(
				"Input %s resulting in %t does not match %t expected result",
				test.input,
				result,
				test.valid,
			)
		}

		if o != test.output {
			t.Errorf("Output %d does not match the expected output %d", o, test.output)
		}
	}

}

type getPromptVerifyFloat32RangeTest struct {
	input, prompt, safeW    string
	inclusive, valid, safeE bool
	min, max, output        float32
}

var gPVF32RTests = []getPromptVerifyFloat32RangeTest{
	{"1.0", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 1.0},
	{"5.6", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 5.6},
	{"10.0", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 10.0},
	{"10.1", "Pick 1.0 - 10.0", "q", true, false, false, 1, 10, 10.1},
	{"0.999999", "Pick 1.0 - 10.0", "q", true, false, false, 1, 10, 0.999999},
	{"q", "Pick 1.0 - 10.0", "q", true, false, true, 1, 10, 0.0},
	{"1.0", "Pick 0.0 - 10.0", "q", false, true, false, 0, 10, 1.0},
	{"5.6", "Pick 0.0 - 10.0", "q", false, true, false, 0, 10, 5.6},
	{"10.0", "Pick 0.0 - 10.1", "q", false, true, false, 1, 10.1, 10.0},
	{"10.1", "Pick 1.0 - 10.1", "q", false, false, false, 1, 10.1, 10.1},
	{"0.999999", "Pick 0.999999 - 10.0", "q", false, false, false, 0.999999, 10, 0.999999},
	{"q", "Pick 1.0 - 10.0", "q", false, false, true, 1, 10, 0.0},
}

func TestPromptVerifyFloat32Range(t *testing.T) {
	for _, test := range gPVF32RTests {

		tmpfile := initTest(&test.input)

		rdr := bufio.NewReader(tmpfile)

		result, safeExit, o, _ := GetPromptVerifyFloat32Range(rdr, test.prompt, test.safeW, test.min, test.max, test.inclusive)

		if safeExit != test.safeE {
			t.Errorf("Test safeExit status %t does not match the expected safeExit status %t", safeExit, test.safeE)
		}

		if result != test.valid {
			t.Errorf(
				"Input %s resulting in %t does not match %t expected result",
				test.input,
				result,
				test.valid,
			)
		}

		if o != test.output {
			t.Errorf("Output %f does not match the expected output %f", o, test.output)
		}
	}
}

type getPromptVerifyFloat64RangeTest struct {
	input, prompt, safeW    string
	inclusive, valid, safeE bool
	min, max, output        float64
}

var gPVF64RTests = []getPromptVerifyFloat64RangeTest{
	{"1.0", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 1.0},
	{"5.6", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 5.6},
	{"10.0", "Pick 1.0 - 10.0", "q", true, true, false, 1, 10, 10.0},
	{"10.1", "Pick 1.0 - 10.0", "q", true, false, false, 1, 10, 10.1},
	{"0.999999", "Pick 1.0 - 10.0", "q", true, false, false, 1, 10, 0.999999},
	{"q", "Pick 1.0 - 10.0", "q", true, false, true, 1, 10, 0.0},
	{"1.0", "Pick 0.0 - 10.0", "q", false, true, false, 0, 10, 1.0},
	{"5.6", "Pick 0.0 - 10.0", "q", false, true, false, 0, 10, 5.6},
	{"10.0", "Pick 0.0 - 10.1", "q", false, true, false, 1, 10.1, 10.0},
	{"10.1", "Pick 1.0 - 10.1", "q", false, false, false, 1, 10.1, 10.1},
	{"0.999999", "Pick 0.999999 - 10.0", "q", false, false, false, 0.999999, 10, 0.999999},
	{"q", "Pick 1.0 - 10.0", "q", false, false, true, 1, 10, 0.0},
}

func TestPromptVerifyFloat64Range(t *testing.T) {
	for _, test := range gPVF64RTests {

		tmpfile := initTest(&test.input)

		rdr := bufio.NewReader(tmpfile)

		result, safeExit, o, _ := GetPromptVerifyFloat64Range(rdr, test.prompt, test.safeW, test.min, test.max, test.inclusive)

		if safeExit != test.safeE {
			t.Errorf("Test safeExit status %t does not match the expected safeExit status %t", safeExit, test.safeE)
		}

		if result != test.valid {
			t.Errorf(
				"Input %s resulting in %t does not match %t expected result",
				test.input,
				result,
				test.valid,
			)
		}

		if o != test.output {
			t.Errorf("Output %f does not match the expected output %f", o, test.output)
		}
	}
}
