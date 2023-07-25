package entity

type Entities struct {
	ExampleEntity ExampleEntityI
}

func NewEntity() Entities {
	return Entities{
		ExampleEntity: NewExampleEntity(),
	}
}
