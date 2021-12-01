package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"flag"
)

func main() {
	var count int
	flag.IntVar(&count, "n", 5, "# of lines to read")
	flag.Parse()

	var in io.Reader
	in = os.Stdin

	buf := bufio.NewScanner(in)

	for i := 0; i<count; i++ {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())

	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: ", err)
	}
}
