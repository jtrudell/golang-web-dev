// - create a struct that holds person fields
// - create a struct that holds secret agent fields and embeds person type
// - attach a method to person: pSpeak
// - attach a method to secret agent: saSpeak
// - create a variable of type person
// - create a variable of type secret agent
// - print a field from person
// - run pSpeak attached to the variable of type person
// - print a field from secret agent
// - run saSpeak attached to the variable of type secret agent
// - run pSpeak attached to the variable of type secret agent

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

func (p person) pSpeak() {
	fmt.Println("Hey there, Secret Agent. What's your badge number?")
}

func (sa secretAgent) saSpeak() {
	fmt.Println("My bade number is", sa.badgeNumber)
}

func main() {
	p := person{"Jill Scott"}
	sa := secretAgent{person{"James Scott"}, 999}
	fmt.Println("Person's name is", p.name)
	p.pSpeak()
	fmt.Println("Secret Agent's name is", sa.name)
	sa.saSpeak()
	sa.pSpeak()
}
