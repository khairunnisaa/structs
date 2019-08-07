package main

import "fmt"

type person struct {
	lastName string
	firstName string
}

func main () {
	nisa := person{"nisa","cantik"}
	fmt.Println(nisa)
}
