package service

import (
	"github.com/joho/godotenv"
	"github.com/peake100/stalkforecaster-go/service/servers"
)

func StartServers() *servers.ServersMonitor {
	// Load .env file. We are only using .env files in development, so we can suppress
	// errors here -- it won't affect production.
	_ = godotenv.Load()

	monitor := servers.NewServiceMonitor()
	monitor.StartServers()
	return monitor
}
