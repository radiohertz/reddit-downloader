package cli

import (
	"log"
	"strconv"

	"github.com/axiiomatic/reddit-downloader/downloader"
)

// Command takes a name and datatype for the command
type Command struct {
	Name  string
	Value string
}

const LIMIT = 25

// Cli initializes the arguments required for the app
type Cli struct {
	Commands []Command
}

// ParseArgs takes in os.Args and parses the system args
func (r *Cli) ParseArgs(args []string) []Command {
	requiredArgs := args[1:]

	if len(requiredArgs) == 0 {
		log.Fatal("No arguments were provided")
	}

	for i, v := range requiredArgs {
		if i%2 == 1 {
			continue
		}

		if i == 0 && v != "--subr" {
			log.Fatal("First argument should be --subr _subredditname_")
		}

		r.Commands = append(r.Commands, Command{
			Name:  v,
			Value: requiredArgs[i+1],
		})
	}

	return r.Commands

}

// Init takes os.args and initalizes the cli
func (r *Cli) Init(args []string) {
	values := r.ParseArgs(args)

	limit := LIMIT

	if len(values) > 1 {
		l, _ := strconv.ParseInt(values[1].Value, 10, 32)

		if l < 0 || l > 100 {
			log.Println("Limit cannot be less than 0 or greater than 100.Defaulting to 25")
			limit = LIMIT
		} else {
			limit = int(l)
		}

	}

	subr := values[0].Value

	downloader.MakeRequestForReddit(subr, limit)

}
