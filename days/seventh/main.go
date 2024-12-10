// seventh is day seven.
package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	equations, err := read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part one:")
	partOne(equations)

	fmt.Println("part two:")
	partTwo(equations)
}

func partTwo(eqs []equation) {
	sum := uint64(0)
	for _, e := range eqs {
		for ops := range possibilities([]func(uint64, uint64) uint64{
			func(a, b uint64) uint64 { return a + b },
			func(a, b uint64) uint64 { return a * b },
			func(a, b uint64) uint64 {
				c, _ := strconv.ParseUint(fmt.Sprintf("%d%d", a, b), 10, 64)
				return c
			},
		}, len(e.inputs)-1) {
			total := e.inputs[0]
			for i, op := range enumerate(ops) {
				total = op(total, e.inputs[i+1])
			}
			if total == e.result {
				sum += total
				break
			}
		}
	}
	fmt.Printf("sum of equations that could work: %d\n", sum)
}

func partOne(eqs []equation) {
	// hilarious brute force
	sum := uint64(0)
	for _, e := range eqs {
		for ops := range possibilities([]func(uint64, uint64) uint64{
			func(a, b uint64) uint64 { return a + b },
			func(a, b uint64) uint64 { return a * b },
		}, len(e.inputs)-1) {
			total := e.inputs[0]
			for i, op := range enumerate(ops) {
				total = op(total, e.inputs[i+1])
			}
			if total == e.result {
				sum += total
				break
			}
		}
	}
	fmt.Printf("sum of equations that could work: %d\n", sum)
}

func enumerate[T any](is iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		j := 0
		for i := range is {
			if !yield(j, i) {
				return
			}
			j++
		}
	}
}

func fmap[A, B any](i iter.Seq[A], f func(A) B) iter.Seq[B] {
	return func(yield func(B) bool) {
		for a := range i {
			if !yield(f(a)) {
				return
			}
		}
	}
}

func possibilities[T any](values []T, length int) iter.Seq[iter.Seq[T]] {
	return func(yield func(iter.Seq[T]) bool) {
		var (
			n    = len(values)
			idcs = make([]int, length)
		)
		for {
			if !yield(func(yield func(T) bool) {
				for _, i := range idcs {
					if !yield(values[i]) {
						return
					}
				}
			}) {
				return
			}
			i := len(idcs) - 1
			for i >= 0 {
				idcs[i] = (idcs[i] + 1) % n
				if idcs[i] != 0 {
					break
				}
				i--
			}
			if i == -1 && idcs[0] == 0 {
				break
			}
		}
	}
}

func read() ([]equation, error) {
	var (
		scan      = bufio.NewScanner(os.Stdin)
		equations []equation
	)
	for scan.Scan() {
		e, err := newEquation(scan.Text())
		if err != nil {
			return nil, err
		}
		equations = append(equations, e)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	return equations, nil
}

type equation struct {
	result uint64
	inputs []uint64
}

func newEquation(s string) (equation, error) {
	r, rem, ok := strings.Cut(s, ": ")
	if !ok {
		return equation{}, fmt.Errorf("invalid equation %q", s)
	}
	var (
		e   equation
		err error
	)
	e.result, err = strconv.ParseUint(r, 10, 64)
	if err != nil {
		return equation{}, err
	}
	for _, num := range strings.Split(rem, " ") {
		i, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			return equation{}, err
		}
		e.inputs = append(e.inputs, i)
	}
	return e, nil
}
