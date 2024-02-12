// Package promptly has helper functions to make CLI Application building easy.
//
// The promptly package should only be used for CLI applications that require
// terminal interaction.
package promptly

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const intR string = `^\d*$`
const floatR string = `^\d*(.\d*)?$`

// GetSimplePromptText gets a single line from the bufio Reader.
// It also removes the line break '\n' from the input.
// Use this if you do not require any validation.
func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Println(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

// GetPromptVerify gets a single line from the bufio Reader, and checks if the
// input is included in the list of valid input. GetPromptVerify returns a
// boolean based on its validity and the input received, whether or not the
// the user took a safe exit, and the input from the user.
func GetPromptVerify(rdr *bufio.Reader, prmpt, safeW string, vi []string, caseS bool) (valid, safeExit bool, input string) {
	in := GetSimplePromptText(rdr, prmpt)

	if strings.EqualFold(in, safeW) {
		return false, true, in
	}

	if !caseS {
		in = strings.ToLower(in)
		stringSliceToLower(&vi)
	}

	viMap := sliceToBoolMap[string](vi)

	return viMap[in], false, in
}

// GetPromptVerifyRegex verifies input against a regex. If an error is thrown,
// then GetPromptVerifyRegex returns false, an empty string, and an error.
// Otherwise, GetPromptVerifyRegex will return a boolean based on if the string
// matches the input, whether or not the user took the safe exit,
// the input as a string, and nil as the error.
func GetPromptVerifyRegex(rdr *bufio.Reader, prmpt, safeW, rS string) (valid, safeExit bool, input string, err error) {
	in := GetSimplePromptText(rdr, prmpt)
	r, err := regexp.Compile(rS)

	if err != nil {
		return false, false, "", errors.New("InvalidRegexError")
	}

	if strings.EqualFold(in, safeW) {
		return false, true, in, nil
	}

	return r.MatchString(in), false, in, nil
}

// GetPromptVerifyIntRange verifies that input is numeric, and that it fits in
// between the min and max int. There is also an option to make the this prompt
// inclusive (incl == true) or exclusive (incl == false)
func GetPromptVerifyIntRange(
	rdr *bufio.Reader,
	prmpt, safeW string,
	min, max int,
	incl bool,
) (valid, safeExit bool, input int, err error) {
	valid, safeExit, in, err := GetPromptVerifyRegex(rdr, prmpt, safeW, intR)

	if err != nil {
		return valid, safeExit, 0, err
	}

	if !valid || safeExit {
		return valid, safeExit, 0, nil
	}

	convIn, err := strconv.Atoi(in)
	if err != nil {
		return false, false, 0, err
	}

	return numericFitsInRange[int](&incl, &convIn, &min, &max)
}

// GetPromptVerifyIntRange verifies that input is numeric, and that it fits in
// between the min and max float32. There is also an option to make the this prompt
// inclusive (incl == true) or exclusive (incl == false)
func GetPromptVerifyFloat32Range(
	rdr *bufio.Reader,
	prmpt, safeW string,
	min, max float32,
	incl bool,
) (valid, safeExit bool, input float32, err error) {
	valid, safeExit, in, err := GetPromptVerifyRegex(rdr, prmpt, safeW, floatR)

	if err != nil {
		return valid, safeExit, 0, err
	}

	if !valid || safeExit {
		return valid, safeExit, 0, nil
	}

	pIn, err := strconv.ParseFloat(in, 32)
	if err != nil {
		return false, false, 0, err
	}

	convIn := float32(pIn)
	return numericFitsInRange[float32](&incl, &convIn, &min, &max)
}

// GetPromptVerifyIntRange verifies that input is numeric, and that it fits in
// between the min and max float64. There is also an option to make the this prompt
// inclusive (incl == true) or exclusive (incl == false)
func GetPromptVerifyFloat64Range(
	rdr *bufio.Reader,
	prmpt, safeW string,
	min, max float64,
	incl bool,
) (valid, safeExit bool, input float64, err error) {
	valid, safeExit, in, err := GetPromptVerifyRegex(rdr, prmpt, safeW, floatR)

	if err != nil {
		return valid, safeExit, 0, err
	}

	if !valid || safeExit {
		return valid, safeExit, 0, nil
	}

	convIn, err := strconv.ParseFloat(in, 32)
	if err != nil {
		return false, false, 0, err
	}

	return numericFitsInRange[float64](&incl, &convIn, &min, &max)
}
