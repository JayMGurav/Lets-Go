package main

import "fmt"

func zeroVal(val int){
	val = 0;
}

func zeroValPtr(valPtr *int){
	*valPtr = 0;
}

func main(){
	
	val := 1;
	zeroVal(val);
	fmt.Println("val: ", val);

	zeroValPtr(&val);
	fmt.Println("changed: ", val);
	fmt.Println("val address: ", &val);

}
