package semver

import (
	"regexp"
	"strconv"
)

const pattern = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

type SemanticVersion struct {
	matched       bool
	versionString string

	Major         int
	Minor         int
	Patch         int
	PreRelease    string
	BuildMetadata string
}

// Matched returns true if the version string was successfully parsed
func (s SemanticVersion) Matched() bool {
	return s.matched
}

// String returns the original version string if it was successfully parsed
func (s SemanticVersion) String() string {
	if s.matched {
		return s.versionString
	}

	return ""
}

// ParseVersion parses the version string and populates the struct if a match is found.
func (s *SemanticVersion) ParseVersion(r *regexp.Regexp) {
	if match := r.FindStringSubmatch(s.versionString); len(match) > 0 {
		s.matched = true

		s.Major, _ = strconv.Atoi(match[r.SubexpIndex("major")])
		s.Minor, _ = strconv.Atoi(match[r.SubexpIndex("minor")])
		s.Patch, _ = strconv.Atoi(match[r.SubexpIndex("patch")])
		s.PreRelease = match[r.SubexpIndex("prerelease")]
		s.BuildMetadata = match[r.SubexpIndex("buildmetadata")]

	}
}

// New creates a new SemanticVersion from a string
func FromString(version string) SemanticVersion {
	return SemanticVersion{versionString: version}
}
