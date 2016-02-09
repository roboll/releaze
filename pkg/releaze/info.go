package releaze

import (
	"encoding/json"

	"github.com/roboll/releaze/pkg/scm"
	"github.com/roboll/releaze/pkg/scm/git"
)

type ReleaseInfo interface {
	Version() string

	Scm() scm.Info
}

var (
	version string
)

type info struct{}

func (*info) Version() string {
	return version
}

func (*info) Scm() scm.Info {
	//TODO detect different scm's with hint arg or presence of values
	return git.Info()
}

func (info *info) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		V   string   `json:"version"`
		Scm scm.Info `json:"scm"`
	}{
		V:   info.Version(),
		Scm: info.Scm(),
	})
}

func Get() ReleaseInfo {
	return &info{}
}
