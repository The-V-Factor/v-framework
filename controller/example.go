package controller

import (
	"fmt"

	"github.com/lam-guo/v-framework/service"
)

type ExampleController struct {
	exampleService service.ExampleServiceI
}

type ExampleControllerI interface {
	Get()
}

func NewExampleController(exampleService service.ExampleServiceI) ExampleControllerI {
	return &ExampleController{exampleService: exampleService}
}

func (e *ExampleController) Get() {
	fmt.Println("into example controller Get()")
	e.exampleService.Get()
}
