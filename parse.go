package admiral

import "errors"

func (c *Command) Parse(args []string) ([]string, error) {
	rest := []string{}

	p := c

	for i, arg := range args {
		// TODO: arguments should not be parsed like this
		// TODO: handle flag errors intstead of panicking
		flagParsed, err := p.tryParseFlag(i, args)
		if err != nil {
			return nil, err
		}
		if !flagParsed {
			if p.tryParseCommand(arg) {
				// TODO: call callback with parsed falgs and args
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

	// Call command callback
	p.Call()

	return rest, nil
}

func (c *Command) tryParseCommand(argument string) bool {
	// Check if arg is a command
	if cmd := c.Command(argument); cmd != nil {
		cmd.Is = true
		return true
	}
	return false
}

func (c *Command) tryParseArg(argValue string) bool {
	// Check if command has any args configured
	if len(c.Args) == 0 {
		return false
	}
	c.pushArg(argValue)
	return true
}
