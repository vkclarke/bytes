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
	done := make(chan int)

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

	go func() {
		for b := range pipe {
			fmt.Fprintf(out, "%x\n", b)
		}

		done <- 0
	}()

	os.Exit(<-done)
}
