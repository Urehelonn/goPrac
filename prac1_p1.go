package main

import (
	"fmt"

	q2 "rsc.io/quote/v2"
	q3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(q2.OptV2())
	fmt.Println(q3.GlassV3())
}
