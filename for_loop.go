package main

import "fmt"

func main(){
	
	i := 3;

	for i <= 6 {
		fmt.Println(i);
		i = i + 1
	}

	for j:=7; j > 0; j-- {
		fmt.Println(j);
	}

	//infinite loop with break

	for {
		fmt.Println("for loop");
		break;
	}

	//print all odd numbers
	for n := 0; n <= 10; n++ {
		if n % 2 == 0{
			continue;
		}
		fmt.Println(n);
	}

}
