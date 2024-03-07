# Admiral
*Flags & Commands*

Command line arguments parser for humans.

## Usage
### TL;DR
```go
type Conf struct {
	Version bool `type:"flag" name:"version" alias:"v" description:"Show version"`

	Network  string `type:"flag" name:"network" alias:"n" description:"Network in CIDR format" required:"true"`
	Confpath string `type:"flag" name:"conf" alias:"c" description:"Path to config file" default:"/etc/app.conf"`

	Start Start `type:"command" name:"start" description:"Start the app in detached mode"`
	Stop  Stop  `type:"command" name:"stop" description:"Stop the app"`
}

type Start struct {
	Network  string `type:"flag" name:"network" alias:"n" description:"Network in CIDR format" required:"true"`
	Logfile  string `type:"flag" name:"logfile" alias:"l" description:"Logfile path" default:"/var/log/app.log"`
	Confpath string `type:"flag" name:"conf" alias:"c" description:"Path to config file" default:"/etc/app.conf"`
}

type Stop struct{}

func main() {
	a := admiral.New("myapp", "My app")
	conf := Conf{}
	a.Configure(&conf)
	a.Flag("version").Handle(func(value interface{}) {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	})

	_, err := a.Parse(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", conf)
}
```
