package main

import (
	"fmt"
	"os"

	"github.com/trueaniki/admiral"
)

type ListenSubcommand struct {
	Host string `type:"flag" name:"host" alias:"h" description:"Host to listen on" required:"true"`
	Port int    `type:"flag" name:"port" alias:"p" description:"Port to listen on" default:"8080" required:"true"`
}

type Conf struct {
	Host   string           `type:"flag" name:"host" alias:"h" description:"Host to listen on" required:"true"`
	Port   int              `type:"flag" name:"port" alias:"p" description:"Port to listen on" default:"8080" required:"true"`
	Listen ListenSubcommand `type:"command" name:"listen" description:"Listen specified host and port"`
}

func main() {
	a := admiral.New("example", "Example app")
	conf := Conf{}
	a.Configure(&conf)
	a.Parse(os.Args)
	fmt.Printf("%+v\n", conf)
	fmt.Println(a.Help())
}
