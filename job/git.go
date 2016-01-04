package job

import (
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

// PullRepo will pull the last update of the repo
func PullRepo() {
	log.Info("Reset of the repo started")
	exec.Command("git", "-C", "/etc/data", "fetch", "origin").Run()
	exec.Command("git", "-C", "/etc/data", "reset", "--hard", "origin/master").Run()
	log.Info("Reset of the repo finished")
}
