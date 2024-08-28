package main

import "github.com/0days-ru/UTMStack/installer/utils"

func InstallTools() error {
	env := []string{"DEBIAN_FRONTEND=noninteractive"}
	err := utils.RunEnvCmd(env, "apt-get", "install", "-y", "cockpit")
	if err != nil {
		return err
	}

	return nil
}
