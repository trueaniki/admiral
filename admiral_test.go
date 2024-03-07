package admiral

import (
	"testing"
)

type subcommand struct {
	Arg1 string `type:"arg" name:"arg1" pos:"0" description:"First argument"`
	Arg2 string `type:"arg" name:"arg2" pos:"1" description:"Second argument"`
}

type conf struct {
	Sub  subcommand `type:"command" name:"sub" description:"Subcommand"`
	Arg1 string     `type:"arg" name:"arg1" pos:"1" description:"First argument"`
	Arg2 string     `type:"arg" name:"arg2" pos:"0" description:"Second argument"`
}

func TestConfigure(t *testing.T) {
	t.Error("Not implemented")
}

func TestParse(t *testing.T) {
	t.Error("Not implemented")
}

func TestHelp(t *testing.T) {
	t.Error("Not implemented")
}
