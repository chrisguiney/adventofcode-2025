package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type idRange struct {
	lower, upper int
}

func parseInputs(file string) ([]idRange, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}
	s := string(b)

	var idRanges []idRange
	for _, r := range strings.Split(s, ",") {
		rr := strings.Split(strings.TrimSpace(r), "-")
		lower, err := strconv.ParseInt(rr[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse lower: %w", err)
		}
		upper, err := strconv.ParseInt(rr[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse upper: %w", err)
		}

		idRanges = append(idRanges, idRange{
			lower: int(lower),
			upper: int(upper),
		})
	}
	return idRanges, nil
}

func isInvalidID(i int) bool {
	s := fmt.Sprint(i)

	for i := 0; i < len(s)/2; i++ {
		pattern := fmt.Sprintf("^(%s)+$", s[:i+1])
		if match, _ := regexp.MatchString(pattern, s); match {
			return true
		}
	}

	return false
}

func findInvalidIDs(idRanges []idRange) []int {
	var invalidRanges []int
	for _, r := range idRanges {
		for i := r.lower; i <= r.upper; i++ {
			if isInvalidID(i) {
				invalidRanges = append(invalidRanges, i)
			}
		}
	}
	return invalidRanges
}

func sumInvalidIDs(ids []int) int {
	var sum int
	for _, i := range ids {
		sum += i
	}
	return sum
}

func main() {
	ranges, err := parseInputs("./day2/input")
	if err != nil {
		log.Fatalf("failed to parse inputs: %v", err)
	}

	invalid := findInvalidIDs(ranges)
	fmt.Println(sumInvalidIDs(invalid))
}
