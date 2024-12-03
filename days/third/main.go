// third is day three.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	input := string(raw)

	fmt.Println("part one:")
	if err := partOne(input); err != nil {
		log.Fatal(err)
	}

	fmt.Println("part two:")
	if err := partTwo(input); err != nil {
		log.Fatal(err)
	}
}

// partOne solves the first part of the problem in nearly the worst possible way.
func partOne(in string) error {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0
	for _, m := range re.FindAllStringSubmatch(in, -1) {
		if len(m) != 3 {
			return fmt.Errorf("impossible match? %v", m)
		}
		x, err := strconv.Atoi(m[1])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}
		sum += x * y
	}
	fmt.Printf("sum of valid muls: %v\n", sum)
	return nil
}

// partTwo solves the second part of the problem, and is possibly even
// more gross than part one.
func partTwo(in string) error {
	var (
		re = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

		sum = 0
		do  = true
	)
	for _, m := range re.FindAllStringSubmatch(in, -1) {
		switch m[0] {
		case "do()":
			do = true
			continue
		case "don't()":
			do = false
			continue
		}
		if !do {
			continue
		}
		if len(m) != 3 {
			return fmt.Errorf("impossible match? %v", m)
		}
		x, err := strconv.Atoi(m[1])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}
		sum += x * y
	}
	fmt.Printf("sum of valid, enabled muls: %d\n", sum)
	return nil
}
