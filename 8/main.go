package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	children []node
	metaData []int
}

func (n *node) checksum() int {
	return n.sumChildren() + n.sumMetaData()
}

func (n *node) sumChildren() int {
	total := 0

	for _, c := range n.children {
		total += c.checksum()
	}

	return total
}

func (n *node) sumMetaData() int {
	total := 0
	for _, m := range n.metaData {
		total += m
	}
	return total
}

func (n *node) value() int {
	if len(n.children) == 0 {
		return n.sumMetaData()
	}

	total := 0

	for _, c := range n.metaData {
		if c-1 < len(n.children) {
			total += n.children[c-1].value()
		}
	}

	return total
}

type LicenceAnalysis struct {
	root     node
	Checksum int
	Value    int
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	analysis := Analyse(f)
	fmt.Printf("%d\n", analysis.Checksum)
	fmt.Printf("%d\n", analysis.Value)
}

func Analyse(r io.Reader) *LicenceAnalysis {
	res := &LicenceAnalysis{root: parseInput(r)}

	res.Checksum = res.root.checksum()
	res.Value = res.root.value()

	return res
}

func parseInput(r io.Reader) node {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	return scanNode(scanner, 0)
}

func scanNode(scanner *bufio.Scanner, depth int) node {
	var nChildren, nMetaData int

	// scan n children
	readInt(scanner, &nChildren)
	// scan n meta data
	readInt(scanner, &nMetaData)

	// fmt.Printf("Scanning node at depth %d with %d children and %d metata\n", depth, nChildren, nMetaData)

	res := node{}

	// scan each child node
	for i := 0; i < nChildren; i++ {
		res.children = append(res.children, scanNode(scanner, depth+1))
	}

	// scan each meta data entry
	for i := 0; i < nMetaData; i++ {
		res.metaData = append(res.metaData, scanMetaData(scanner))
	}

	return res
}

func scanMetaData(scanner *bufio.Scanner) int {
	var i int
	readInt(scanner, &i)
	return i
}

func readInt(scanner *bufio.Scanner, i *int) {
	scanner.Scan()
	_, err := fmt.Sscanf(scanner.Text(), "%d", i)

	if err != nil {
		panic(err)
	}
}
