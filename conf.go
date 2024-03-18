package admiral

import (
	"reflect"
	"strconv"
)

func (c *Command) applyConfig(v reflect.Value) {
	// Iterate over the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)

		if isFlagField(field) {
			c.addFlag(buildFlag(field, v.Field(i)))
		} else if isArgField(field) {
			c.addArg(buildArg(field, v.Field(i)))
		} else if isCommandField(field) {
			c.addCommand(buildCommand(field, v.Field(i)))
		}
	}
}

func isFlagField(f reflect.StructField) bool {
	tag := f.Tag
	explicit := tag.Get("type") == "flag"
	implicit := tag.Get("name") != "" && tag.Get("alias") != ""
	return explicit || implicit
}

func isArgField(f reflect.StructField) bool {
	tag := f.Tag
	explicit := tag.Get("type") == "arg"
	implicit := tag.Get("name") != "" &&
		f.Type.Kind() == reflect.String &&
		!isFlagField(f) &&
		!isCommandField(f)
	return explicit || implicit
}

func isCommandField(f reflect.StructField) bool {
	tag := f.Tag
	explicit := tag.Get("type") == "command"
	implicit := tag.Get("name") != "" && f.Type.Kind() == reflect.Struct
	return explicit || implicit
}

func buildFlag(f reflect.StructField, v reflect.Value) *Flag {
	tag := f.Tag

	name := tag.Get("name")
	alias := tag.Get("alias")
	description := tag.Get("description")
	fieldType := f.Type.String()
	defaultValue := tag.Get("default")
	required := tag.Get("required")

	return &Flag{
		Name:        name,
		Alias:       alias,
		Description: description,

		dataType: fieldType,

		defaultValue: defaultValue,
		required:     required == "true",

		set: func(value interface{}) {
			v.Set(reflect.ValueOf(value))
		},
	}
}

func buildCommand(f reflect.StructField, v reflect.Value) *Command {
	t := v.Type()
	tag := f.Tag

	name := tag.Get("name")
	description := tag.Get("description")

	if t.Kind() != reflect.Struct {
		panic("Command field must be a struct")
	}

	cmd := &Command{
		Name:        name,
		Description: description,
		Commands:    []*Command{},
		Flags:       []*Flag{},
		Args:        []*Arg{},

		get: func() interface{} {
			return v.Addr().Interface()
		},
	}

	cmd.applyConfig(v)

	return cmd
}

func buildArg(f reflect.StructField, v reflect.Value) *Arg {
	tag := f.Tag

	name := tag.Get("name")
	description := tag.Get("description")
	posStr, hasPos := tag.Lookup("pos")
	if !hasPos {
		posStr = "-1"
	}
	pos, err := strconv.Atoi(posStr)
	if err != nil {
		panic("Arg position must be an integer")
	}
	required, hasRequired := tag.Lookup("required")
	if !hasRequired {
		required = "false"
	}

	return &Arg{
		Name:        name,
		Description: description,
		pos:         pos,
		required:    required == "true",
		set: func(value string) {
			v.SetString(value)
		},
	}
}
