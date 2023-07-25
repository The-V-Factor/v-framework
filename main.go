package main

import (
	"github.com/lam-guo/v-framework/controller"
	"github.com/lam-guo/v-framework/entity"
	"github.com/lam-guo/v-framework/service"
)

func main() {
	e := entity.NewEntity()
	s := service.NewService(e)
	c := controller.NewController(s)
	c.ExampleController.Get()
}
