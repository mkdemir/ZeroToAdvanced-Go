package main

import (
	"fmt"
	"test/models"
)

func main() {

	insan1 := models.NewPerson("Mustafa", "mkdemir")
	insan2 := models.NewPerson("Mustafa", "test")

	models.AddPerson(insan1)
	models.AddPerson(insan2)

	// err := insan.ChangeNickname("fsdfsdfdsfsdfsdfsdfdsf")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Println(*insan1)
	fmt.Println(*insan2)

}
