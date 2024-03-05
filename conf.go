package admiral

import "reflect"

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

func buildFlag(f reflect.StructField) *Flag {
	tag := f.Tag

	name := tag.Get("name")
	alias := tag.Get("alias")
	description := tag.Get("description")
	fieldType := f.Type.String()
	defaultValue := tag.Get("default")
	required := tag.Get("required")

	return &Flag{
		Name:         name,
		Alias:        alias,
		Description:  description,
		dataType:     fieldType,
		defaultValue: defaultValue,
		required:     required == "true",
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
	}

	cmd.applyConfig(v)

	return cmd
}
