package vo

type InstallationVO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewInstallationVO() *InstallationVO {
	return &InstallationVO{}
}
