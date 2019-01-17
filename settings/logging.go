package settings

import (
	"github.com/namsral/flag"
)

// Logging is the structure to configure the logger
type Logging struct {
	Level string
}

// LoadEnv load log variables from environment variables
func (c *Logging) LoadEnv(fs *flag.FlagSet) {
	fs.StringVar(&c.Level, "logging_level", "error", "Minimal level to log")
}
