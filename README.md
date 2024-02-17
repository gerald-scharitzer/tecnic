# Tecnic ⚙️

provides technical building blocks for applications based on the [Go programming language](https://go.dev/).

Its package structure is based on the [Go standard library](https://pkg.go.dev/std).

[![Go Reference](https://pkg.go.dev/badge/gopkg.in/gerald-scharitzer/tecnic.v0.svg)](https://pkg.go.dev/gopkg.in/gerald-scharitzer/tecnic.v0)

# Features ✨

Get pointers to variables derived from constants and literals.

# Use 🔌

Use the library by importing the module's packages with

```go
import (
    tp "gopkg.in/gerald-scharitzer/tecnic.v0/pkg"
)
```

where `pkg` is the package path relative to the module.

# Develop 🚀

with Git and Go.

## Cycle

1. Get with `git clone https://github.com/gerald-scharitzer/tecnic.git`
2. Enter with `cd tecnic`
3. Increase version in [version.go#Version](version.go#Version) and [version_test.go#TestVersion](version_test.go#TestVersion)
4. Test with `go test ./...`
5. Check with `go vet ./...`
6. Build with `go build ./...`
7. Tag with `git tag semver`
8. Push with `git push origin semver`
9. Publish with `GOPROXY=proxy.golang.org go list -m gopkg.in/gerald-scharitzer/tecnic.vn@semver`

where `semver` is `v` followed by the [semantic version](https://semver.org/spec/v2.0.0.html) (e.g. `v0.0.0`)
and `vn` is `v` followed by the major version number (e.g. `v0`)

## Rules

- Use `fmt.Println` instead of `println`, because [`println`](https://pkg.go.dev/builtin@go1.21.6#println)
is implementation-specific and not guaranteed to stay in the language.
- Use `fmt.Print` instead of `print`, because [`print`](https://pkg.go.dev/builtin@go1.21.6#print)
is implementation-specific and not guaranteed to stay in the language.
