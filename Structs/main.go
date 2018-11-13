package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	sex       string
	age       int
	contact   contactInfo
}
type contactInfo struct {
	email      string
	postalCode string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	jono := person{
		firstName: "jono",
		lastName:  "echols",
		sex:       "m",
		age:       32,
		contact: contactInfo{
			email:      "jfg.echols@gmail.com",
			postalCode: "v6k1p6",
		},
	}
	jono.updateName("jonathan")
	jono.print()

}
