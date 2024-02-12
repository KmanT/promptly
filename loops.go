package promptly

import (
	"bufio"
	"fmt"
	"strings"
)

// GetPromptVerifyLoop attempts to get input until the input is valid. If the
// input received is invalid, GetPromptVerify will loop again and ask for input
// again. Once the input is valid GetPromptVerify will return whether or not
// the user took the safe exit and the user's choice.
func GetPromptVerifyLoop(rdr *bufio.Reader, prmpt, safeW string, vi []string, caseS bool) (safeExit bool, input string) {
	if !caseS {
		stringSliceToLower(&vi)
	}

	vIM := sliceToBoolMap[string](vi)

	for {
		in := GetSimplePromptText(rdr, prmpt)

		if strings.EqualFold(in, safeW) {
			return true, in
		}

		if !caseS {
			in = strings.ToLower(in)
		}

		isValid := vIM[in]
		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again\n", in)
		} else {
			return false, in
		}
	}
}

// GetPromptVerifyRegexLoop attempts to get input until the input is valid. If
// the regex 'rS' is invalid, the program will panic as the set up to this would
// be fundamentally incorrect and cause an infinite loop. Otherwise, it
// verifies the input. If the input is invalid, then it will attempt to get the
// input again. If the input is valid, then the input will be returned along
// with a safe-exit status of false. If the user used a safe word, then the
// safe exit state will be true and returned along with the input.
func GetPromptVerifyRegexLoop(rdr *bufio.Reader, prmpt, safeW string, rS string) (safeExit bool, input string) {
	for {
		isValid, safeE, in, err := GetPromptVerifyRegex(rdr, prmpt, safeW, rS)
		if err != nil {
			fmt.Printf("Regex '%s' is invalid", rS)
			panic(-1)
		}

		if safeE {
			return true, in
		}

		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again\n", in)
		} else {
			return false, in
		}
	}
}
