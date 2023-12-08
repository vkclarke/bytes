package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	pipe := make(chan byte, os.Getpagesize())

	go func() {
		in := bufio.NewReader(os.Stdin)

		for b, eof := in.ReadByte(); eof == nil; b, eof = in.ReadByte() {
			pipe <- b
		}

		close(pipe)
	}()

	for b := range pipe {
		fmt.Println(b)
	}
}
