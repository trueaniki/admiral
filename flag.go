package admiral

import (
	"fmt"
	"strconv"
)

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

func parseFlagValue(value, dataType string) (interface{}, error) {
	switch dataType {
	case "int":
		return parseFlagInt(value)
	case "int64":
		return parseFlagInt64(value)
	case "int32":
		return parseFlagInt32(value)
	case "int16":
		return parseFlagInt16(value)
	case "int8":
		return parseFlagInt8(value)
	case "uint":
		return parseFlagUint(value)
	case "uint64":
		return parseFlagUint64(value)
	case "uint32":
		return parseFlagUint32(value)
	case "uint16":
		return parseFlagUint16(value)
	case "uint8":
		return parseFlagUint8(value)
	case "float":
		return parseFlagFloat(value)
	case "bool":
		return parseFlagBool(value)
	case "string":
		return parseFlagString(value)
	default:
		return nil, fmt.Errorf("unknown data type %s", dataType)
	}
}

func parseFlagInt(value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an integer", value)
	}
	return v, nil
}

func parseFlagInt64(value string) (int64, error) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an integer", value)
	}
	return v, nil
}

func parseFlagInt32(value string) (int32, error) {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an integer", value)
	}
	return int32(v), nil
}

func parseFlagInt16(value string) (int16, error) {
	v, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an integer", value)
	}
	return int16(v), nil
}

func parseFlagInt8(value string) (int8, error) {
	v, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an integer", value)
	}
	return int8(v), nil
}

func parseFlagUint(value string) (uint, error) {
	v, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an unsigned integer", value)
	}
	return uint(v), nil
}

func parseFlagUint64(value string) (uint64, error) {
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an unsigned integer", value)
	}
	return v, nil
}

func parseFlagUint32(value string) (uint32, error) {
	v, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an unsigned integer", value)
	}
	return uint32(v), nil
}

func parseFlagUint8(value string) (uint8, error) {
	v, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an unsigned integer", value)
	}
	return uint8(v), nil
}

func parseFlagUint16(value string) (uint16, error) {
	v, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not an unsigned integer", value)
	}
	return uint16(v), nil
}

func parseFlagFloat(value string) (float64, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("flag %s value is not a float", value)
	}
	return v, nil
}

func parseFlagBool(value string) (bool, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("flag %s value is not a boolean", value)
	}
	return v, nil
}

func parseFlagString(value string) (string, error) {
	return value, nil
}
