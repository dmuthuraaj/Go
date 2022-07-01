package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	first_name, last_name string
}

type Employee struct {
	Person
	employee_id string
}

type Customer struct {
	Person
	customer_id, phone_number string
}

type CorporateModel interface {
	FullName() string
	ShowId() string
}

type MarketingModel interface {
	CorporateModel
	ShowPhoneNumber() string
}

func (e Employee) FullName() string {
	return e.Person.last_name + ", " + e.Person.first_name
}

func (e Employee) ShowId() string {
	return e.employee_id
}

func (c Customer) FullName() string {
	return c.Person.first_name + " " + c.Person.last_name
}

func (c Customer) ShowId() string {
	return c.customer_id
}

func (s *Customer) ShowPhoneNumber() string {
	var area_code = (s.phone_number)[0:3]
	var prefix = (s.phone_number)[3:6]
	var line = (s.phone_number)[6:10]
	return "(" + area_code + ")" + prefix + "-" + line
}

func FormatPersonalGreeting(p CorporateModel) string {
	switch reflect.TypeOf(p).String() {
	case "main.Customer":
		return "Welcome valued customer, " + p.FullName() + "!  Your customer id is: " + p.ShowId() + ""
	case "main.Employee":
		return "Employee number: " + p.ShowId() + " (a.k.a. " + p.FullName() + ") - return to work"
	default:
		return "No data found"
	}
}

func main() {
	/*Load some data*/
	drone := Employee{
		Person:      Person{first_name: "Jonah", last_name: "Simms"},
		employee_id: "10ne3mpl0y33",
	}
	patron := Customer{
		Person:       Person{first_name: "Moesha", last_name: "Mitchel"},
		customer_id:  "1amacu5t0m3r",
		phone_number: "2135559753",
	}
	/*end data loading*/
	fmt.Println(" ** Greet Customer **")
	fmt.Println(FormatPersonalGreeting(patron))
	fmt.Printf("We will text you at %s as soon as your pizza is ready!\n\n", patron.ShowPhoneNumber())
	fmt.Println(" ** Greet Employee **")
	fmt.Println(FormatPersonalGreeting(drone))
}
