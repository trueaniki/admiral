package admiral

type Flag struct {
	Name        string
	Alias       string
	Description string

	dataType     string
	defaultValue string
	required     bool

	parent *Command

	cb func()
}

func (f *Flag) Handle(cb func()) {
	f.cb = cb
}
