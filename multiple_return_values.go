package main

import "fmt";

func multi_return() (int, int) {
	return 3, 7;
}

func main(){
	var a, b int = multi_return();
	fmt.Println(a, b);
	
	_, c := multi_return();
	fmt.Println(c)

}
