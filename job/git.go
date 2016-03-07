package job

import (
	"os"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

const (
	repoPath = "/etc/data"
)

// UpdateSource will do the job of creating the git repo
// if needed and then make sure that the last version
// is used
func UpdateSource() bool {
	log.Info("Reset of the repo started")
	_, err := os.Lstat(repoPath)
	if err != nil {
		cloneRepo(repoPath)
	}
	needDBUpdate := pullRepo(repoPath)
	log.Info("Reset of the repo finished")
	return needDBUpdate
}

func cloneRepo(path string) {
	exec.Command("git", "clone", "https://github.com/helphone/data.git", "/etc/data").Run()
}

func pullRepo(path string) (needDBUpdate bool) {
	exec.Command("git", "-C", path, "fetch", "origin").Run()

	currentCommit, _ := exec.Command("git", "-C", path, "rev-parse", "@").Output()
	remoteCommit, _ := exec.Command("git", "-C", path, "rev-parse", "@{u}").Output()

	if string(currentCommit) != string(remoteCommit) {
		exec.Command("git", "-C", path, "reset", "--hard", "origin/master").Run()
		return true
	}
	return false
}
