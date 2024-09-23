package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
)

func bytes(in io.ByteReader) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for {
			if b, err := in.ReadByte(); err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			} else if !yield(b) {
				break
			}
		}
	}
}

func main() {
	for b := range bytes(bufio.NewReader(os.Stdin)) {
		fmt.Println(b)
	}
}
