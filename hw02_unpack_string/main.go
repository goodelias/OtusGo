package main

import "fmt"

func main() {
	result, err := Unpack("f33")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
