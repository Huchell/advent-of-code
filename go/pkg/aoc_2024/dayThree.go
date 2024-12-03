package aoc2024

import (
	"bufio"
	"bytes"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

type DayThree struct{}

var mulRegex regexp.Regexp = *regexp.MustCompile(`mul\(\d+,\d+\)`)
var doRegex regexp.Regexp = *regexp.MustCompile(`do\(\)`)
var donotRegex regexp.Regexp = *regexp.MustCompile(`don't\(\)`)

var mulPrefix []byte = []byte("mul(")
var mulSuffix []byte = []byte(")")
var comma []byte = []byte(",")

func (DayThree) PartOne(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(tokenizePartOne)

	sum := 0
	for scanner.Scan() {
		token := scanner.Bytes()
		if len(token) == 0 {
			continue
		}

		token, _ = bytes.CutPrefix(token, mulPrefix)
		token, _ = bytes.CutSuffix(token, mulSuffix)

		numSlices := bytes.SplitN(token, comma, 2)
		first, err := strconv.Atoi(string(numSlices[0]))
		if err != nil {
			return -1, err
		}
		second, err := strconv.Atoi(string(numSlices[1]))
		if err != nil {
			return -1, err
		}

		sum += first * second
	}
	return sum, nil
}

func (DayThree) PartTwo(input []byte) (any, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(tokenizePartTwo)

	active := true
	sum := 0
	for scanner.Scan() {
		token := scanner.Bytes()
		if len(token) == 0 {
			continue
		}

		switch len(token) {
		case 4: // do()
			active = true
			break
		case 7: // don't()
			active = false
			break
		default: //mul
			if !active {
				continue
			}

			token, _ = bytes.CutPrefix(token, mulPrefix)
			token, _ = bytes.CutSuffix(token, mulSuffix)

			numSlices := bytes.SplitN(token, comma, 2)
			first, err := strconv.Atoi(string(numSlices[0]))
			if err != nil {
				return -1, err
			}
			second, err := strconv.Atoi(string(numSlices[1]))
			if err != nil {
				return -1, err
			}

			sum += first * second
			break
		}
	}
	return sum, nil
}

func tokenizePartOne(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for advance = 0; advance < len(data); advance++ {
		switch data[advance] {
		case 'm':
			mulAdvance, mulToken, found := mulToken(data, advance)
			if !found {
				continue
			}
			token = mulToken
			advance += mulAdvance
			return
		}
	}

	if !atEOF {
		return 0, nil, nil
	}

	err = bufio.ErrFinalToken
	return
}

func tokenizePartTwo(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for advance = 0; advance < len(data); advance++ {
		switch data[advance] {
		case 'm':
			mulAdvance, mulToken, found := mulToken(data, advance)
			if !found {
				continue
			}
			token = mulToken
			advance += mulAdvance
			return

		case 'd':
			tokenAdvance, toggleToken, found := toggleToken(data, advance)
			if !found {
				continue
			}

			token = toggleToken
			advance += tokenAdvance
			return
		}
	}

	if !atEOF {
		return 0, nil, nil
	}

	err = bufio.ErrFinalToken
	return
}

func mulToken(data []byte, startIdx int) (advance int, token []byte, found bool) {
	advance = len(mulPrefix)
	if startIdx+advance >= len(data) {
		return 0, nil, false
	}

	token = data[startIdx : startIdx+advance]
	if !slices.Equal(token, mulPrefix) {
		found = false
		return
	}

	_, firstNum, found := numToken(data, startIdx+advance)
	advance += len(firstNum)
	if !found {
		found = false
		return
	}
	token = append(token, firstNum...)

	if data[startIdx+advance] != byte(',') {
		found = false
		return
	}
	token = append(token, data[startIdx+advance])
	advance += 1

	_, secondNum, found := numToken(data, startIdx+advance)
	advance += len(secondNum)
	if !found {
		found = false
		return
	}
	token = append(token, secondNum...)

	if data[startIdx+advance] != byte(')') {
		found = false
		return
	}
	token = append(token, byte(')'))
	advance += 1

	found = true
	return
}

func numToken(data []byte, startIdx int) (advance int, token []byte, found bool) {
	token = []byte{}
	for advance := 0; startIdx+advance < len(data); advance++ {
		b := rune(data[startIdx+advance])
		if !unicode.IsDigit(b) {
			break
		}

		found = true
		token = append(token, byte(b))
	}
	return
}

var doPrefix []byte = []byte("do")
var donotSuffix []byte = []byte("n't")
var doSuffix []byte = []byte("()")

func toggleToken(data []byte, startIdx int) (advance int, token []byte, found bool) {
	advance = len(doPrefix)
	if startIdx+advance >= len(data) {
		return 0, nil, false
	}

	token = data[startIdx : startIdx+advance]
	if !slices.Equal(token, doPrefix) {
		found = false
		return
	}

	donotSuffixToken := data[startIdx+advance : startIdx+advance+len(donotSuffix)]
	if slices.Equal(donotSuffixToken, donotSuffix) {
		token = append(token, donotSuffixToken...)
		advance += len(donotSuffixToken)
	}

	doSuffixToken := data[startIdx+advance : startIdx+advance+len(doSuffix)]
	if !slices.Equal(doSuffixToken, doSuffix) {
		return
	}
	token = append(token, doSuffixToken...)
	advance++

	found = true
	return
}
