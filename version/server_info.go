// Copyright 2018 The QOS Authors

package version

var (
	// Version is the current version of Explorer
	Version = "0.2.0"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
