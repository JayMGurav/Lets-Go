package main

import "fmt"

func main(){

	if 7%2 == 0 {
		fmt.Println("7 is even");
	}else {
		fmt.Println("7 is odd");
	}


	if j := 8; j < 0 {
		fmt.Println("J is negative number");
	} else if j == 0 {
		fmt.Println("j is zero!");
	} else {
		fmt.Println("j is positive numnber");
	}
}
