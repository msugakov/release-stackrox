package version

// version value is embedded at the build time, see Makefile.
var version string = "unknown"

func GetVersion() string {
	return version
}
