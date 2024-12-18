// first is the first of December.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/pfcm/aoc24/it"
)

func main() {
	left, right, err := read()
	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(left)
	sort.Ints(right)

	fmt.Println("part one:")
	partOne(left, right)

	fmt.Println("part two:")
	partTwo(left, right)
}

func partTwo(left, right []int) {
	counts := make(map[int]int, len(right))
	for _, r := range right {
		counts[r] = counts[r] + 1
	}

	total := 0
	for _, l := range left {
		total += l * counts[l]
	}
	fmt.Println(total)
}

func partOne(left, right []int) {
	total := 0
	for l, r := range it.Zip(slices.Values(left), slices.Values(right)) {
		d := l - r
		if d < 0 {
			d = -d
		}
		total += d
	}
	fmt.Println(total)
}

func read() (left, right []int, err error) {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := strings.Fields(scan.Text())
		if len(line) != 2 {
			return nil, nil, fmt.Errorf("invalid line %q", scan.Text())
		}
		l, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, nil, fmt.Errorf("parsing line %q: %w", scan.Text(), err)
		}
		r, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, nil, fmt.Errorf("parsing line %q: %w", scan.Text(), err)
		}
		left = append(left, l)
		right = append(right, r)
	}
	if err := scan.Err(); err != nil {
		return nil, nil, err
	}
	return
}
