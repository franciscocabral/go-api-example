package settings

import (
	"os"

	"github.com/namsral/flag"
)

// Application is the structure to configure the application
type Application struct {
	Name        string
	HostName    string
	Version     string
	CompanyName string
	Logging     *Logging
}

// LoadEnv - load environment variables
func (c *Application) LoadEnv(fs *flag.FlagSet) {
	c.HostName, _ = os.Hostname()

	fs.StringVar(&c.Name, "name", "Example-API", "Application Name")
	fs.StringVar(&c.Version, "version", "", "Application Version")
	fs.StringVar(&c.CompanyName, "company_name", "Barba's Co.", "Application Company Name")
}
