package main

import "fmt"

type Product struct{
  title string
  id string 
  price float64
}

func main(){
prices := []float64{10.99, 8.99}
fmt.Println(prices[0:1])
prices[1] = 9.99

updatedPrices := append(prices, 5.99)
fmt.Println(updatedPrices, prices)


discountPrices := []float64{101.99,80.99,20.59}
prices = append(prices, discountPrices...)
fmt.Println(prices)
}

  // var productNames = [4]string{"A Book"}
  // prices := [4]float64{10.99, 9.99, 45.99,20.0}
  // fmt.Println(prices)

  // productNames[2] = "A Carpet"

  // fmt.Println(productNames)

  // fmt.Println(prices[0])

  // // featuredPrice := prices[:3]
  // // featuredPrice := prices[1:]
  // featuredPrice := prices[1:3]
  // featuredPrice[0] = 199.99
  // highlightedPrice := featuredPrice[:1]
  // fmt.Println(featuredPrice)
  // fmt.Println(highlightedPrice)
  // fmt.Println(prices)

  // fmt.Println(len(featuredPrice), cap(featuredPrice))
  // // fmt.Println(&featuredPrice[0], &prices[1])

  // highlightedPrice = highlightedPrice[:3]
  //   fmt.Println(highlightedPrice)
  // fmt.Println(len(featuredPrice), cap(featuredPrice))