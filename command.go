package admiral

import (
	"fmt"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Commands    []*Command
	Flags       []*Flag
	Args        []string

	parent *Command

	cb func()
}

func (c *Command) Handle(cb func()) {
	c.cb = cb
}

func (c *Command) Command(name string) *Command {
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

func (c *Command) Flag(name string) *Flag {
	for _, flag := range c.Flags {
		if flag.Name == name {
			return flag
		}
	}
	return nil
}

func (c *Command) AddCommand(name, description string) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		parent:      c,
	}
	c.Commands = append(c.Commands, command)
	return command
}

func (c *Command) AddFlag(name, alias, description string) *Flag {
	flag := &Flag{
		Name:        fmt.Sprintf("--%s", name),
		Alias:       fmt.Sprintf("-%s", alias),
		Description: description,
		parent:      c,
	}
	c.Flags = append(c.Flags, flag)
	return flag
}

func (c *Command) AddArg(name, description string) {
	c.Args = append(c.Args, name)
}

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
