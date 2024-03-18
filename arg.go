package admiral

// Arg entity represents a positional argument
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

	// Shows if positional argument is required
	required bool

	// Set value callback to the field in the config struct
	set func(string)

	// Parent command
	parent *Command

	// Callback to be called when positional argument is found in args
	cb func(string)
}

// Adds callback, which will be called when positional argument is found in args
func (a *Arg) Handle(cb func(string)) {
	a.cb = cb
}

// Set value and call all side effects
func (a *Arg) Call(value string) {
	a.Is = true
	a.Value = value
	a.set(value)
	if a.cb != nil {
		a.cb(value)
	}
}

// Set if positional argument is required
func (a *Arg) SetRequired(b bool) *Arg {
	a.required = b
	return a
}
