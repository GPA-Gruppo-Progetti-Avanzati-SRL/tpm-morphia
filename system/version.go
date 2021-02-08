package system

import "runtime"

type Info struct {
	TPMMorphia string
	BuildNum   string
	GitSHA     string

	GOOS   string
	GOARCH string
}

const DefaultTpmMorphiaVersion = "0.0.0"
const DefaultTpmMorphiaBuildNum = "UNSET"
const DefaultTpmMorphiaGitSHA = "UNSET"

const DefaultTpmMorphiaBuildNum_EnvVar = "DEVOPS_BUILDNUM"
const DefaultTpmMorphiaGitSHA_EnvVar = "DEVOPS_GITSHA"

var SysInfo = Info{
	TPMMorphia: DefaultTpmMorphiaVersion,
	BuildNum:   DefaultTpmMorphiaBuildNum,
	GitSHA:     DefaultTpmMorphiaGitSHA,
	GOOS:       runtime.GOOS,
	GOARCH:     runtime.GOARCH,
}
