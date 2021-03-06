package shellcligen

import "fmt"

// CLIProgram ...
type CLIProgram struct {
	Help      string      `json:"message" yaml:"help_message"`
	Options   []CLIOption `json:"options" yaml:"options"`
	SafeFlags bool        `json:"safe_flags" yaml:"safe_flags"`
}

// Name ...
type Name struct {
	Short, Long string
}

// CLIOption ...
type CLIOption struct {
	// LongName ...
	LongName string `json:"long_name" yaml:"long_name"`

	// ShortName ...
	ShortName string `json:"short_name" yaml:"short_name"`

	// Required ...
	Required bool `json:"required" yaml:"required"`

	// ArgsRequired ...
	ArgsRequired bool `json:"args_required" yaml:"args_required"`

	// ConflictsWith ...
	ConflictsWith []string `json:"conflicts_with" yaml:"conflicts_with"`

	// Help ...
	Help bool `json:"is_help" yaml:"is_help"`
}

func (cliopt CLIOption) String() string {
	return fmt.Sprintf("Long name: `%s`, Short name: `%s`, Required: %t",
		cliopt.LongName, cliopt.ShortName, cliopt.Required)
}
