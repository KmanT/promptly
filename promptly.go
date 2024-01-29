package promptly

import (
	"KmanT/qwikvert"

	"bufio"
	"fmt"
	"reflect"
	"strings"
)

func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Print(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func GetPromptVerify(rdr *bufio.Reader, prmpt string, validInpt map[interface{}]bool) string {
	var in string
	isValid := false
	convType := reflect.TypeOf(validInpt).Key().String()

	for !isValid {
		in = GetSimplePromptText(rdr, prmpt)

		isValid, _ = validateInput(in, convType, validInpt)

		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again", in)
		}
	}

	return in
}

func GetPromptVerifyConvert(rdr *bufio.Reader, prmpt string, validInpt map[interface{}]bool) interface{} {
	var converted interface{}
	isValid := false
	convType := reflect.TypeOf(validInpt).Key().String()

	for !isValid {
		in := GetSimplePromptText(rdr, prmpt)

		isValid, converted = validateInput(in, convType, validInpt)

		if !isValid {
			fmt.Printf("Input '%s' is invalid. Try again", in)
		}
	}

	return converted
}

func validateInput(in, convType string, validInpt map[interface{}]bool) (bool, interface{}) {
	var convIn interface{}
	convIn, err := qwikvert.ConvertInputToType(&in, convType)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Try again")
		return false, -1
	}

	return validInpt[convIn], convIn
}
