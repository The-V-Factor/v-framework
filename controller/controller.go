package controller

import "github.com/lam-guo/v-framework/service"

type Controllers struct {
	ExampleController ExampleControllerI
}

func NewController(services service.Services) Controllers {
	return Controllers{
		ExampleController: NewExampleController(services.ExampleService),
	}
}
