package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var nl, nw, nc = 0, 0, 0
	var isInside = false

	// Read the input
	for {
		c, err := reader.ReadByte()
		if err != nil {
			break
		}

		nc++

		if c == '\n' {
			nl++
		}

		if c == ' ' || c == '\n' || c == '\t' {
			isInside = false
			continue
		}

		if !isInside {
			isInside = true
			nw++
		}
	}

	fmt.Printf("%d %d %d\n", nl, nw, nc)
}
