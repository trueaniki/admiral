package admiral

type Arg struct {
	// Positional argument name
	Name string
	// Positional argument description
	Description string

	// Explicitly set position in the list of positional arguments
	Pos int

	// Shows positional argument presence in args
	Is bool
	// Stores positional argument value as soon as it is found in args
	Value string

	required bool

	set func(string)

	parent *Command

	cb func(string)
}

func (a *Arg) Handle(cb func(string)) {
	a.cb = cb
}

func (a *Arg) Call(value string) {
	a.Is = true
	a.Value = value
	a.set(value)
	if a.cb != nil {
		a.cb(value)
	}
}

func (a *Arg) SetRequired(b bool) *Arg {
	a.required = b
	return a
}
