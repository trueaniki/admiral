package admiral

import (
	"io"
	"os"
	"reflect"
)

type command = Command

type Admiral struct {
	command

	Stdout io.Writer
	Exit   func(int)
}

func New(name, description string) *Admiral {
	a := &Admiral{
		command: Command{
			Name:        name,
			Description: description,
		},
		Stdout: os.Stdout,
		Exit:   os.Exit,
	}
	// Set the root command to itself for further propagation
	a.root = a
	// Add help flag to the root command
	a.addHelpFlag()
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
