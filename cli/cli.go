package cli

import "os"

type Command interface {
	Execute() error
}

var (
	commands = map[string]func(args *Args) Command{
		"create": newCreateCommand,
	}
)

// GetCommandStub checks if a registered command has been passed via the CLI
// and returns it to the caller where the command can be executed
func GetCommandStub() Command {
	args := &Args{os.Args}

	// remove executable name
	args.pop()
	if args.len() == 0 { // no command line arguments passed
		return nil
	}

	mainCommand := args.pop()
	if command, ok := commands[mainCommand]; ok {
		return command(args)
	}

	return nil
}
