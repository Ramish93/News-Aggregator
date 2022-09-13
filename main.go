package main

import "pluralsight/go-person/organization"



func main() {
	p:= organization.NewPerson(
		firstName: "ramish", lastName: "hassae"
	)
	println(p.ID())
	println(p.FullName())
}