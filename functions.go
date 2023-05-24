package main

import "fmt";


func add(a int, b int) int {
	return a + b;
}

func multiply(a, b int) int {
	return a * b;
}

func main(){
	sum := add(2, 3);
	fmt.Println("Sum: ", sum);

	product := multiply(2, 3);
	fmt.Println("Product: ", product);
}
