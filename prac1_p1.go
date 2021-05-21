package main

import (
	"fmt"
)

func main() {
	fvs := favs()
	for fv := range fvs {
		fmt.Println(fv)
	}
}

func favs() []string {
	return []string{"Apple", "Banana", "Orange", "Pineapple"}
}
