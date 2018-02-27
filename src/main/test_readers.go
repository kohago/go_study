package main

import (
	"fmt"
	"io"
	"strings"
)

func testReaders() {
	r := strings.NewReader("Hello! new Reader")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v", n, err, b)
		fmt.Println("b[:n]=%q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
