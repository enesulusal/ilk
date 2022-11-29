package main

import "fmt"

func main() {

	var iller map[string]string
	iller = make(map[string]string)
	iller["06"] = "ankara"
	fmt.Println(iller)
	iller["06"] = "yeni il"
	fmt.Println(iller)

	il06 := iller["06"]
	fmt.Println(il06)

	il01 := iller["01"]
	fmt.Println(il01)
	delete(iller, "06")
	il06, ok := iller["06"]

	if ok {

		fmt.Println(il06)
	} else {

		fmt.Println("06 bulunamadı")
	}
}
