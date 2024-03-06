package admiral

type Arg struct {
	Name        string
	Description string
	// Commands    []Command

	Pos int

	Is    bool
	Value string

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
