package main

import "fmt"


func main(){
num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// addNum(num, 10)
iterate(num)
}

// func addNum (num []int, n int){
// 	for i :=0; i <= n; i++{
// 		num = append(num, i)
// 	}
// 	fmt.Println(num)
// }

func iterate (num []int){
for _,value:= range num{
if value % 2 ==0{
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
}
}