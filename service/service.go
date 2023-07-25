package service

import "github.com/lam-guo/v-framework/entity"

type Services struct {
	ExampleService ExampleServiceI
}

func NewService(entities entity.Entities) Services {
	return Services{
		ExampleService: NewExampleService(entities.ExampleEntity),
	}
}
