package entity

import "fmt"

type ExampleEntity struct {
	Name string
}

type ExampleEntityI interface {
	Get()
}

func NewExampleEntity() ExampleEntityI {
	return &ExampleEntity{
		Name: "example entity",
	}
}

func (e *ExampleEntity) Get() {
	fmt.Println("into example entity Get()")
	fmt.Println(e.Name)
}
