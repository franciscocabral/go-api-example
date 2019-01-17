package settings

import (
	"os"

	"fmt"

	"github.com/namsral/flag"
)

// EnvPrefix - Prefix for environment variables
const EnvPrefix = "API_EXAMPLE"

var config *Application

func init() {
	config = load(EnvPrefix)
}

func load(prefix string) *Application {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], prefix, 0)

	c := &Application{
		Logging: new(Logging),
	}

	c.LoadEnv(fs)
	c.Logging.LoadEnv(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(fmt.Sprintf("Error on parsing variables: %s", err.Error()))
	}

	return c
}

func Get() *Application {
	return config
}
