package shellcligen

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrOpeningInputFile        = errors.New("error opening input file")
	ErrReadingInputFile        = errors.New("error reading input file")
	ErrParsingInputFile        = errors.New("error parsing input file")
	ErrMissingRequiredArgument = errors.New("error missing required argument")
)

// ParseCLIProgram ...
func ParseCLIProgram(filename string) (CLIProgram, error) {
	file, err := os.Open(filename)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error opening input file: %w", ErrOpeningInputFile)
	}

	defer file.Close()

	cli := CLIProgram{}

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error reading input file: %w", ErrReadingInputFile)
	}

	err = yaml.Unmarshal(fileContent, &cli)
	if err != nil {
		return CLIProgram{}, fmt.Errorf("error parsing input file: %w", err)
	}

	return cli, nil
}

func hasRequiredOptions(cliProgram *CLIProgram) bool {
	required := false

	for _, option := range cliProgram.Options {
		if option.Required {
			required = true
			break
		}
	}

	return required
}
