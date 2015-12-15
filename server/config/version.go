package config

var versionMajor = "0.1"
var versionMinor string

func Version() string {
	return versionMajor + versionMinor
}
