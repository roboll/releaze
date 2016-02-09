package git

import (
	"encoding/json"

	"github.com/roboll/releaze/pkg/scm"
)

var (
	commit string
	branch string
)

type info struct{}

func (*info) Commit() string {
	return commit
}

func (*info) Branch() string {
	return branch
}

func Info() scm.Info {
	return &info{}
}

func (git *info) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Commit string `json:"commit"`
		Branch string `json:"branch"`
	}{
		Commit: git.Commit(),
		Branch: git.Branch(),
	})
}
