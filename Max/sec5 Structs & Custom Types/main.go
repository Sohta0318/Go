package main

import "fmt"

type str string

func (text str) log(){
  fmt.Println(text)
}

func test(){
var name str
name = "Test"
name.log()
}