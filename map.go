package main

import "fmt"

func main(){

	m := make(map[string]int);
	fmt.Println("map: ", m);

	m["k1"] = 1;
	m["k2"] = 2;
	fmt.Println("intilised map: ", m);

	v1 := m["k1"];
	fmt.Println("value of k1: ", v1);

	//length
	fmt.Println("length of map: ", len(m));

	delete(m, "k1");
	fmt.Println("deleted k1: ", m);

	_, prs := m["k2"]
	fmt.Println("prs: ", prs);

	//initialised map
	i := map[string]int{"foo": 1, "boo": 2};
	fmt.Println("initialised map: ", i);

}
