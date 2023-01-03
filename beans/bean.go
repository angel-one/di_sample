package beans

import (
	"log"
	"reflect"
	"sync"
)

type Bean interface {
	GetName() string
	Initialize(map[string]interface{})
}

var deps map[string][]string
var repository map[string]interface{}
var nameToBean map[string]Bean
var mutex sync.RWMutex

func init() {
	deps = make(map[string][]string)
	repository = make(map[string]interface{})
	nameToBean = make(map[string]Bean)
}

func Register(bean Bean, requires []string) {

	mutex.Lock()
	defer mutex.Unlock()

	name := bean.GetName()
	log.Println("Registering ", name, " with deps ", deps[name])

	deps[name] = []string{}
	for _, dep := range requires {
		deps[name] = append(deps[name], dep)
	}

	nameToBean[name] = bean
	origRef := reflect.ValueOf(bean).Interface()

	repository[name] = origRef

}

func Initialise() {
	/*
	   Does a DFS based topological sort .
	   Uses inWork as  recurrsion stack to keep track of cycles in graph
	*/
	visited := make(map[string]bool)
	inWork := make(map[string]bool)
	for name, _ := range repository {
		helper(name, visited, inWork)
	}

}

func helper(beanName string, visited map[string]bool, inWork map[string]bool) {

	if _, exists := visited[beanName]; exists {
		return
	}
	log.Println("Initialising bean ", beanName)
	if _, exists := inWork[beanName]; exists {
		panic("cycle in deps!!")
	}

	inWork[beanName] = true

	if len(deps[beanName]) != 0 {
		for _, dep := range deps[beanName] {

			helper(dep, visited, inWork)

		}
	}

	delete(inWork, beanName)
	visited[beanName] = true

	bean := nameToBean[beanName]
	(bean).Initialize(repository)
	return

}
