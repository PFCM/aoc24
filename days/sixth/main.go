// sixth is the sixth day.
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func main() {
	m, err := read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part one:")
	partOne(m)

	fmt.Println("part two:")
	partTwo(m)
}

func loops(m [][]byte, g guard) bool {
	var (
		seen = map[guard]bool{g: true}
		ok   bool
	)
	tick := func() bool {
		m[g.pos.y][g.pos.x] = 'X'
		g, ok = tick(m, g)
		if !ok {
			return false
		}
		m[g.pos.y][g.pos.x] = '^'
		return true
	}
	for tick() {
		if seen[g] {
			return true
		}
		seen[g] = true
	}
	return false
}

func partTwo(m [][]byte) {
	var g guard
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '^' {
				g.pos = pair{x, y}
			}
		}
	}
	g.dir = pair{0, -1}

	numLoops := 0
	for i := range m {
		for j := range m[i] {
			switch m[i][j] {
			case '#':
				continue
			case '^':
				continue
			}
			fmt.Printf("\r%03d, %03d", i, j)
			m2 := deepCopy(m)
			m2[i][j] = '#'
			if loops(m2, g) {
				numLoops++
			}
		}
	}
	fmt.Printf("\n%d possible loops\n", numLoops)
}

type guard struct {
	pos, dir pair
}

func (g guard) next() pair {
	return pair{g.pos.x + g.dir.x, g.pos.y + g.dir.y}
}

func tick(m [][]byte, g guard) (guard, bool) {
	next := g.next()
	if next.x < 0 || next.x >= len(m[0]) {
		return g, false
	}
	if next.y < 0 || next.y >= len(m) {
		return g, false
	}
	if m[next.y][next.x] == '#' {
		g.dir = rot90(g.dir)
		return g, true
	}
	g.pos = next
	return g, true
}

func partOne(m [][]byte) {
	m = deepCopy(m)
	var g guard
	for y := range m {
		for x := range m[y] {
			if m[y][x] == '^' {
				g.pos = pair{x, y}
			}
		}
	}
	g.dir = pair{0, -1}

	var ok bool
	tick := func() bool {
		m[g.pos.y][g.pos.x] = 'X'
		g, ok = tick(m, g)
		if !ok {
			return false
		}
		m[g.pos.y][g.pos.x] = '^'
		return true
	}
	// print(m)
	for tick() {
		// print(m)
	}

	exes := 0
	for _, r := range m {
		for _, b := range r {
			if b == 'X' {
				exes++
			}
		}
	}
	fmt.Printf("%d unique positions\n", exes)
}

func print(m [][]byte) {
	fmt.Println()
	for _, row := range m {
		fmt.Printf("%s\n", row)
	}
}

func rot90(d pair) pair {
	switch d {
	case pair{0, -1}:
		return pair{1, 0}
	case pair{0, 1}:
		return pair{-1, 0}
	case pair{-1, 0}:
		return pair{0, -1}
	case pair{1, 0}:
		return pair{0, 1}
	}
	panic("hopefully unreachable")
}

type pair struct {
	x, y int
}

func deepCopy(bs [][]byte) [][]byte {
	out := slices.Clone(bs)
	for i, bs := range out {
		out[i] = slices.Clone(bs)
	}
	return out
}

func read() ([][]byte, error) {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	b := bytes.Split(raw, []byte("\n"))
	if len(b[len(b)-1]) == 0 {
		b = b[:len(b)-1]
	}
	return b, nil
}
