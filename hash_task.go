package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {

	var str string = "hello world"

	for i := 0; i < 100000; i++ {
		new_str := str + strconv.Itoa(i)

		h := sha256.Sum256([]byte(new_str))

		// convert result to string
		res := fmt.Sprintf("%x", h)

		if res[0:4] == "0000" {
			fmt.Println(res)
			fmt.Println(i)
			break
		}

	}
}
