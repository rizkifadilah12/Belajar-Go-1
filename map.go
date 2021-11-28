package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "budi",
		"address": "gatau",
	}
	person["title"] = "programer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	//menghapus map / data dalam map
	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "budi"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
