package promptly

import (
	"cmp"
	"strings"
)

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

// numericFitsInRange checks if input fits within a min and max. If incl is true,
// the check will be inclusive; if false it will be exclusive.
func numericFitsInRange[O cmp.Ordered](incl *bool, in, min, max *O) (bool, bool, O, error) {
	if *incl {
		return !(*in < *min || *in > *max), false, *in, nil
	} else {
		return !(*in <= *min || *in >= *max), false, *in, nil
	}
}
