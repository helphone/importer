package job

import (
	"io/ioutil"
	"testing"
)

func TestCloneAndPullRepo(t *testing.T) {
	path, err := ioutil.TempDir("", "git2go")
	if err != nil {
		t.Errorf("Error during the getting of temp dir, err: %v", err)
	}

	cloneRepo(path)
	pullRepo(path)
}
