package model

import "fmt"

// что-то с координатами
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

func (p Person) String() string {
	return fmt.Sprintf("имя: %s, возраст: %d, мужчина: %v, номер: %d", p.Name, p.Age, p.IsMan, p.Id)
}
