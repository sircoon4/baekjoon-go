package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	dogString := `|\_/|
|q p|   /}
( 0 )"""\`
	dogString = dogString + "\n|\"^\"`    |\n"
	dogString = dogString + `||_/=\\__|`

	fmt.Fprint(writer, dogString)
}
