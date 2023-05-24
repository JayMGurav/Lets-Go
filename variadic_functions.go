package main

import "fmt";

func sum(nums ...int) int {
	total := 0;
	for _, num := range nums {
		total += num;
	}
	return total;
}

func main(){
		
	sum_a := sum(1,2);
	fmt.Println("sum of 1,2 is: ", sum_a);

	sum_b := sum(1,2, 3);
	fmt.Println("sum of 1,2 ,3 is: ", sum_b);
	
	nums := []int{1, 2, 3, 4, 5};
	var array_sum = sum(nums...);
	fmt.Println("array sum: ", array_sum)
}
