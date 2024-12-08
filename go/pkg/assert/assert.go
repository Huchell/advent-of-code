package assert

import "log"

func ErrorIsNil(err error) {
	if err == nil {
		return
	}
	log.Fatal(err)
}

func IsEqual[T comparable](result, expected T) {
	if result == expected {
		return
	}
	log.Fatalf("%+v != %+v\n\tExpected: %+v\n\tResult: %+v", result, expected, expected, result)
}
