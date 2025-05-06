package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
// var revenue float64
// var expense float64
// var taxRate float64

// fmt.Print("Revenue: ")
// fmt.Scan(&revenue)

// fmt.Print("Expense: ")
// fmt.Scan(&expense)

// fmt.Print("Tax Rate: ")
// fmt.Scan(&taxRate)

revenue, err := getUserInfo("Revenue: ")

if err != nil{
  fmt.Println(err)
  return
}

expense, err := getUserInfo("Expense: ")

if err != nil{
  fmt.Println(err)
  return
}
taxRate, err:= getUserInfo("Tax Rate: ")

if err != nil{
  fmt.Println(err)
  return
}
// if err1 != nil || err2 != nil || err3 != nil{
//   fmt.Println(err1)
//   return
// }

// ebt := revenue -expense
// profit := ebt * (1 - taxRate/100)
// ration := ebt/ profit

ebt, profit, ration := calculateFinancial(revenue, expense, taxRate)

fmt.Printf("%.1f\n",ebt)
fmt.Printf("%.1f\n",profit)
fmt.Printf("%.3f\n",ration)

storeResultToFile(ebt, profit, ration)
}

func calculateFinancial(revenue, expense, taxRate float64)(float64, float64, float64){
  ebt := revenue -expense
  profit := ebt * (1 - taxRate/100)
  ration := ebt/ profit
  return ebt, profit, ration
}

func getUserInfo(infoText string) (float64, error){
var userInput float64
fmt.Print(infoText)
fmt.Scan(&userInput)
if userInput <= 0{
  return 0, errors.New("Value Must be greater than 0")
}
return userInput, nil
}

func storeResultToFile(ebt, profit, ration float64){
  results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRation: %.1f\n",ebt, profit, ration)
os.WriteFile("calculated.txt", []byte(results), 0644)
}