package admiral

import (
	"testing"
)

type subcommand struct {
	Arg1  string `type:"arg" name:"arg1" pos:"0" description:"First argument"`
	Arg2  string `type:"arg" name:"arg2" pos:"1" description:"Second argument"`
	Flag1 string `type:"flag" name:"flag1" alias:"f" description:"First flag" required:"true"`
	Flag2 string `type:"flag" name:"flag2" alias:"g" description:"Second flag" default:"default"`
}

type conf struct {
	Sub   subcommand `type:"command" name:"sub" description:"Subcommand"`
	Arg1  string     `type:"arg" name:"arg1" pos:"1" description:"First argument"`
	Arg2  string     `type:"arg" name:"arg2" pos:"0" description:"Second argument"`
	Flag1 string     `type:"flag" name:"flag1" alias:"f" description:"First flag" required:"true"`
	Flag2 string     `type:"flag" name:"flag2" alias:"g" description:"Second flag" default:"default"`
}

func TestConfigure(t *testing.T) {
	a := New("test", "Test")
	c := conf{}
	a.Configure(&c)

	if len(a.Commands) != 1 {
		t.Errorf("Expected 1 command, got %d", len(a.Commands))
	}
	if len(a.Args) != 2 {
		t.Errorf("Expected 2 args, got %d", len(a.Args))
	}
	if len(a.Flags) != 2 {
		t.Errorf("Expected 2 flags, got %d", len(a.Flags))
	}

	if a.Command("sub") == nil {
		t.Error("Expected subcommand")
	}
	if a.Arg("arg1") == nil {
		t.Error("Expected arg1")
	}
	if a.Arg("arg2") == nil {
		t.Error("Expected arg2")
	}
	if a.Flag("flag1") == nil {
		t.Error("Expected flag1")
	}
	if a.Flag("flag2") == nil {
		t.Error("Expected flag2")
	}

	if a.Command("sub").Description != "Subcommand" {
		t.Errorf("Expected description 'Subcommand', got '%s'", a.Command("sub").Description)
	}
	if a.Arg("arg1").Description != "First argument" {
		t.Errorf("Expected description 'First argument', got '%s'", a.Arg("arg1").Description)
	}
	if a.Arg("arg2").Description != "Second argument" {
		t.Errorf("Expected description 'Second argument', got '%s'", a.Arg("arg2").Description)
	}
	if a.Flag("flag1").Description != "First flag" {
		t.Errorf("Expected description 'First flag', got '%s'", a.Flag("flag1").Description)
	}
	if a.Flag("flag2").Description != "Second flag" {
		t.Errorf("Expected description 'Second flag', got '%s'", a.Flag("flag2").Description)
	}

	if a.Flag("flag1").required != true {
		t.Errorf("Expected required flag1, got %t", a.Flag("flag1").required)
	}
	if a.Flag("flag2").required != false {
		t.Errorf("Expected required flag2, got %t", a.Flag("flag2").required)
	}
	if a.Flag("flag2").defaultValue != "default" {
		t.Errorf("Expected default 'default', got %s", a.Flag("flag2").defaultValue)
	}

	if a.Args[0].Name != "arg2" {
		t.Errorf("Expected arg2, got %s", a.Args[0].Name)
	}
	if a.Args[1].Name != "arg1" {
		t.Errorf("Expected arg1, got %s", a.Args[1].Name)
	}

}

func TestParse(t *testing.T) {
	t.Error("Not implemented")
}

func TestHelp(t *testing.T) {
	t.Error("Not implemented")
}
