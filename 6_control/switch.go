package main

import "fmt"
func switch_int(i int) {
    switch i {
		case 0:
			fmt.Printf("0")
		case 1:
			fmt.Printf("1")
		case 2:
			fmt.Printf("2 fall")
			fallthrough
		case 3:
			fmt.Printf("3")
		case 4, 5, 6:
			fmt.Printf("|4, 5, 6|")
		default:
			fmt.Printf("Default")
	}

}
func main() {
	switch_int(0)
	switch_int(1)
	switch_int(2)
	switch_int(3)
	switch_int(4)
	switch_int(5)
	switch_int(6)
	switch_int(7)
	switch_int(8)
}
