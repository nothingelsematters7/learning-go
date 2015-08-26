package main

import (
	"fmt"
	"github.com/nothingelsematters7/stringutil"
	"os"
	"strings"
)

func main() {
	who := "World!"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}

	greeting := "Hello " + who

	fmt.Println(stringutil.Reverse(greeting))

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
