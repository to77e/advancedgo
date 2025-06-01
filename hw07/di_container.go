package hw07

import "errors"

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	providers map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		providers: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.providers[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, ok := c.providers[name]
	if !ok {
		return nil, errors.New("service not found")
	}

	fn, ok := constructor.(func() interface{})
	if !ok {
		return nil, errors.New("invalid constructor type")
	}
	return fn(), nil
}
