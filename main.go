package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fmt.Println(data[os.Args[1]])
		// } else {
		// for k, v := range data {
		// 	fmt.Println(k, ":", v)
		// }
	}
}
