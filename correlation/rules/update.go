package rules

import (
	"log"
	"os/exec"
	"time"

	"github.com/0days-ru/UTMStack/correlation/utils"
)

func Update(updateReady chan bool) {
	first := true
	cnf := utils.GetConfig()
	if cnf.UseSystemRules == "true" {
		for {
			log.Println("Downloading rules")
			
				cnf := utils.GetConfig()

				rm := exec.Command("rm", "-R", cnf.RulesFolder+"system")
				_ = rm.Run()
				
				clone := exec.Command("git", "clone", "https://github.com/utmstack/rules.git", cnf.RulesFolder+"system")
				_ = clone.Run()

				if first {
					first = false
					updateReady <- true
				}
				
				log.Println("Rules updated")
			time.Sleep(48 * time.Hour)
		}
	} else {
		log.Println("System rules update disabled")
		updateReady <- true
	}
}

