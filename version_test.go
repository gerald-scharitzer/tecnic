package tecnic

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestVersion(t *testing.T) {
	want := "0.1.0-alpha.0"
	version := Version()
	assert.Equal(t, want, version.String())
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		want    string
		version SemVer
	}{
		// Test case 1: Version with no pre-release
		{want: "1.2.3", version: SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: ""}},

		// Test case 2: Version with pre-release
		{want: "4.5.6-alpha", version: SemVer{Major: 4, Minor: 5, Patch: 6, PreRelease: "alpha"}},

		// Test case 3: Version with build
		{want: "7.8.9+build", version: SemVer{Major: 7, Minor: 8, Patch: 9, Build: "build"}},

		// Test case 4: Version with pre-release and build
		{want: "10.11.12-alpha+build", version: SemVer{Major: 10, Minor: 11, Patch: 12, PreRelease: "alpha", Build: "build"}},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, test.version.String())
	}
}
