package models

import "errors"

type Person struct { // struck class'a denk geliyor. Value object
	id        int64
	firstName string
	lastName  string
	nickName  string
}

func NewPerson(firstName string, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}

}

func (p *Person) SetNickName(nickName string) error {

	if len(nickName) == 0 {
		return errors.New("Nickname must not be empty")
	}

	p.nickName = nickName
	return nil
}
