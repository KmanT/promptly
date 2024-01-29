package promptly

import (
	"math"
	"reflect"
	"strconv"
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
	{"1234545", "int32"},
	{"0.0", "float32"},
	{"0.0", "float64"},
	{"foobar", "string"},

	{strconv.Itoa(math.MaxInt), "int"},
	{strconv.Itoa(math.MaxInt8), "int8"},
	{strconv.Itoa(math.MaxInt16), "int16"},
	{strconv.Itoa(math.MaxInt32), "int32"},
	{strconv.Itoa(math.MaxInt64), "int64"},

	{strconv.Itoa(math.MinInt), "int"},
	{strconv.Itoa(math.MinInt8), "int8"},
	{strconv.Itoa(math.MinInt16), "int16"},
	{strconv.Itoa(math.MinInt32), "int32"},
	{strconv.Itoa(math.MinInt64), "int64"},

	{strconv.FormatUint(uint64(math.MaxUint), 10), "uint"},
	{strconv.FormatUint(uint64(math.MaxUint8), 10), "uint8"},
	{strconv.FormatUint(uint64(math.MaxUint16), 10), "uint16"},
	{strconv.FormatUint(uint64(math.MaxUint32), 10), "uint32"},
	{strconv.FormatUint(uint64(math.MaxUint64), 10), "uint64"},

	{strconv.FormatUint(uint64(0), 10), "uint"},
	{strconv.FormatUint(uint64(0), 10), "uint8"},
	{strconv.FormatUint(uint64(0), 10), "uint16"},
	{strconv.FormatUint(uint64(0), 10), "uint32"},
	{strconv.FormatUint(uint64(0), 10), "uint64"},

	{strconv.FormatFloat(float64(math.MaxFloat32), 'f', 2, 32), "float32"},
	{strconv.FormatFloat(float64(math.MaxFloat64), 'f', 2, 64), "float64"},
	{strconv.FormatFloat(float64(0), 'f', 2, 32), "float32"},
	{strconv.FormatFloat(float64(0), 'f', 2, 64), "float64"},
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
