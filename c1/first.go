package main

import (
	"fmt"
	"os"
	"strings"
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
	for index,arg := range os.Args {
		s += temp + arg
		temp = " "
		fmt.Printf("index is %d and data is %s \n",index,arg)
	}
	fmt.Printf("s2 is %s \n",s)

	//使用join来拼接字符串
	fmt.Println(strings.Join(os.Args," \n"))
}
