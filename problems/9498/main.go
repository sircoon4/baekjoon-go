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

	var a int
	fmt.Fscanln(reader, &a)

	if a >= 90 {
		fmt.Fprint(writer, "A")
	} else if a >= 80 {
		fmt.Fprint(writer, "B")
	} else if a >= 70 {
		fmt.Fprint(writer, "C")
	} else if a >= 60 {
		fmt.Fprint(writer, "D")
	} else {
		fmt.Fprint(writer, "F")
	}
}
