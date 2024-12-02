// second is December the second.
package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reports, err := read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part 1:")
	partOne(reports)
	fmt.Println("part 2:")
	partTwo(reports)
}

func partOne(reports [][]int) {
	safe := 0
	for _, r := range reports {
		positive := false
		if r[1] < r[0] {
			positive = true
		}
		if all(difference(slices.Values(r)), func(i int) bool {
			// This feels like it should be simpler.
			switch {
			case positive && i < 0:
				return false
			case !positive && i > 0:
				return false
			}
			if i < 0 {
				i = -i
			}
			return i >= 1 && i <= 3
		}) {
			safe++
		}
	}
	fmt.Printf("%d safe\n", safe)
}

func partTwo(reports [][]int) {
	safe := func(r func() iter.Seq[int]) bool {
		positive := false
		for d := range difference(r()) {
			if d > 0 {
				positive = true
			}
			break
		}
		return all(difference(r()), func(i int) bool {
			switch {
			case positive && i < 0:
				return false
			case !positive && i > 0:
				return false
			}
			if i < 0 {
				i = -i
			}
			return i >= 1 && i <= 3
		})
	}
	total := 0
	for _, r := range reports {
		if safe(func() iter.Seq[int] { return slices.Values(r) }) {
			total++
			continue
		}
		for i := range r {
			if safe(func() iter.Seq[int] {
				return skip(slices.Values(r), i)
			}) {
				total++
				break
			}
		}
	}
	fmt.Printf("%d safe\n", total)
}

func all[T any](s iter.Seq[T], p func(T) bool) bool {
	for v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}

func skip[T any](seq iter.Seq[T], i int) iter.Seq[T] {
	return func(yield func(T) bool) {
		j := 0
		for t := range seq {
			if j == i {
				j++
				continue
			}
			j++
			if !yield(t) {
				return
			}
		}
	}
}

func difference(is iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		var (
			prev  = 0
			first = true
		)
		for i := range is {
			if first {
				first = false
				prev = i
				continue
			}
			if !yield(prev - i) {
				return
			}
			prev = i
		}
	}
}

func read() ([][]int, error) {
	var (
		results [][]int
		scan    = bufio.NewScanner(os.Stdin)
	)
	for scan.Scan() {
		l := strings.Split(scan.Text(), " ")
		row := make([]int, len(l))
		for i := range l {
			level, err := strconv.Atoi(l[i])
			if err != nil {
				return nil, err
			}
			row[i] = level
		}
		results = append(results, row)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
