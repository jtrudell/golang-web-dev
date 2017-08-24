// - create an interface type that both person and secretAgent implement
// - declare a func with a parameter of the interface’s type
// - call that func in main and pass in a value of type person
// - call that func in main and pass in a value of type secretAgent

package main

import (
	"fmt"
)

type person struct {
	name string
}

type secretAgent struct {
	person
	badgeNumber int
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Hey there, Secret Agent. What's your badge number?")
}

func (sa secretAgent) speak() {
	fmt.Println("My badge number is", sa.badgeNumber)
}

func talk(h human) {
	h.speak()
}

func main() {
	p := person{"Jill Scott"}
	sa := secretAgent{person{"James Scott"}, 999}
	talk(p)
	talk(sa)
}
