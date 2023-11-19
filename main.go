package main

import (
	"errors"
	"fmt"
)

var name string

func enter_club(age int) (string, error) {
	if age >= 18 && age < 70 {
		return fmt.Sprintf("%s, you are allowed in because you're %d years old", name, age), nil
	} else if age >= 70 {
		return fmt.Sprintf("%s, they won't let you in because you're already %d years old", name, age), errors.New("You're to old")
	} else {
		return fmt.Sprintf("%s, they won't let you in because you're only %d years old", name, age), errors.New("You're too young")
	}

}

func main() {

	var mess string
	var err error
	name = "Ivan"
	mess, err = enter_club(18)
	fmt.Println(mess, err)

}
