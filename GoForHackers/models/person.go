package models

import "errors"

type Person struct {
	name     string
	nickName string
}

var (
	persons []*Person
)

func NewPerson(name, nickName string) *Person {
	p := Person{
		name:     name,
		nickName: nickName,
	}

	return &p
}

func AddPerson(person *Person) {
	persons = append(persons, person)
}

func (p *Person) ChangeNickname(nickName string) error {

	if len(nickName) > 0 {
		return errors.New(":)")
	}
	p.nickName = nickName

	return nil
}
