package admiral

import (
	"errors"
	"strings"
)

func (c *Command) tryParseFlag(i int, args []string) bool {
	flagName := args[i]

	// Determine flag value
	flagValue := ""
	if strings.Contains(flagName, "=") {
		flagValue = strings.Split(flagName, "=")[1]
		flagName = strings.Split(flagName, "=")[0]
	} else if i+1 < len(args) {
		flagValue = args[i+1]
	}

	// Check if arg is a flag
	if !strings.HasPrefix(flagName, "--") && !strings.HasPrefix(flagName, "-") {
		return false
	}

	// Check if arg is a flag name
	if strings.HasPrefix(flagName, "--") {
		if err := c.parseFlagByName(flagName, flagValue); err != nil {
			panic(err)
		}
		return true
	}
	// Check if arg is a flag alias
	if strings.HasPrefix(flagName, "-") && len(flagName) == 2 {
		if err := c.parseFlagByAlias(flagName, flagValue); err != nil {
			panic(err)
		}
		return true
	}
	// Check if arg is a flag group
	if strings.HasPrefix(flagName, "-") && len(flagName) > 2 {
		if err := c.parseFlagGroup(flagName); err != nil {
			panic(err)
		}
		return true
	}
	return false
}

func (c *Command) parseFlagByName(flagName string, flagValue string) error {
	// Remove the -- prefix if it exists
	flagName = strings.TrimPrefix(flagName, "--")

	// Check if flag exists
	if flag := c.Flag(flagName); flag != nil {
		// Set flag value
		if flagValue != "" {
			v, err := parseFlagValue(flagValue, flag.dataType)
			if err != nil {
				return err
			}
			flag.Is = true
			flag.Value = v
			flag.set(v)
		} else {
			flag.Is = true
			flag.Value = true
			flag.set(true)
		}
	} else {
		return errors.New("Flag " + flagName + " does not exist")
	}
	return nil
}

func (c *Command) parseFlagByAlias(flagAlias string, flagValue string) error {
	// Remove the - prefix if it exists
	flagAlias = strings.TrimPrefix(flagAlias, "-")

	// Check if flag exists
	if flag := c.FlagByAlias(flagAlias); flag != nil {
		// Set flag value
		if flagValue != "" {
			v, err := parseFlagValue(flagValue, flag.dataType)
			if err != nil {
				return err
			}
			flag.Is = true
			flag.Value = v
			flag.set(v)
		} else {
			flag.Is = true
			flag.Value = true
			flag.set(true)
		}
	} else {
		return errors.New("Flag " + flagAlias + " does not exist")
	}
	return nil
}

func (c *Command) parseFlagGroup(flagGroup string) error {
	// Remove the - prefix
	flagGroup = flagGroup[1:]

	flags := strings.Split(flagGroup, "")
	for _, flag := range flags {
		if err := c.parseFlagByAlias(flag, "true"); err != nil {
			return err
		}
	}
	return nil
}
