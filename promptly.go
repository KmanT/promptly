// Package promptly has helper functions to make CLI Application building easy.
//
// The promptly package should only be used for CLI applications that require
// terminal interaction.
package promptly

import (
	"bufio"
	"cmp"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

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
// boolean based on its validity and the input received.
func GetPromptVerify(rdr *bufio.Reader, prmpt string, vi []string, caseS bool) (bool, string) {
	in := GetSimplePromptText(rdr, prmpt)

	if !caseS {
		in = strings.ToLower(in)
		stringSliceToLower(&vi)
	}

	viMap := sliceToBoolMap[string](vi)

	return viMap[in], in
}

// GetPromptVerifyRegex verifies input against a regex. If an error is thrown,
// then GetPromptVerifyRegex returns false, an empty string, and an error.
// Otherwise, GetPromptVerifyRegex will return a boolean based on if the string
// matches the input, the input as a string, and nil as the error
func GetPromptVerifyRegex(rdr *bufio.Reader, prmpt string, rS string) (bool, string, error) {
	in := GetSimplePromptText(rdr, prmpt)
	r, err := regexp.Compile(rS)

	if err != nil {
		return false, "", errors.New("InvalidRegexError")
	}

	return r.MatchString(in), in, nil
}

// GetPromptVerifyLoop attempts to get input until the input is valid. If the
// input received is invalid, GetPromptVerify will loop again and ask for input
// again. Once the input is valid GetPromptVerify will return the user's choice.
func GetPromptVerifyLoop(rdr *bufio.Reader, prmpt string, vi []string, caseS bool) string {
	var in string
	isValid := false

	if !caseS {
		stringSliceToLower(&vi)
	}

	vIM := sliceToBoolMap[string](vi)

	for !isValid {
		in = GetSimplePromptText(rdr, prmpt)
		if !caseS {
			in = strings.ToLower(in)
		}

		isValid = vIM[in]
		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again", in)
		}
	}
	return in
}

// sliceToBoolMap is helper function that takes in a slice of types that
// implement cmp.Ordered (as O) and returns a map with the O type as a key
// and bool as the value. Effectively creating a set
func sliceToBoolMap[O cmp.Ordered](slice []O) map[O]bool {
	m := make(map[O]bool)

	for _, el := range slice {
		m[el] = true
	}

	return m
}

// stringSliceToLower is a helper function that mutates a string slice to
// lowercase.
func stringSliceToLower(slice *[]string) {
	for i, el := range *slice {
		(*slice)[i] = strings.ToLower(el)
	}
}
