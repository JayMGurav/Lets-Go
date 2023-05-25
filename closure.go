package main

import "fmt";


func countSeq() func() int {
	count:=0;
	return func() int {
		count++;
		return count;
	}
}

func main(){
	initCount := countSeq();

	fmt.Println(initCount());
	fmt.Println(initCount());
	fmt.Println(initCount());
	fmt.Println(initCount());
}
