package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, e1 := strconv.ParseUint("789", 0, 64)
	fmt.Println(u, e1)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e2 := strconv.Atoi("wat")
	fmt.Println(e2)
}
