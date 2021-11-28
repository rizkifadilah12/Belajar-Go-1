package main

import (
	"fmt"
)

func main() {
	type NoKtp string
	type Married bool

	var noKtpBudi NoKtp = "12121212"
	var MarriedStatus Married = true
	fmt.Println(noKtpBudi)
	fmt.Println(MarriedStatus)
}
