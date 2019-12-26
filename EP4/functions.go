package main

import "fmt"

func imTheMan() string {

	fmt.Println("I'm the man")
	return "I'm the man"

}

func imTheManName(name string) string {

	fmt.Println(name + " is the man")
	return name + " is the man"

}

func main() {

	imTheMan()
	manName := imTheManName("Chase")
	fmt.Println(manName)

}
