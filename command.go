package admiral

import (
	"fmt"
	"strings"
)

// Command entity represents a command with subcommands, flags and positional arguments
type Command struct {
	// Command name
	Name string
	// Command description
	Description string
	// Subcommands
	Commands []*Command
	// Flags
	Flags []*Flag
	// Positional arguments
	Args []*Arg

	// Shows command presence in args
	Is bool

	// Parent command
	parent *Command
	// Root Admiral instance
	root *Admiral

	// Returns value of the struct configured to this command
	get func() interface{}

	// Callback to be called when command is found in args
	cb func(opts interface{})
}

// Adds callback, which will be called when command is found in args
func (c *Command) Handle(cb func(opts interface{})) {
	c.cb = cb
}

// Sets command presence to true and calls callback
func (c *Command) Call() {
	c.Is = true
	if c.cb != nil {
		c.cb(c.get())
	}
}

// Finds subcommand by name
// Returns nil if subcommand is not found
func (c *Command) Command(name string) *Command {
	for _, cmd := range c.Commands {
		if cmd.Name == name {
			return cmd
		}
	}
	return nil
}

// Finds flag by name
// Returns nil if flag is not found
func (c *Command) Flag(name string) *Flag {
	for _, flag := range c.Flags {
		if flag.Name == name {
			return flag
		}
	}
	return nil
}

// Finds flag by alias
// Returns nil if flag is not found
func (c *Command) FlagByAlias(alias string) *Flag {
	for _, flag := range c.Flags {
		if flag.Alias == alias {
			return flag
		}
	}
	return nil
}

// Finds postional argument by name
// Returns nil if the argument is not found
func (c *Command) Arg(name string) *Arg {
	for _, arg := range c.Args {
		if arg.Name == name {
			return arg
		}
	}
	return nil
}

// Adds arg value to first arg object that has no value
func (c *Command) pushArg(argValue string) {
	for _, arg := range c.Args {
		if !arg.Is {
			arg.Call(argValue)
			return
		}
	}
}

// Add command raw
func (c *Command) addCommand(cmd *Command) {
	if cmd.parent == nil {
		cmd.parent = c
	}
	if cmd.root == nil {
		cmd.root = c.root
	}
	cmd.addHelpFlag()
	c.Commands = append(c.Commands, cmd)
}

func (c *Command) addHelpFlag() {
	if c.Flag("help") == nil {
		c.AddFlag("help", "h", "Show help").Handle(func(_ interface{}) {
			c.root.Stdout.Write([]byte(c.Help()))
			c.root.Exit(0)
		})
	}
}

// Add subcommand
func (c *Command) AddCommand(name, description string) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		parent:      c,
		// Propagate root reference
		root: c.root,
	}
	c.addCommand(command)
	return command
}

// Add flag raw
func (c *Command) addFlag(flag *Flag) {
	if flag.parent == nil {
		flag.parent = c
	}
	if flag.dataType == "" {
		flag.dataType = "bool"
	}
	c.Flags = append(c.Flags, flag)
}

// Add flag
func (c *Command) AddFlag(name, alias, description string) *Flag {
	flag := &Flag{
		Name:        name,
		Alias:       alias,
		Description: description,
		parent:      c,
	}
	c.addFlag(flag)
	return flag
}

// Add arg raw
func (c *Command) addArg(arg *Arg) {
	if arg.parent == nil {
		arg.parent = c
	}
	// If specific position is set, add arg to that position
	if arg.pos != -1 {
		for len(c.Args) <= arg.pos {
			c.Args = append(c.Args, nil)
		}
		c.Args[arg.pos] = arg
		// If no position is set, add arg to the end
	} else {
		c.Args = append(c.Args, arg)
	}
}

// Add arg
func (c *Command) AddArg(name, description string) {
	arg := &Arg{
		Name:        name,
		Description: description,
		parent:      c,
	}
	c.addArg(arg)
}

// Build usage string for command
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

var spaces4 = strings.Repeat(" ", 4)
var spaces5 = strings.Repeat(" ", 5)

// Builds help message for command
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
			alias := flag.Alias
			if alias == "" {
				alias = " "
			}
			if len(flag.Name+alias) > longest {
				longest = len(alias + flag.Name)
			}
		}

		for _, flag := range c.Flags {
			alias := flag.Alias
			if alias == "" {
				alias = " "
			}
			length := len(alias + flag.Name)
			gap := strings.Repeat(" ", longest-length)
			// Write alias
			if alias != " " {
				s.Write([]byte(fmt.Sprintf("  -%s,", alias)))
			} else {
				s.Write([]byte(spaces5))
			}
			// Write name
			s.Write([]byte(fmt.Sprintf(" --%s", flag.Name)))
			// Write gap
			s.Write([]byte(fmt.Sprintf("%s%s", gap, spaces4)))
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
