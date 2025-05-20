package main

import "fmt"

func main() {
  // 1
hobbies := [3]string{"Travel", "Work out", "Movie"}
fmt.Println(hobbies)

// 2
fmt.Println(hobbies[0])
fmt.Println(hobbies[1:3])

// 3
newHobbies := hobbies[0:2]
newHobbies2 := hobbies[:2]
fmt.Println(newHobbies2)


// 4
reNewHobbies := newHobbies[1:3]
fmt.Println(reNewHobbies)

// 5
courseGoals := []string{"Go complete", "Master Go"}

// 6
courseGoals[1] = "Create project"
newCourseGoals := append(courseGoals, "New Goal")
fmt.Println(newCourseGoals)

// 7
type Product struct{
  title string
  id string
  price float64
}
products := []Product{
  Product{"first title", "1", 32.22},
  Product{"second title", "2", 40.1},
}
newProducts := append(products, Product{"third title", "3", 50.0})
fmt.Println(newProducts)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.