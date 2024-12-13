package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	catString := `\    /\
 )  ( ')
(  /  )
 \(__)|`
	fmt.Fprint(writer, catString)
}
