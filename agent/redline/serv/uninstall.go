package serv

import (
	"github.com/threatwinds/logger"
	"github.com/0days-ru/UTMStack/agent/redline/configuration"
	"github.com/0days-ru/UTMStack/agent/redline/utils"
)

func UninstallService(h *logger.Logger) {
	// Uninstall service
	err := utils.StopService(configuration.SERV_NAME)
	if err != nil {
		h.Fatal("error stopping %s: %v", configuration.SERV_NAME, err)
	}
	err = utils.UninstallService("%s")
	if err != nil {
		h.Fatal("error uninstalling %s: %v", configuration.SERV_NAME, err)
	}
}
