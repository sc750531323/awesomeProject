package main

import (
	"fmt"
	"os"
)

func main() {
	var s,temp string
	for i := 0; i < len(os.Args); i++ {
		s += temp + os.Args[i]
		temp = " "
	}
	fmt.Printf("s is %s \n",s)
}
