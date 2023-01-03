package main

import (
	"di_sample/beans"
)

type BeanA struct {
	val int
}

var theBeanA BeanA

func init() {
	theBeanA = BeanA{}
	beans.Register(&theBeanA, nil)
}

func (this BeanA) getVal() int {
	return this.val
}

func (this BeanA) setVal(value int) {
	this.val = value
}

func (this BeanA) GetName() string {
	return "A"
}

func (this *BeanA) Initialize(deps map[string]interface{}) {
	this.val = 47 // dummy!
}
