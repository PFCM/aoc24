// fifth is day five.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var topoPartOneFlag = flag.Bool("topo", false, "whether to use the topological sort approach for part one")

func main() {
	flag.Parse()

	rules, updates, err := read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("part one:")
	if *topoPartOneFlag {
		if err := partOneTopo(rules, updates); err != nil {
			log.Fatal(err)
		}
	} else {
		partOne(rules, updates)
	}

	fmt.Println("part two:")
	if err := partTwo(rules, updates); err != nil {
		log.Fatal(err)
	}
}

func partOne(rules map[int]map[int]bool, updates [][]int) {
	total := 0
	for _, u := range updates {
		if inOrder(rules, u) {
			total += u[len(u)/2]
		}
	}
	fmt.Printf("sum of middle of in order updates: %d\n", total)
}

func partOneTopo(rules map[int]map[int]bool, updates [][]int) error {
	total := 0
	for _, u := range updates {
		if inOrderTopo(rules, u) {
			total += u[len(u)/2]
		}
	}
	fmt.Printf("sum of middle of in order updates: %d\n", total)
	return nil
}

func partTwo(rules map[int]map[int]bool, updates [][]int) error {
	total := 0
	for _, u := range updates {
		if inOrder(rules, u) {
			continue
		}
		// re-order
		ordered, err := topoSubgraph(rules, u)
		if err != nil {
			return err
		}
		total += ordered[len(ordered)/2]
	}
	fmt.Printf("sum of middle of re-ordered updates: %d\n", total)
	return nil
}

func topoSubgraph(edges map[int]map[int]bool, nodeList []int) ([]int, error) {
	// it's just a dfs with some odd bookkeeping.
	var (
		visit func(n int) error
		temp  = make(map[int]bool)
		seen  = make(map[int]bool)
		nodes = make(map[int]bool)
		out   []int
	)
	visit = func(n int) error {
		if seen[n] {
			return nil
		}
		if !nodes[n] {
			// Not in the relevant subgraph.
			return nil
		}
		if temp[n] {
			return fmt.Errorf("cycle detected involving %d", n)
		}
		temp[n] = true
		for m := range edges[n] {
			if err := visit(m); err != nil {
				return err
			}
		}
		seen[n] = true
		out = append(out, n)
		return nil
	}
	for _, n := range nodeList {
		nodes[n] = true
	}
	for n := range nodes {
		if seen[n] {
			continue
		}
		if err := visit(n); err != nil {
			return nil, err
		}
	}
	slices.Reverse(out)
	return out, nil
}

func inOrderTopo(rules map[int]map[int]bool, update []int) bool {
	ordering, err := topoSubgraph(rules, update)
	if err != nil {
		panic(err)
	}
	for len(update) > 0 && len(ordering) > 0 {
		switch u, o := update[0], ordering[0]; {
		case u == o:
			update = update[1:]
			ordering = ordering[1:]
		default:
			ordering = ordering[1:]
		}
	}
	if len(ordering) == 0 && len(update) != 0 {
		return false
	}
	return true
}

func inOrder(rules map[int]map[int]bool, update []int) bool {
	// n^2 :/
	for i, a := range update {
		for _, b := range update[:i] {
			if rules[a][b] {
				return false
			}
		}
	}
	return true
}

func read() (map[int]map[int]bool, [][]int, error) {
	scan := bufio.NewScanner(os.Stdin)
	rules := make(map[int]map[int]bool)
	for scan.Scan() {
		l := scan.Text()
		if l == "" {
			break
		}
		a, b, ok := strings.Cut(l, "|")
		if !ok {
			return nil, nil, fmt.Errorf("invalid rule: %q", l)
		}
		x, err := strconv.Atoi(a)
		if err != nil {
			return nil, nil, err
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			return nil, nil, err
		}
		if _, ok := rules[x]; !ok {
			rules[x] = make(map[int]bool)
		}
		rules[x][y] = true
	}
	if err := scan.Err(); err != nil {
		return nil, nil, err
	}

	updates := [][]int{nil}
	for scan.Scan() {
		raw := strings.Split(scan.Text(), ",")
		for _, r := range raw {
			x, err := strconv.Atoi(r)
			if err != nil {
				return nil, nil, err
			}
			updates[len(updates)-1] = append(updates[len(updates)-1], x)
		}
		updates = append(updates, nil)
	}
	if err := scan.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates[:len(updates)-1], nil
}
