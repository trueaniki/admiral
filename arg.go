package admiral

type Arg struct {
	Name        string
	Description string
	Commands    []Command

	Is    bool
	Value string

	set func(string)

	parent *Command
}
