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
	fmt.Printf("s1 is %s \n",s)
	s = ""
	//迭代方式2
	for _,arg := range os.Args {
		s += temp + arg
		temp = " "
	}
	fmt.Printf("s2 is %s \n",s)
}
