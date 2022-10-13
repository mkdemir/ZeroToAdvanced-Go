package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() { // struct metodu oldu
	fmt.Println("Eklendi : ", c.firstName)
}

func (c customer) update() { // struct metodu oldu
	fmt.Println("Update : ", c.firstName)
}

func Demo2() {
	c := customer{firstName: "Mustafa", lastName: "Demir", age: 23}
	c.save()
	c.update()
}
