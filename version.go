package tecnic

import (
	"strconv"
	"strings"
)

// Semantic version https://semver.org/
type SemVer struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
}

// The semantic version of this module
func Version() SemVer {
	return SemVer{Major: 0, Minor: 1, Patch: 0, PreRelease: "", Build: ""}
}

// Implements the Stringer interface
func (version *SemVer) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(version.Major))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Minor))
	builder.WriteRune('.')
	builder.WriteString(strconv.Itoa(version.Patch))
	if len(version.PreRelease) > 0 {
		builder.WriteRune('-')
		builder.WriteString(version.PreRelease)
	}
	if len(version.Build) > 0 {
		builder.WriteRune('+')
		builder.WriteString(version.Build)
	}
	return builder.String()
}
