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
	root   *Admiral

	// Returns value of the struct configured to this command
	get func() interface{}

	cb func(opts interface{})
}

// Adds callback, which will be called when command is found in args
func (c *Command) Handle(cb func(opts interface{})) {
	c.cb = cb
}

func (c *Command) Call() {
	c.Is = true
	if c.cb != nil {
		c.cb(c.get())
	}
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

func (c *Command) usage() string {
	s := strings.Builder{}
	path := c.Name
	for p := c.parent; p != nil; p = p.parent {
		path = fmt.Sprintf("%s %s", p.Name, path)
	}
	s.Write([]byte(path))
	if len(c.Flags) > 0 {
		s.Write([]byte(" [options]"))
	}

	if len(c.Commands) > 0 {
		s.Write([]byte(" <command>"))
	}

	for _, arg := range c.Args {
		s.Write([]byte(fmt.Sprintf(" <%s>", arg.Name)))
	}

	return s.String()
}

// Builds help message for command
// TODO: show default values for flags
// TODO: show required flags
// TODO: handle case when flag has no alias
// TODO: show correct usage
func (c *Command) Help() string {
	s := strings.Builder{}

	// Write command usage
	s.Write([]byte(fmt.Sprintf("Usage: %s\n", c.usage())))

	// Write command description
	s.Write([]byte(fmt.Sprintf("\n%s\n", c.Description)))

	// Write subcommands
	if len(c.Commands) > 0 {
		s.Write([]byte("\nCommands:\n"))
		for _, cmd := range c.Commands {
			s.Write([]byte(fmt.Sprintf("  %s\t\t%s\n", cmd.Name, cmd.Description)))
		}
	}

	// Write flags
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
			// Write alias if it exists
			if flag.Alias == "" {
				s.Write([]byte("  "))
			} else {
				s.Write([]byte(fmt.Sprintf("  -%s,", flag.Alias)))
			}
			// Write name
			s.Write([]byte(fmt.Sprintf(" --%s,", flag.Name)))
			// Write gap
			s.Write([]byte(fmt.Sprintf("%s\t", gap)))
			// Write description
			s.Write([]byte(flag.Description))

			// Write default value if it exists
			if flag.defaultValue != "" {
				s.Write([]byte(fmt.Sprintf(" (default: %s)", flag.defaultValue)))
			}

			// Write required flag if it is required
			if flag.required {
				s.Write([]byte(" (required)"))
			}

			// End line
			s.Write([]byte("\n"))
		}
	}

	return s.String()
}
