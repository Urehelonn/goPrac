package main

import (
	"fmt"
)

func main() {
	fvs := Favs()
	for fv := range fvs {
		fmt.Println(fvs[fv])
	}
}

func Favs() []string {
	return []string{"Apple", "Banana", "Orange", "Pineapple"}
}
