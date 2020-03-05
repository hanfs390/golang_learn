package main

import "fmt"
func judge(a int) (flag int) {
	if a > 5 {
		fmt.Println("a > 5")
		return 1;
	} else {
		fmt.Println("a <= 5")
		return 0;	
	}
}
func main() {
	var a = 5 
	judge(a)
}
