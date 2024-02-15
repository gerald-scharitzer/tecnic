package tecnic

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestVersionString(t *testing.T) {
	// Test case 1: Version with no pre-release
	want1 := "1.2.3"
	version1 := SemVer{Major: 1, Minor: 2, Patch: 3, PreRelease: ""}
	assert.Equal(t, want1, version1.String())

	// Test case 2: Version with pre-release
	want2 := "4.5.6-alpha"
	version2 := SemVer{Major: 4, Minor: 5, Patch: 6, PreRelease: "alpha"}
	assert.Equal(t, want2, version2.String())
}
