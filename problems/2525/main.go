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

	var a, b, c int
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &c)

	afterCookHours := a + (c / 60)
	afterCookMinutes := b + c%60

	if afterCookMinutes >= 60 {
		afterCookMinutes -= 60
		afterCookHours++
	}

	if afterCookHours >= 24 {
		afterCookHours %= 24
	}

	fmt.Fprint(writer, afterCookHours, " ", afterCookMinutes)
}
