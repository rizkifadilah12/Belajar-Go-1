package main

import (
	"fmt"
)

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	//var lulusNilai bool = nilaiAkhir > 80
	//var lulusAbsensi bool = absensi > 80

	//var lulus bool = lulusNilai && lulusAbsensi

	fmt.Println(nilaiAkhir > 80 && absensi > 80)
}
