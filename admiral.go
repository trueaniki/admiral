package admiral

type command = Command

type Admiral struct {
	command
}

func New(name, description string) *Admiral {
	return &Admiral{
		command: Command{
			Name:        name,
			Description: description,
		},
	}
}

func (p *Admiral) Parse() {

}
