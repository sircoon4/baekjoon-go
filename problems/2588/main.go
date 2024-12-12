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
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	s1 := b / 10
	r1 := b % 10
	s2 := s1 / 10
	r2 := s1 % 10
	r3 := s2

	fmt.Fprintln(writer, a*r1)
	fmt.Fprintln(writer, a*r2)
	fmt.Fprintln(writer, a*r3)
	fmt.Fprint(writer, a*b)
}
