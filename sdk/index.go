package sdk

type Sdk interface {
	Versions()
	LocalVersions()
	UseVersion(version string)
	Install(version string)
	Uninstall(version string)
	SetEnv()
}
