package main

import "fmt"

func main(){
	var s []string
	fmt.Println("un-init: ", s, len(s), cap(s))

	s = make([]string, 3);
	fmt.Println("emp: ", s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("initiate: ", s)
	
	s = append(s, "d");
	s = append(s, "e");
	fmt.Println("appended: ", s)
	
	c := make([]string, len(s));
	copy(c, s);
	fmt.Println("copy: ", c);

	l := s[3:]
	fmt.Println("sliced: ", l);

	t := []string{"1", "2", "3"}
	fmt.Println("initialised: ", t);

	twoD := make([][]int, 3);
	for i:=0; i < 3; i++ {
		innerLen := i + 1;
		twoD[i] = make([]int, innerLen);
		for j:=0; j < innerLen; j++ {
			twoD[i][j] = i + j;
		}
	}

	fmt.Println("2D: ", twoD);

}
