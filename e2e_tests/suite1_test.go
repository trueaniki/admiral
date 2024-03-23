package admiral_test

import (
	"testing"

	"github.com/trueaniki/admiral"
)

type Conf struct {
	Sub   Subcommand `type:"command" name:"sub" description:"Subcommand"`
	Arg1  string     `type:"arg" name:"arg1" pos:"1" description:"First argument"`
	Arg2  string     `type:"arg" name:"arg2" pos:"0" description:"Second argument"`
	Flag1 string     `type:"flag" name:"flag1" alias:"f" description:"First flag" required:"true"`
	Flag2 string     `type:"flag" name:"flag2" alias:"g" description:"Second flag" default:"default"`
}

type Subcommand struct {
	Arg1  string `type:"arg" name:"arg1" pos:"0" description:"First argument"`
	Arg2  string `type:"arg" name:"arg2" pos:"1" description:"Second argument"`
	Flag1 string `type:"flag" name:"flag1" alias:"f" description:"First flag" required:"true"`
	Flag2 string `type:"flag" name:"flag2" alias:"g" description:"Second flag" default:"default"`
}

func TestSuite1_1(t *testing.T) {
	a := admiral.New("test", "Test")
	c := &Conf{}
	a.Configure(c)

	a.Parse([]string{"sub", "arg2", "arg1", "-f", "flag1"})

	if a.Command("sub") == nil {
		t.Error("Expected subcommand")
	}
	sub := a.Command("sub")
	if !sub.Is {
		t.Error("Expected subcommand to be set")
	}

	if !sub.Arg("arg1").Is {
		t.Error("Expected sub arg1 to be set")
	}
	if sub.Arg("arg1").Value != "arg1" {
		t.Error("Expected sub arg1 to be arg1")
	}
	if !sub.Arg("arg2").Is {
		t.Error("Expected sub arg2 to be set")
	}
	if sub.Arg("arg2").Value != "arg2" {
		t.Error("Expected sub arg2 to be arg2")
	}
	if !sub.Flag("flag1").Is {
		t.Error("Expected sub flag1 to be set")
	}
	if sub.Flag("flag1").Value != "flag1" {
		t.Error("Expected sub flag1 to be flag1")
	}
}

func TestSuite1_2(t *testing.T) {
	a := admiral.New("test", "Test")
	c := &Conf{}
	a.Configure(c)

	a.Parse([]string{"arg2", "arg1", "-f", "flag1"})

	if !a.Arg("arg1").Is {
		t.Error("Expected arg1 to be set")
	}
	if a.Arg("arg1").Value != "arg1" {
		t.Error("Expected arg1 to be arg1")
	}
	if !a.Arg("arg2").Is {
		t.Error("Expected arg2 to be set")
	}
	if a.Arg("arg2").Value != "arg2" {
		t.Error("Expected arg2 to be arg2")
	}
	if !a.Flag("flag1").Is {
		t.Error("Expected flag1 to be set")
	}
	if a.Flag("flag1").Value != "flag1" {
		t.Error("Expected flag1 to be flag1")
	}
	if !a.Flag("flag2").Is {
		t.Error("Expected flag2 to be set")
	}
	if a.Flag("flag2").Value != "default" {
		t.Error("Expected flag2 to be default")
	}
}
