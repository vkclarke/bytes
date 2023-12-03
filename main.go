package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := os.Stdout

	pipe := make(chan byte, 32)

	go func() {
		for {
			b, err := in.ReadByte()
			if err != nil {
				break
			}
			pipe <- b
		}

		close(pipe)
	}()

	for b := range pipe {
		fmt.Fprintln(out, b)
	}
}
