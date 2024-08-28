package processor

import (
	"github.com/0days-ru/UTMStack/office365/utils"
	"github.com/utmstack/config-client-go/types"
)

func PullLogs(startTime string, endTime string, group types.ModuleGroup) {
	utils.Logger.Info("starting log sync for : %s from %s to %s", group.GroupName, startTime, endTime)

	agent := GetOfficeProcessor(group)

	err := agent.GetAuth()
	if err != nil {
		utils.Logger.ErrorF("error getting auth token: %v", err)
		return
	}

	err = agent.StartSubscriptions()
	if err != nil {
		utils.Logger.ErrorF("error starting subscriptions: %v", err)
		return
	}

	agent.GetLogs(startTime, endTime, group)
}
