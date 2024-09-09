package main

import "fmt"

func main() {

	i := 5
	f := 5.0
	s := "Name"
	b := true
	print(i, f, s, b)

	//array
	names := [...]string{"Rajib", "Kohinur", "Tasfiya", "Tahiya", "Wazia"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
