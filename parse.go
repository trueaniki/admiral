package admiral

func (c *Command) Parse(args []string) []string {
	rest := []string{}

	p := c

	for i, arg := range args {
		if !p.tryParseFlag(i, args) && !p.tryParseArg(arg) {
			if p.tryParseCommand(arg) {
				p = p.Command(arg)
			} else {
				rest = append(rest, arg)
			}
		}
	}

	return rest
}

func (c *Command) tryParseCommand(argument string) bool {
	// Check if arg is a command
	if cmd := c.Command(argument); cmd != nil {
		cmd.Is = true
		return true
	}
	return false
}

func (c *Command) tryParseArg(argument string) bool {
	// Check if arg is an argument
	if arg := c.Arg(argument); arg != nil {
		arg.Is = true
		arg.Value = argument
		arg.set(argument)
		return true
	}
	return false
}
