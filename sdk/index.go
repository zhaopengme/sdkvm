package sdk

type Sdk interface {
	Versions()
	CurrentVersions()
	UseVersion(version string)
	Install(version string)
	Uninstall(version string)
}
