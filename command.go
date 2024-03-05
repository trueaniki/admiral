package admiral

import (
	"fmt"
	"reflect"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Commands    []*Command
	Flags       []*Flag
	Args        []*Arg

	Is bool

	parent *Command

	cb func()
}

// Adds callback, which will be called when command is found in args
func (c *Command) Handle(cb func()) {
	c.cb = cb
}

// Finds subcommand by name
func (c *Command) Command(name string) *Command {
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

// Finds flag by name
func (c *Command) Flag(name string) *Flag {
	for _, flag := range c.Flags {
		if flag.Name == name {
			return flag
		}
	}
	return nil
}

func (c *Command) FlagByAlias(alias string) *Flag {
	for _, flag := range c.Flags {
		if flag.Alias == alias {
			return flag
		}
	}
	return nil
}

func (c *Command) Arg(name string) *Arg {
	for _, arg := range c.Args {
		if arg.Name == name {
			return arg
		}
	}
	return nil
}

func (c *Command) addCommand(cmd *Command) {
	if cmd.parent == nil {
		cmd.parent = c
	}
	if cmd.cb == nil {
		cmd.cb = func() {}
	}
	c.Commands = append(c.Commands, cmd)
}

// Adds subcommand
func (c *Command) AddCommand(name, description string) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		parent:      c,
	}
	c.addCommand(command)
	return command
}

func (c *Command) addFlag(flag *Flag) {
	if flag.parent == nil {
		flag.parent = c
	}
	if flag.cb == nil {
		flag.cb = func() {}
	}
	c.Flags = append(c.Flags, flag)
}

// Adds flag
func (c *Command) AddFlag(name, alias, description string) *Flag {
	flag := &Flag{
		Name:        fmt.Sprintf("--%s", name),
		Alias:       fmt.Sprintf("-%s", alias),
		Description: description,
		parent:      c,
	}
	c.addFlag(flag)
	return flag
}

func (c *Command) addArg(arg *Arg) {
	if arg.parent == nil {
		arg.parent = c
	}
	c.Args = append(c.Args, arg)
}

func (c *Command) AddArg(name, description string) {
	arg := &Arg{
		Name:        name,
		Description: description,
		parent:      c,
	}
	c.addArg(arg)
}

// Builds help message for command
func (c *Command) Help() string {
	s := strings.Builder{}

	path := c.Name
	for p := c.parent; p != nil; p = p.parent {
		path = fmt.Sprintf("%s %s", p.Name, path)
	}

	s.Write([]byte(fmt.Sprintf("Usage: %s\n", path)))
	s.Write([]byte(fmt.Sprintf("\n%s\n", c.Description)))

	if len(c.Commands) > 0 {
		s.Write([]byte("\nCommands:\n"))
		for _, cmd := range c.Commands {
			s.Write([]byte(fmt.Sprintf("  %s\t\t%s\n", cmd.Name, cmd.Description)))
		}
	}

	if len(c.Flags) > 0 {
		s.Write([]byte("\nOptions:\n"))

		longest := 0
		for _, flag := range c.Flags {
			if len(flag.Name+flag.Alias) > longest {
				longest = len(flag.Alias + flag.Name)
			}
		}

		for _, flag := range c.Flags {
			len := len(flag.Alias + flag.Name)
			gap := strings.Repeat(" ", longest-len)
			s.Write([]byte(fmt.Sprintf("  %s, %s%s\t%s\n", flag.Alias, flag.Name, gap, flag.Description)))
		}
	}

	return s.String()
}

func (c *Command) applyConfig(v reflect.Value) {
	// Iterate over the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)

		if isFlagField(field) {
			c.addFlag(buildFlag(field, v.Field(i)))
		} else if isArgField(field) {
			c.addArg(buildArg(field, v.Field(i)))
		} else if isCommandField(field) {
			c.addCommand(buildCommand(field, v.Field(i)))
		}
	}
}
