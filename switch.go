package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2
	fmt.Println("Write ", i, "as " )

	switch i {
		case 1: 
			fmt.Println("one")
		case 2:
			fmt.Println("two");
		case 3: 
			fmt.Println("three");
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("Its weekend");
		default:
			fmt.Println("Its week day");
	}

	t := time.Now();
	switch {
		case t.Hour() < 12:
			fmt.Println("Its before noon :)")
		default:
			fmt.Println("its after noon");
	}

	what_i_am := func(i interface{}) {
		switch 	t := i.(type) {
			case bool:
				fmt.Println("I'm Boolean");
			case int:
				fmt.Println("I'm Int");
			default:
				fmt.Println("dunno the type %T", t);				
		}
	}

	what_i_am(true)
	what_i_am(1)
	what_i_am("hey")

}
