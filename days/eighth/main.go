// eighth is the eighth day.
package main

import (
	"bytes"
	"fmt"
	"io"
	"iter"
	"log"
	"os"
)

func main() {
	aerials, bounds, err := read()
	if err != nil {
		log.Fatal(err)
	}
	aerialsByFreq := make(map[byte][]aerial)
	for _, a := range aerials {
		aerialsByFreq[a.freq] = append(aerialsByFreq[a.freq], a)
	}

	fmt.Println("part one:")
	partOne(aerialsByFreq, bounds)
	fmt.Println("part two:")
	partTwo(aerialsByFreq, bounds)
}

func partOne(aerialsByFreq map[byte][]aerial, bounds point) {
	inBounds := func(p point) bool {
		if p[0] < 0 || p[1] < 0 {
			return false
		}
		if p[0] >= bounds[0] || p[1] >= bounds[1] {
			return false
		}
		return true
	}
	antinodes := make(map[point]bool)
	for _, as := range aerialsByFreq {
		for i, a := range as[:len(as)-1] {
			for _, b := range as[i+1:] {
				for _, an := range a.antinodes(b) {
					if !inBounds(an) {
						continue
					}
					antinodes[an] = true
				}
			}
		}
	}
	fmt.Printf("unique antinodes: %d\n", len(antinodes))
}

func partTwo(aerialsByFreq map[byte][]aerial, bounds point) {
	antinodes := make(map[point]bool)
	out := make([][]byte, int(bounds[1]))
	for i := range out {
		out[i] = make([]byte, int(bounds[0]))
		for j := range out[i] {
			out[i][j] = '.'
		}
	}
	for f, as := range aerialsByFreq {
		fmt.Printf("%c:\n", f)
		for _, a := range as {
			out[int(a.loc[1])][int(a.loc[0])] = f
		}
		for i, a := range as[:len(as)-1] {
			for _, b := range as[i+1:] {
				for an := range a.antinodes2(b, bounds) {
					fmt.Printf("\t%v\n", an)
					if an != a.loc && an != b.loc {
						out[int(an[1])][int(an[0])] = '#'
					}
					antinodes[an] = true
				}
			}
		}
	}
	fmt.Printf("unique antinodes: %d\n", len(antinodes))
	fmt.Printf("%s\n", bytes.Join(out, []byte("\n")))
}

type point [2]float64

func (p point) sub(q point) point {
	return point{p[0] - q[0], p[1] - q[1]}
}

func (p point) add(q point) point {
	return point{p[0] + q[0], p[1] + q[1]}
}

type aerial struct {
	freq byte
	loc  point
}

func (a aerial) antinodes(b aerial) [2]point {
	d := a.loc.sub(b.loc)
	return [2]point{
		a.loc.add(d),
		b.loc.sub(d),
	}
}

func (a aerial) antinodes2(b aerial, bounds point) iter.Seq[point] {
	inBounds := func(p point) bool {
		if p[0] < 0 || p[1] < 0 {
			return false
		}
		if p[0] >= bounds[0] || p[1] >= bounds[1] {
			return false
		}
		return true
	}
	return func(yield func(point) bool) {
		d := point{a.loc[0] - b.loc[0], a.loc[1] - b.loc[1]}
		for p := a.loc; inBounds(p); p = p.add(d) {
			if !yield(p) {
				return
			}
		}
		for p := b.loc; inBounds(p); p = p.sub(d) {
			if !yield(p) {
				return
			}
		}

	}
}

func read() ([]aerial, point, error) {
	all, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, point{}, err
	}
	// todo also compute the bounds
	var (
		aerials []aerial
		lines   = bytes.Split(all, []byte("\n"))
	)
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	for y, l := range lines {
		for x, c := range l {
			if c != '.' {
				aerials = append(aerials, aerial{
					freq: c,
					loc:  point{float64(x), float64(y)},
				})
			}
		}
	}
	return aerials, point{float64(len(lines[0])), float64(len(lines))}, nil
}
