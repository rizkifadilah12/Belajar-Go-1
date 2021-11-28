package main

import "fmt"

func main() {
	name := "budi"

	if name == "budi" {
		fmt.Println("hallo budi")
	} else if name == "joko" {
		fmt.Println("hallo joko")
	} else {
		fmt.Println("kamu siapa?")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
