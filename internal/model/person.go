package model

type Person struct {
	Id    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	IsMan bool   `json:"isMan"`
}

func New(id int, age int, name string, isMan bool) Person {
	return Person{
		Id:    id,
		Age:   age,
		Name:  name,
		IsMan: isMan,
	}
}
