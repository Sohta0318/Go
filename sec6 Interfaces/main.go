package main

import "fmt"

type englishBot struct {}
type spanishBot struct {}

func main(){
eb := englishBot{}
sb := spanishBot{}

printGreeting(eb)
// printGreeting(sb)
}

func printGreeting(eb englishBot){
fmt.Println(eb.getGreeting())
}

// func printGreeting(sb englishBot){
// fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string{
	// Very custom logic for generating an english greeting
return "Hi There!"
}

func (spanishBot) getGreeting() string{
		// Very custom logic for generating a spanish greeting
		return "Hola!"
}