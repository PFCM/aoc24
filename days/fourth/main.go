// fourth is day four.
package main

import (
	"bytes"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
	"slices"

	"github.com/pfcm/aoc24/it"
)

func main() {
	in, err := read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part one:")
	partOne(in)

	fmt.Println("part two:")
	partTwo(in)
}

func partOne(in [][]byte) {
	total := 0
	for x := range in {
		for y := range in[x] {
			for d := range directions(in, x, y) {
				if startsWith(d, []byte("XMAS")) {
					total++
				}
			}
		}
	}
	fmt.Printf("total XMAS: %d\n", total)
}

func partTwo(in [][]byte) {
	total := 0
	mas, sam := []byte("MAS"), []byte("SAM")
	type direction struct {
		dx, dy int
	}
	for x := range in {
		for y := range in[x] {
			switch {
			case startsWith(dir(x, y, 1, 1, in), mas):
			case startsWith(dir(x, y, 1, 1, in), sam):
			default:
				continue
			}
			// :()
			if x+2 < len(in) && startsWith(dir(x+2, y, -1, 1, in), mas) {
				total++
				continue
			}
			if x+2 < len(in) && startsWith(dir(x+2, y, -1, 1, in), sam) {
				total++
				continue
			}
			if y+2 < len(in) && startsWith(dir(x, y+2, 1, -1, in), mas) {
				total++
				continue
			}
			if y+2 < len(in) && startsWith(dir(x, y+2, 1, -1, in), sam) {
				total++
				continue
			}
		}
	}
	fmt.Printf("total X-MAS: %d\n", total)
}

func startsWith(bs iter.Seq[byte], want []byte) bool {
	i := 0
	for a, b := range it.Zip(bs, slices.Values(want)) {
		i++
		if a != b {
			return false
		}
	}
	return i == len(want)
}

func dir(x, y, dx, dy int, in [][]byte) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for x >= 0 && x < len(in) && y >= 0 && y < len(in[x]) {
			if !yield(in[x][y]) {
				return
			}
			x += dx
			y += dy
		}
	}
}

func directions(in [][]byte, x, y int) iter.Seq[iter.Seq[byte]] {
	return func(yield func(iter.Seq[byte]) bool) {
		type direction struct {
			dx, dy int
		}
		for _, d := range []direction{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		} {
			if !yield(dir(x, y, d.dx, d.dy, in)) {
				return
			}
		}
	}
}

func read() ([][]byte, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return bytes.Split(b, []byte{'\n'}), nil
}
