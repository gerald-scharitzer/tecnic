// provides technical building blocks for applications based on the [Go programming language](https://go.dev/)
// with a package structure based on the [Go standard library](https://pkg.go.dev/std)
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
}

func Version() SemVer {
	return SemVer{Major: 0, Minor: 0, Patch: 0, PreRelease: ""}
}

func VersionString() string {
	version := Version()
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
	return builder.String()
}
