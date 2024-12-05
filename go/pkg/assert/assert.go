package assert

import "testing"

func ErrorIsNil(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Fatal(err)
}

func IsEqual[T comparable](t *testing.T, result, expected T) {
	if result == expected {
		return
	}
	t.Fatalf("%+v != %+v\n\tExpected: %+v\n\tResult: %+v", result, expected, expected, result)
}
