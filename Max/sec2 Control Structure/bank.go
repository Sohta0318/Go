package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
  data, err := os.ReadFile(accountBalanceFile)

  if err != nil{
    return 1000, errors.New("Fail to find balance file.")
  }

  balanceText := string(data)
  balance, err :=strconv.ParseFloat(balanceText, 64)

  if err != nil{
    return 1000, errors.New("Failed to parse stored balance value.")
  }

  return balance, nil
}
func writeBalanceToFile(balance float64){
  balanceText := fmt.Sprint(balance)
os.WriteFile(accountBalanceFile, []byte(balanceText),0644)
}

func main(){
  var accountBalance, err  = getBalanceFromFile()

  if err != nil{
    fmt.Println("ERROR")
    fmt.Println(err)
    fmt.Println("----------")
    // panic("Can't continue, Sorry.")
  }

  fmt.Println("Welcome to Go bank!!")

  for {
  fmt.Println("What do you want to do?")
  fmt.Println("1. Check the balance")
  fmt.Println("2. Deposit money")
  fmt.Println("3. Withdraw money")
  fmt.Println("4. Exit")

  var choice int
  fmt.Print("Your choice: ")
  fmt.Scan(&choice)

  switch(choice){
  case 1: 
  fmt.Println("Your Balance is ", accountBalance)
  case 2: 
  fmt.Print("Your deposit: ")
  var depositAmount float64
  fmt.Scan(&depositAmount)

if depositAmount <= 0{
fmt.Println("Invalid amount. Must be greater then 0.")
// return
continue
}

  accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
  writeBalanceToFile(accountBalance)
  fmt.Println("Balance updated!! New amount:", accountBalance)
case 3:
  fmt.Print("Withdrawal amount: ")
  var withdrawalAmount float64
  fmt.Scan(&withdrawalAmount)

if withdrawalAmount <= 0{
fmt.Println("Invalid amount. Must be greater then 0.")
return
}

if withdrawalAmount > accountBalance{
fmt.Println("Invalid amount. You can't withdraw more then you have.")
return
}

  accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
  writeBalanceToFile(accountBalance)

  fmt.Println("Balance updated!! New amount:", accountBalance)
default: 
fmt.Println("Goodbye!")
fmt.Println("Thank you for choosing our bank!")

    return 
    // break
  }

//   // wantsCheckBalance:= choice ==1
//   if choice ==1 {
//     fmt.Println("Your Balance is ", accountBalance)
//   }else if choice == 2{
//     fmt.Print("Your deposit: ")
//     var depositAmount float64
//     fmt.Scan(&depositAmount)

// if depositAmount <= 0{
//   fmt.Println("Invalid amount. Must be greater then 0.")
//   // return
//   continue
// }

//     accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
//     fmt.Println("Balance updated!! New amount:", accountBalance)
//   }else if choice == 3{
//     fmt.Print("Withdrawal amount: ")
//     var withdrawalAmount float64
//     fmt.Scan(&withdrawalAmount)

// if withdrawalAmount <= 0{
//   fmt.Println("Invalid amount. Must be greater then 0.")
//   return
// }

// if withdrawalAmount > accountBalance{
//   fmt.Println("Invalid amount. You can't withdraw more then you have.")
//   return
// }

//     accountBalance -= withdrawalAmount // accountBalance = accountBalance - withdrawalAmount
//     fmt.Println("Balance updated!! New amount:", accountBalance)
//   }else{
//     fmt.Println("Goodbye!")
//     // return 
//     break
//   }
}

// fmt.Println("Thank you for choosing our bank!")
}