package service

import (
	"fmt"

	"github.com/lam-guo/v-framework/entity"
)

type ExampleService struct {
	exampleEntity entity.ExampleEntityI
}

type ExampleServiceI interface {
	Get()
}

func NewExampleService(exampleEntity entity.ExampleEntityI) ExampleServiceI {
	return &ExampleService{exampleEntity}
}

func (e *ExampleService) Get() {
	fmt.Println("into example service Get()")
	e.exampleEntity.Get()
}
