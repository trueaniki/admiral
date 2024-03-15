# Admiral
*Flags & Commands*

Command line arguments parser for humans.

## TL;DR
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
	cli := admiral.New("myapp", "My app")
	conf := Conf{}
	cli.Configure(&conf)
	cli.Flag("version").Handle(func(value interface{}) {
		fmt.Println("Version 1.0.0")
		os.Exit(0)
	})

	_, err := cli.Parse(os.Args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", conf)
}
```


## Core concepts
### Command
Entity representing commands or subcommand of application
### Admiral
Root command represnting the binary
### Arg
Positinal argument of a command
### Flag
Flag (argument with '--' or '-' prefix) of a command

## Usage
First thing to do is to initialize `admiral` instance:
```go
cli := admiral.New("myApp", "My App")
```
Here first argument is an entryopoint name and second is application description. Both are used in help message.
### Configuring using struct
You can use structs to configure the parser:
```go
type Conf struct {
	Host   string           `type:"flag" name:"host" alias:"h" description:"Host to listen on" required:"true"`
	Port   int              `type:"flag" name:"port" alias:"p" description:"Port to listen on" default:"8080"`
	Listen ListenSubcommand `type:"command" name:"listen" description:"Listen specified host and port"`
}
```
Each entity is defined using tags. The tags are:
- `type` - defines entity type: *flag, arg, command*.
- `name` - entity name.

	For flags acts like a long name when using in cli. For `type:"flag" name:"example"` flag is available as `--example`.

	For commands define callable name of the command. For `type:"command" name:"example"` command is used like `app example`.

	For arguments name appear in help message.
- `alias` - Flag alias.
- `description` - Entity description. Appear in help message.
- `required` - Is flag or argument required or not for this command.
- `default` - Default value for flag.
- `pos` - Positional argument's position starting from 0. If used for one argument, should be used for all.
Flag data type is recognized automatically based on property type.
Argument property must have `string` type.

After defining structure make variable of it:
```go
conf := &Conf{}
```
Then call `Configure` method:
```go
cli.Configure(conf)
```
Then call `Parse` method:
```go
rest, err := cli.Parse(os.Args)
```
`rest` variable contains all cli argument that were not parsed for some reason.
Use `conf` variable to access parsed data.
#### Configuring subcommands
To configure subcommand you can use structs with same structure as root struct.
```go
type Conf struct {
	Listen Listen `type:"command" name:"listen" description:"Listen specified host and port"`
}

type Listen struct {
	Host   string           `type:"flag" name:"host" alias:"h" description:"Host to listen on" required:"true"`
	Port   int              `type:"flag" name:"port" alias:"p" description:"Port to listen on" default:"8080"`
}
```

## Configuring using methods
It's also possible to configure using methods instead of using structs:
```go
cli := admiral.New("myApp", "My App")
cli.AddFlag("host", "h", "Host to listen on").SetType("string").SetRequired(true)
cli.AddFlag("port", "p", "Port to listen on").SetType("int").SetDefault("8080")

cli.AddCommand("listen")

cli.Command("listen").
	AddFlag("host", "h", "Host to listen on").
	SetDataType("string").
	SetRequired(true)

cli.Command("listen").
	AddFlag("port", "p", "Port to listen on").
	SetType("int").
	SetDefault("8080")
```
After parsing arguments:
```go
rest, err := cli.Parse(os.Args)
```
You can access result by using `.Is` and `.Value` properties. `.Is` stands for entity presence in arguments. `.Value` contains actual value.
```go
if !cli.Command("listen").Is {
	host := cli.Flag("host").Value
	port := cli.Flag("port").Value
} else {
	host := cli.Command("listen").Flag("host").Value
	port := cli.Command("listen").Flag("port").Value
}

```
## Adding handlers
*TBA*
