package promptly

import (
	"bufio"
	"errors"
	"fmt"

	//"os"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetSimplePromptText(rdr *bufio.Reader, prmpt string) string {
	fmt.Print(prmpt)
	txt, _ := rdr.ReadString('\n')
	return strings.TrimRight(txt, "\n")
}

func GetTypeAndBits(typeName string) (string, int, error) {
	if !regexp.MustCompile(`[a-z]+\d{1,3}`).MatchString(typeName) {
		return typeName, 0, nil
	}

	name := regexp.MustCompile(`[a-z]+`).FindString(typeName)
	bitStr := regexp.MustCompile(`\d{1,3}`).FindString(typeName)

	bits, err := strconv.Atoi(bitStr)
	if err != nil {
		fmt.Printf("Warning: could not find bits in %q", typeName)
		return name, 0, nil
	}

	if bits < 8 {
		fmt.Print("Error: the number of bits must be greater than 8.")
		return name, 0, errors.New("BitOutOfRangeError")
	}

	if bits >= 128 {
		fmt.Print("Error: the number of bits must be less than or equal to 128.")
		return name, 0, errors.New("BitOutOfRangeError")
	}

	boL2 := math.Log2(float64(bits))

	if math.Mod(boL2, 1.0) != 0.0 {
		fmt.Print("Error: the number of bits must be a power of 2.")
		return name, 0, errors.New("BitNotPowerOfTwoError")
	}

	return name, bits, nil
}

func ConvertInputToType(input *string, t string) (interface{}, error) {
	name, bits, _ := GetTypeAndBits(t)

	switch name {
	case "int":
		if bits <= 0 {
			return strconv.Atoi(*input)
		} else {
			return stringToInt(input, &bits)
		}
	case "uint":
		return stringToUint(input, &bits)
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

func stringToInt(input *string, bits *int) (interface{}, error) {
	o, err := strconv.ParseInt(*input, 10, *bits)
	if err != nil {
		fmt.Printf("Error: input %s is not valid for int conversion", *input)
		return 0, err
	}

	if *bits >= 64 {
		return o, nil
	}

	switch *bits {
	case 8:
		var cI int8 = int8(o)
		return cI, nil
	case 16:
		var cI int16 = int16(o)
		return cI, nil
	case 32:
		var cI int32 = int32(o)
		return cI, nil
	default:
		return o, nil
	}
}

func stringToUint(input *string, bits *int) (interface{}, error) {
	o, err := strconv.ParseUint(*input, 10, *bits)
	if err != nil {
		fmt.Printf("Error: input %s is not valid for unit conversion", *input)
		return 0, err
	}

	if *bits >= 64 {
		return o, nil
	}

	switch *bits {
	case 0:
		var cI uint = uint(o)
		return cI, nil
	case 8:
		var cI uint8 = uint8(o)
		return cI, nil
	case 16:
		var cI uint16 = uint16(o)
		return cI, nil
	case 32:
		var cI uint32 = uint32(o)
		return cI, nil
	default:
		return o, nil
	}
}

// TODO: prompt verification
