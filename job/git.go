package job

import (
	"os"

	log "github.com/Sirupsen/logrus"
	git "gopkg.in/libgit2/git2go.v23"
)

const (
	repoPath = "/etc/data"
)

// UpdateSource will do the job of creating the git repo
// if needed and then make sure that the last version
// is used
func UpdateSource() error {
	log.Info("Reset of the repo started")
	_, err := os.Lstat(repoPath)
	if err != nil {
		err = cloneRepo(repoPath)
		if err != nil {
			return err
		}
	}
	err = pullRepo(repoPath)
	log.Info("Reset of the repo finished")
	return err
}

func cloneRepo(path string) error {
	cloneOptions := &git.CloneOptions{
		Bare:           false,
		CheckoutBranch: "master",
	}
	_, err := git.Clone("https://github.com/helphone/data.git", path, cloneOptions)
	return err
}

func pullRepo(path string) (err error) {
	repo, err := git.OpenRepository(path)
	if err != nil {
		log.Errorf("Open repo fails, err: %v", err)
		return
	}

	remote, err := repo.Remotes.Lookup("origin")
	if err != nil {
		log.Errorf("Lookup remote fails, err: %v", err)
		return
	}

	err = remote.Fetch([]string{}, nil, "")
	if err != nil {
		log.Errorf("Cannot fetch the remote, err: %v", err)
		return
	}

	err = repo.SetHead("refs/remotes/origin/master")
	if err != nil {
		log.Errorf("Cannot set the HEAD on the remote, err: %v", err)
		return
	}

	repo.StateCleanup()
	return
}
