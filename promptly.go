package promptly

import (
	"bufio"
	"fmt"

	//"os"
	"regexp"
	"strconv"
	"strings"
)

func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Print(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func GetTypeAndBits(typeName string) (string, int) {
	if !regexp.MustCompile(`[a-z]+\d+`).MatchString(typeName) {
		return typeName, 0
	}

	name := regexp.MustCompile(`[a-z]+`).FindString(typeName)
	bitStr := regexp.MustCompile(`\d+`).FindString(typeName)

	bits, err := strconv.Atoi(bitStr)
	if err != nil {
		fmt.Printf("Warning: could not find bits in %q", typeName)
		bits = 0
	}

	return name, bits
}

func ConvertInputToType(input *string, t string) (interface{}, error) {
	name, bits := GetTypeAndBits(t)

	switch name {
	case "int":
		if bits <= 0 {
			return strconv.Atoi(*input)
		} else {
			return strconv.ParseInt(*input, 10, bits)
		}
	case "float":
		return strconv.ParseFloat(*input, bits)
	case "rune":
		return strconv.ParseInt(*input, 10, 32)
	case "complex":
		return strconv.ParseComplex(*input, bits)
	default:
		return *input, nil
	}
}
