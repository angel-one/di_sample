package main

import (
	beans "di_sample/beans"
)

type BeanB struct {
	beanA *BeanA
}

var theBeanB BeanB

func init() {
	theBeanB = BeanB{}
	beans.Register(&theBeanB, []string{"A"})
}

func (this BeanB) getVal() int {
	return this.beanA.val
}

func (this BeanB) setVal(value int) {
	this.beanA.val = value
}

func (this BeanB) GetName() string {
	return "B"
}

func (this *BeanB) Initialize(deps map[string]interface{}) {

	beanARef, ok := deps["A"]
	if !ok {
		panic("BeanB creation error : BeanA not found")
	}

	someBean, ok := beanARef.(*BeanA)
	if !ok {
		panic("Bean B creation error : Bean A type error ")
	}

	this.beanA = someBean
}
