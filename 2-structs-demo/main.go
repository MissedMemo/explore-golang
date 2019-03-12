package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	p1 := person{
		"John",
		"Smith",
		contactInfo{
			"john@gmail.com",
			94604,
		},
	}

	fmt.Println("p1", p1)

	p2 := person{
		firstName: "Tom",
		lastName:  "Jones",
		contact: contactInfo{
			email:   "tom@yahoo.com",
			zipCode: 12345,
		},
	}

	fmt.Println("p2", p2)

	var p3 person
	p3.firstName = "Sarah"
	p3.lastName = "Jones"
	fmt.Println("p3", p3)
	fmt.Printf("%+v", p3) // prints field names and values
}
