package promptly

import (
	"reflect"
	"testing"
)

type getTypeAndBitsCase struct {
	arg, typeName string
	bits          int
}

var tAndBTests = []getTypeAndBitsCase{
	{"string", "string", 0},
	{"int", "int", 0},
	{"int16", "int", 16},
	{"int32", "int", 32},
	{"float32", "float", 32},
	{"complex128", "complex", 128},
}

func TestGetTypeAndBits(t *testing.T) {

	for _, test := range tAndBTests {
		name, bits, _ := GetTypeAndBits(test.arg)

		if name != test.typeName || bits != test.bits {
			t.Errorf(
				"Output %q and %d not equal to %q and %d",
				name,
				bits,
				test.typeName,
				test.bits,
			)
		}
	}
}

type convertInputToTypeTest struct {
	i, t string
}

var cItTCases = []convertInputToTypeTest{
	{"16", "int"},
	{"-1234", "int"},
	{"1234545", "int64"},
	{"foobar", "string"},
}

func TestConvertInputToType(t *testing.T) {
	for _, test := range cItTCases {

		output, err := ConvertInputToType(&test.i, test.t)

		if err != nil {
			t.Errorf("Could not convert %q to %q", test.i, test.t)
		}

		outT := reflect.TypeOf(output).String()

		if outT != test.t {
			t.Errorf(
				"Output %v could not be be converted to %q, but instead %q",
				output,
				test.t,
				outT,
			)
		}
	}
}
