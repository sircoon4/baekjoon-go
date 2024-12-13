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
	fmt.Fscanln(reader, &a, &b, &c)

	var max int
	var same int
	var dice [6]int

	same = a
	dice[a-1] += 1
	max = a

	if dice[b-1] > 0 {
		same = b
	}
	dice[b-1] += 1
	if b > max {
		max = b
	}

	if dice[c-1] > 0 {
		same = c
	}
	dice[c-1] += 1
	if c > max {
		max = c
	}

	if dice[same-1] == 3 {
		fmt.Fprint(writer, 10000+same*1000)
	} else if dice[same-1] == 2 {
		fmt.Fprint(writer, 1000+same*100)
	} else {
		fmt.Fprint(writer, max*100)
	}
}
