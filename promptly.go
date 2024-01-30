package promptly

import (
	"bufio"
	"cmp"
	"fmt"
	"strings"
)

func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Println(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func GetPromptVerify(rdr *bufio.Reader, prmpt string, vi []string, caseS bool) (bool, string) {
	in := GetSimplePromptText(rdr, prmpt)

	if !caseS {
		in = strings.ToLower(in)
		stringSliceToLower(&vi)
	}

	viMap := sliceToBoolMap[string](vi)

	return viMap[in], in
}

// TODO: Verify with Regex func

func GetPromptVerifyLoop(rdr *bufio.Reader, prmpt string, vi []string, caseS bool) string {
	var in string
	isValid := false

	if !caseS {
		stringSliceToLower(&vi)
	}

	vIM := sliceToBoolMap[string](vi)

	for !isValid {
		// isValid, in = GetPromptVerify(rdr, prmpt, vi, caseS)
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

func sliceToBoolMap[O cmp.Ordered](slice []O) map[O]bool {
	m := make(map[O]bool)

	for _, el := range slice {
		m[el] = true
	}

	return m
}

func stringSliceToLower(slice *[]string) {
	for i, el := range *slice {
		(*slice)[i] = strings.ToLower(el)
	}
}
