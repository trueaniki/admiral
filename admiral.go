package admiral

import (
	"io"
	"os"
	"reflect"
)

type command = Command

// Admiral is the root command and therefore the entry point of the application.
type Admiral struct {
	// Derive from Command.
	command

	// Allows to set writer for help messages.
	Stdout io.Writer
	// Allows to set exit function
	// which is called after help message is shown
	// or when an error occurs.
	Exit func(int)
}

// Create a new Admiral instance
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

// Parse and apply config struct to create commands, flags and args.
// Panics if the passed interface is not a pointer to a struct
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
