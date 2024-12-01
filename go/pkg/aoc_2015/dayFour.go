package aoc2015

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type DayFour struct{}

func (DayFour) PartOne(input []byte) (any, error) {
	val := 1
	var checksum string
	for !strings.HasPrefix(checksum, "00000") {
		valBytes := []byte(strconv.Itoa(val))
		checkValue := bytes.Join([][]byte{input, valBytes}, []byte{})
		checksum = fmt.Sprintf("%x", md5.Sum(checkValue))

		val += 1
	}

	return val - 1, nil
}

func (DayFour) PartTwo(input []byte) (any, error) {
	val := 1
	var checksum string
	for !strings.HasPrefix(checksum, "000000") {
		valBytes := []byte(strconv.Itoa(val))
		checkValue := bytes.Join([][]byte{input, valBytes}, []byte{})
		checksum = fmt.Sprintf("%x", md5.Sum(checkValue))

		val += 1
	}

	return val - 1, nil
}
