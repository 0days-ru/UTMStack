package configuration

import (
	"path/filepath"

	"github.com/0days-ru/UTMStack/agent/runner/utils"
)

type ServicesBin struct {
	AgentServiceBin   string
	RedlineServiceBin string
	UpdaterServiceBin string
}

const (
	MASTERVERSIONENDPOINT = "/management/info"
	INSTALLER_LOG_FILE    = "utmstack_agent_installer.log"
	Bucket                = "https://cdn.utmstack.com/agent_updates/"
	AgentManagerPort      = "9000"
	LogAuthProxyPort      = "50051"
)

func GetCertPath() string {
	path, _ := utils.GetMyPath()
	return filepath.Join(path, "certs", "utm.crt")
}

func GetKeyPath() string {
	path, _ := utils.GetMyPath()
	return filepath.Join(path, "certs", "utm.key")
}

func GetCaPath() string {
	path, _ := utils.GetMyPath()
	return filepath.Join(path, "certs", "ca.crt")
}
