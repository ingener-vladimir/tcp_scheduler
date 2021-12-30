package model

type Person struct {
	id    int    `json:"id"`
	age   int    `json:"age"`
	name  string `json:"name"`
	isMan bool   `json:"isMan"`
}

func New(id int, age int, name string, isMan bool) Person {
	return Person{
		id:    id,
		age:   age,
		name:  name,
		isMan: isMan,
	}
}
