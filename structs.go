package main

import "fmt"

type person struct {
	name string
	age int
}

func new_person(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p;
}


func main(){
	fmt.Println(person{"Bob", 20});

	fmt.Println(person{name: "Jay", age: 24});

	fmt.Println(person{name: "Joe"})

	s := person{name: "hwllo"}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

}
