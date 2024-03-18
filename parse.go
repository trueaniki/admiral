package admiral

import "errors"

// Parse command line arguments.
// Sets values to the struct fields and calls callbacks.
// Returns rest of the arguments that were not parsed for any reason.
func (c *Command) Parse(args []string) ([]string, error) {
	// Remove the first argument, which is the application name
	args = args[1:]

	rest := []string{}
	p := c

	for i, arg := range args {
		// TODO: arguments should not be parsed like this
		flagParsed, err := p.tryParseFlag(i, args)
		if err != nil {
			return nil, err
		}
		if !flagParsed {
			if p.tryParseCommand(arg) {
				p = p.Command(arg)
			} else if !p.tryParseArg(arg) {
				rest = append(rest, arg)
			}
		}
	}

	// Handle default flag value and required flags
	for _, flag := range p.Flags {
		if !flag.Is && flag.defaultValue != "" {
			v, err := parseFlagValue(flag.defaultValue, flag.dataType)
			if err != nil {
				return nil, err
			}
			flag.Call(v)
		}
		if flag.required && !flag.Is {
			return nil, errors.New("Flag " + flag.Name + " is required")
		}
	}

	// Handle required args
	for _, arg := range p.Args {
		if !arg.Is && arg.required {
			return nil, errors.New("Argument " + arg.Name + " is required")
		}
	}

	// Call command callback
	p.Call()

	return rest, nil
}

// Try to parse command
func (c *Command) tryParseCommand(argument string) bool {
	// Check if arg is a command
	if cmd := c.Command(argument); cmd != nil {
		cmd.Is = true
		return true
	}
	return false
}

// Try to parse argument
func (c *Command) tryParseArg(argValue string) bool {
	// Check if command has any args configured
	if len(c.Args) == 0 {
		return false
	}
	c.pushArg(argValue)
	return true
}
