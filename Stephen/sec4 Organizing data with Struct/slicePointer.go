package main

import "fmt"

func slicePointer(){
	mySlice := [] string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice (s []string){
	s[0] = "Bye"
}