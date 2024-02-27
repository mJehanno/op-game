package version

import "github.com/blang/semver"

type VersionManager struct {
	current string
}

func NewVersionManager(current semver.Version) *VersionManager {
	return &VersionManager{
		current: current.String(),
	}
}

func (v VersionManager) GetCurrentVersion() string {
	return v.current
}
