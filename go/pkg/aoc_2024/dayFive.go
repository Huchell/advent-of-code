package aoc2024

import (
	"bufio"
	"bytes"
	"huchell/aoc/pkg/assert"
	"math"
	"slices"
	"strconv"
)

type DayFive struct{}

func (DayFive) PartOne(input []byte) (any, error) {
	rules, pagesList := parseRulesAndPages(input)

	sum := 0
	for _, pages := range pagesList {
		if !isPageListInOrder(pages, rules) {
			continue
		}

		midIdx := int(math.Floor(float64(len(pages)) / 2))
		sum += pages[midIdx]
	}
	return sum, nil
}

func (DayFive) PartTwo(input []byte) (any, error) {
	rules, pagesList := parseRulesAndPages(input)

	sum := 0
	for _, pages := range pagesList {
		if isPageListInOrder(pages, rules) {
			continue
		}

		sortPages(pages, rules)
		midIdx := int(math.Floor(float64(len(pages)) / 2))
		sum += pages[midIdx]
	}
	return sum, nil
}

func parseRulesAndPages(input []byte) (RuleList, [][]int) {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	rules := make(RuleList)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}

		parts := bytes.Split(line, []byte("|"))
		left, err := strconv.Atoi(string(parts[0]))
		assert.ErrorIsNil(err)
		right, err := strconv.Atoi(string(parts[1]))
		assert.ErrorIsNil(err)

		addRule(rules, left, right)
	}

	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}

		parts := bytes.Split(line, []byte(","))

		pages := make([]int, len(parts))
		for idx, part := range parts {
			val, err := strconv.Atoi(string(part))
			assert.ErrorIsNil(err)
			pages[idx] = val
		}
		updates = append(updates, pages)
	}

	return rules, updates
}

type RuleList = map[int][]int

func addRule(rules RuleList, key, val int) {
	existing, ok := rules[key]
	if !ok {
		existing = []int{}
	}

	existing = append(existing, val)
	rules[key] = existing
}

type Rule struct {
	left, right int
}

func (rule Rule) applies(pages []int) bool {
	return slices.Contains(pages, rule.left) && slices.Contains(pages, rule.right)
}

func isPageListInOrder(pages []int, ruleList RuleList) bool {
	beforeSlice := []int{}
	for _, page := range pages {
		rules, ok := ruleList[page]
		if !ok {
			beforeSlice = append(beforeSlice, page)
			continue
		}

		for _, rule := range rules {
			if !slices.Contains(beforeSlice, rule) {
				continue
			}
			return false
		}
		beforeSlice = append(beforeSlice, page)
	}
	return true
}

func sortPages(pages []int, ruleList RuleList) {
	isSorted := false
	for !isSorted {
		for idx, page := range pages {
			rules, ok := ruleList[page]
			if !ok {
				continue
			}

			minIdx := idx
			for _, rule := range rules {
				ruleIdx := slices.Index(pages, rule)
				if ruleIdx == -1 || ruleIdx > idx {
					continue
				}
				minIdx = min(minIdx, ruleIdx)
			}

			swap := pages[minIdx]
			pages[idx] = swap
			pages[minIdx] = page
		}

		isSorted = isPageListInOrder(pages, ruleList)
	}
}
