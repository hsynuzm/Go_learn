package main

import "fmt"

func main() {

	//Fixed
	/*first method
	var names [3]string

	names[0] = "jhon"
	names[1] = "james"
	names[2] = "cena"

	fmt.Println(names)*/

	var names = [3]string{"jhon", "james", "cena"}
	names[1] = "jane"
	fmt.Println(names)
}
