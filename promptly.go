package promptly

import (
	"bufio"
	"fmt"
	"strings"
)

func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Println(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func GetPromptVerify(rdr *bufio.Reader, prmpt string, validInpt map[string]bool) (bool, string) {
	in := GetSimplePromptText(rdr, prmpt)

	return validInpt[in], in
}

func GetPromptVerifyLoop(rdr *bufio.Reader, prmpt string, validInpt map[string]bool) string {
	var in string
	isValid := false

	for !isValid {
		isValid, in = GetPromptVerify(rdr, prmpt, validInpt)

		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again", in)
		}
	}
	return in
}
