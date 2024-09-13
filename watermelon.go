package main

import (
	"fmt"
	"os"
)

func main() {
	var i int
	fmt.Fscanf(os.Stdin, "%d", &i)
	if i %2 == 0 && i > 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
