package semver

import (
	"regexp"
	"strconv"
)

var regex *regexp.Regexp

// https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
const pattern = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

func init() {
	regex = regexp.MustCompile(pattern)
}

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
func (s *SemanticVersion) ParseVersion() {
	if match := regex.FindStringSubmatch(s.versionString); len(match) > 0 {
		s.matched = true

		s.Major, _ = strconv.Atoi(match[regex.SubexpIndex("major")])
		s.Minor, _ = strconv.Atoi(match[regex.SubexpIndex("minor")])
		s.Patch, _ = strconv.Atoi(match[regex.SubexpIndex("patch")])
		s.PreRelease = match[regex.SubexpIndex("prerelease")]
		s.BuildMetadata = match[regex.SubexpIndex("buildmetadata")]

	}
}

// FromString creates a new SemanticVersion from a string
func FromString(version string) SemanticVersion {
	return SemanticVersion{versionString: version}
}
