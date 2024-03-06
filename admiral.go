package admiral

import "reflect"

type command = Command

type Admiral struct {
	command
}

func New(name, description string) *Admiral {
	a := &Admiral{
		command: Command{
			Name:        name,
			Description: description,
		},
	}
	// Set the root command to itself for further propagation
	a.root = a
	return a
}

func (a *Admiral) Configure(conf interface{}) {
	v := reflect.ValueOf(conf)
	// Check if the passed interface is a pointer and points to a struct
	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		v = v.Elem() // Dereference the pointer to get the underlying struct
	} else {
		panic("Config must be a pointer to a struct")
	}
	a.command.applyConfig(v)
}
