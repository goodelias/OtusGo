package main

import "fmt"

func main() {
	result, err := Unpack("")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
