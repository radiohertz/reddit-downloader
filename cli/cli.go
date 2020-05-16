package cli

import (
	"log"
)

// Command takes a name and datatype for the command
type Command struct {
	Name  string
	Value string
}

// Cli initializes the arguments required for the app
type Cli struct {
	Commands []Command
}

// ParseArgs takes in os.Args and parses the system args
func (r *Cli) ParseArgs(args []string) {
	requiredArgs := args[1:]

	if !(len(requiredArgs) > 0) {
		log.Fatal("No arguments were provided")
	}

}
