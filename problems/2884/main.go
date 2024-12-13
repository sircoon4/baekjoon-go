package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	if b >= 45 {
		fmt.Fprint(writer, a, b-45)
	} else if a == 0 {
		fmt.Fprint(writer, 23, b+15)
	} else {
		fmt.Fprint(writer, a-1, b+15)
	}
}
