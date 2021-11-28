package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Budi":
		fmt.Println("Hallo Budi")
	case "Joko":
		fmt.Println("Hallo Joko")
	default:
		fmt.Println("kamu siapa?")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama Terlalu Panjang")
	//case false:
	//	fmt.Println("Nama Sudah Benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
