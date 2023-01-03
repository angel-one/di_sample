package main

import (
	beans "di_sample/beans"
	"log"
)

func main() {
	beans.Initialise() // Note : not to be called in Init()!!
	log.Println("hello world")
	log.Println("beanB", theBeanB.getVal())
}
