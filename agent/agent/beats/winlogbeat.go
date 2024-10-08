package beats

import (
	"fmt"
	"path/filepath"

	"github.com/threatwinds/logger"
	"github.com/threatwinds/validations"
	"github.com/0days-ru/UTMStack/agent/agent/configuration"
	"github.com/0days-ru/UTMStack/agent/agent/logservice"
	"github.com/0days-ru/UTMStack/agent/agent/utils"
)

type Winlogbeat struct{}

func (w Winlogbeat) Install() error {
	path, err := utils.GetMyPath()
	if err != nil {
		return fmt.Errorf("error getting current path: %v", err)
	}

	winlogPath := filepath.Join(path, "beats", "winlogbeat")
	beatConfig := BeatConfig{
		LogsPath:    filepath.Join(winlogPath, "logs"),
		LogFileName: "windowscollector",
	}

	if isInstalled, err := utils.CheckIfServiceIsInstalled(configuration.WinServName); err != nil {
		return fmt.Errorf("error checking if %s service is installed: %v", configuration.WinServName, err)
	} else if !isInstalled {
		err = utils.CreatePathIfNotExist(beatConfig.LogsPath)
		if err != nil {
			return fmt.Errorf("error creating %s folder", beatConfig.LogsPath)
		}

		configFile := filepath.Join(winlogPath, "winlogbeat.yml")
		templateFile := filepath.Join(path, "templates", "winlogbeat.yml")
		err = utils.GenerateFromTemplate(beatConfig, templateFile, configFile)
		if err != nil {
			return fmt.Errorf("error configuration from %s: %v", templateFile, err)
		}

		err = utils.Execute("sc",
			winlogPath,
			"create",
			configuration.WinServName,
			"binPath=",
			fmt.Sprintf("\"%s\\winlogbeat.exe\" --environment=windows_service -c \"%s\\winlogbeat.yml\" --path.home \"%s\" --path.data \"C:\\ProgramData\\winlogbeat\" --path.logs \"C:\\ProgramData\\winlogbeat\\logs\" -E logging.files.redirect_stderr=true", winlogPath, winlogPath, winlogPath),
			"DisplayName=",
			configuration.WinServName,
			"start=",
			"auto")
		if err != nil {
			return fmt.Errorf("error installing %s service: %s", configuration.WinServName, err)
		}

		err = utils.Execute("sc", winlogPath, "start", configuration.WinServName)
		if err != nil {
			return fmt.Errorf("error starting %s service: %s", configuration.WinServName, err)
		}
	}

	return nil
}

func (w Winlogbeat) SendSystemLogs(h *logger.Logger) {
	logLinesChan := make(chan []string)
	path, err := utils.GetMyPath()
	if err != nil {
		h.ErrorF("error getting current path: %v", err)
	}
	winbLogPath := filepath.Join(path, "beats", "winlogbeat", "logs")

	go utils.WatchFolder("windowscollector", winbLogPath, logLinesChan, configuration.BatchCapacity, h)
	for logLine := range logLinesChan {
		validatedLogs := []string{}
		for _, log := range logLine {
			validatedLog, _, err := validations.ValidateString(log, false)
			if err != nil {
				h.ErrorF("error validating log: %s: %v", log, err)
				continue
			}
			validatedLogs = append(validatedLogs, validatedLog)
		}
		logservice.LogQueue <- logservice.LogPipe{
			Src:  string(configuration.LogTypeWindowsAgent),
			Logs: validatedLogs,
		}
	}
}

func (w Winlogbeat) Uninstall() error {
	if isInstalled, err := utils.CheckIfServiceIsInstalled(configuration.WinServName); err != nil {
		return fmt.Errorf("error checking if %s is running: %v", configuration.WinServName, err)
	} else if isInstalled {
		err = utils.StopService(configuration.WinServName)
		if err != nil {
			return fmt.Errorf("error stopping %s: %v", configuration.WinServName, err)
		}
		err = utils.UninstallService(configuration.WinServName)
		if err != nil {
			return fmt.Errorf("error uninstalling %s: %v", configuration.WinServName, err)
		}
	}

	return nil
}
