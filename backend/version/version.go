package version

type VersionManager struct {
	current string
}

func NewVersionManager() *VersionManager {
	return &VersionManager{
		current: "1.0.0",
	}
}

func (v VersionManager) GetCurrentVersion() string {
	return v.current
}
