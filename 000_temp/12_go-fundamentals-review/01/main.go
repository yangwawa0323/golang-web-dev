package main

import (
	"fmt"
	"strconv"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	fname string
	// lname         string
	licenseToKill bool
}

func (p person) speak() string {
	return "Yo, what's up"
}
func (p person) intro() string {
	return p.fname + " " + p.lname
}

func (sa secretAgent) speak() string {
	return "shaken, not stirred"
}

/*
	If comment the `secretAgent` receiver `intro()`` function,
	it's will call `person` receiver `intro()``
	because in the struct secretAgent included the `person`
	struct and with an additional attribte `licenseToKill`.
	but if you redefined the struct like
	`type secretAgent struct{
		fname string
		lname string
		licenseToKill bool
	}`
	, above struct will not work by lacking of `intro()` function.
*/
func (sa secretAgent) intro() string {
	return sa.fname + " " + sa.lname +
		" has licenseToKill: " + strconv.FormatBool(sa.licenseToKill)
}

type human interface {
	speak() string
	intro() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func introduce(h human) {
	fmt.Println(h.intro())
}

func main() {

	p1 := new(person)
	// p1 := person{
	// 	"Miss",
	// 	"Moneypenny",
	// }
	p1.fname = "Miss"
	p1.lname = "Moneypenny"

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	saySomething(p1)
	introduce(p1)
	saySomething(sa1)
	introduce(sa1)
}
