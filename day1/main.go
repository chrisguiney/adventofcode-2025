package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	dialMax   = 99
	dialStart = 50
)

type rotation struct {
	dir   byte
	steps int
}

func readInput(fname string) ([]rotation, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed opening input file %s: %w", fname, err)
	}
	defer file.Close()
	var rotations []rotation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("failed converting %s to int: %w", line[1:], err)
		}
		rotations = append(rotations, rotation{dir: line[0], steps: steps})
	}
	return rotations, nil
}

func applyRotation(rotation rotation, dial, dialAtZero int) (int, int) {
	switch rotation.dir {
	case 'R':
		for i := 0; i < rotation.steps; i++ {
			dial++
			if dial > dialMax {
				dial = 0
			}
			if dial == 0 {
				dialAtZero++
			}
		}
	case 'L':
		for i := 0; i < rotation.steps; i++ {
			dial--
			if dial < 0 {
				dial = dialMax
			}
			if dial == 0 {
				dialAtZero++
			}
		}
	default:
		log.Fatalf("unknown direction %q", rotation.dir)
	}
	return dial, dialAtZero
}

func applyRotations(rotations []rotation, dial int) int {
	var dialAtZero = 0
	for _, rotation := range rotations {
		dial, dialAtZero = applyRotation(rotation, dial, dialAtZero)
	}
	return dialAtZero
}

func main() {
	rotations, err := readInput("./day1/input")
	if err != nil {
		log.Fatal(err)
	}

	var dialAtZero = applyRotations(rotations, dialStart)

	fmt.Println(dialAtZero)
}
