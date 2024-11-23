package version

import (
	"context"
	"sync"
	"time"

	"golang.org/x/mod/semver"
)

var latest struct {
	sync.Once
	vers string
}

// IsLatest checks if the current version is the latest one.
func IsLatest() bool {
	return semver.Compare(Version(), Latest()) >= 0
}

// Latest returns the latest version number discovered via one or more sources.
// If it cannot detect the latest version, it will return the current one.
func Latest() string {
	latest.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		vers, err := latestGithub(ctx)
		if err == nil {
			Log.Info("latest on github", "version", vers)
			latest.vers = vers
			return
		}
		Log.Warn("cannot check version", "err", err)
		// give up
		latest.vers = Version()
	})
	return latest.vers
}

func semverIsValid(vers string) bool {
	return semver.IsValid(vers)
}

func semverIsDev(vers string) bool {
	return !semverIsValid(vers) || semver.Prerelease(vers) != ""
}

func semverLatestFromList(names []string) string {
	var vers []string
	for _, tag := range names {
		if !semver.IsValid(tag) {
			continue
		}
		vers = append(vers, tag)
	}
	if len(vers) == 0 {
		return ""
	}
	semver.Sort(vers)
	return vers[len(vers)-1]
}
