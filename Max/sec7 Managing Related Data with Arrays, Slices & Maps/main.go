package main

import "fmt"

type FloatMap map[string]float64

func (m FloatMap) output(){
fmt.Println(m)
}

func main(){
  userNames := make([] string, 2,5)
  // userNames := [] string{}

  userNames[0] = "Julie"
  userNames = append(userNames, "Max")
  userNames = append(userNames, "Manual")

  fmt.Println(userNames)

  courseRatings := make(FloatMap, 3)
  // courseRatings := map[string]float64{}

  courseRatings["go"] = 4.7
  courseRatings["react"] = 4.8
  courseRatings["angular"] = 4.7

  // fmt.Println(courseRatings)

  courseRatings.output()

  for index, value := range userNames{
    // ..
    fmt.Println("Index: ",index)
    fmt.Println("Value: ",value)
  }

  for key, value := range courseRatings{
        // ..
    fmt.Println("Key: ",key)
    fmt.Println("Value: ",value)
  }

}