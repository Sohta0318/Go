package main

import "fmt"

func main(){
  websites := map[string]string{
    "Google": "https://google.com",
    "Amazon Web Service": "https://aws.com",
  }
  fmt.Println(websites)
  fmt.Println(websites["Amazon Web Service"])

  websites["LinkedIn"] = "https://linkedin.com"
  fmt.Println(websites)

  delete(websites, "Google")
  fmt.Println(websites)
}