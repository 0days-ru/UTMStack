package serv

import (
	"github.com/kardianos/service"
	"github.com/0days-ru/UTMStack/agent/updater/configuration"
)

// GetConfigServ creates and returns a pointer to a service configuration structure.
func GetConfigServ() *service.Config {
	svcConfig := &service.Config{
		Name:        configuration.SERV_NAME,
		DisplayName: "UTMStack Updater",
		Description: "UTMStack Updater Service",
	}

	return svcConfig
}
