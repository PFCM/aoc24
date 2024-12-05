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

func directions(in [][]byte, x, y int) iter.Seq[iter.Seq[byte]] {
	return func(yield func(iter.Seq[byte]) bool) {
		dir := func(dx, dy int) iter.Seq[byte] {
			x, y := x, y
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
			if !yield(dir(d.dx, d.dy)) {
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
