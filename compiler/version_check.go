//go:build go1.16 && !go1.19
// +build go1.16,!go1.19

package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Version is the GopherJS compiler version string.
var Version = runtime.Version()

// GoVersion is the current Go 1.x version that GopherJS is compatible with.
var GoVersion = runtime.Version()[:6]

// CheckGoVersion checks the version of the Go distribution
// at goroot, and reports an error if it's not compatible
// with this version of the GopherJS compiler.
func CheckGoVersion(goroot string) error {
	if nvc, err := strconv.ParseBool(os.Getenv("GOPHERJS_SKIP_VERSION_CHECK")); err == nil && nvc {
		return nil
	}
	v, err := goRootVersion(goroot)
	if err != nil {
		return fmt.Errorf("unable to detect Go version for %q: %w", goroot, err)
	}
	if !strings.HasPrefix(v, GoVersion) {
		return fmt.Errorf("GopherJS %s requires a %v distribution, but found version %s", Version, GoVersion, v)
	}
	return nil
}

// goRootVersion defermines Go release for the given GOROOT installation.
func goRootVersion(goroot string) (string, error) {
	v, err := os.ReadFile(filepath.Join(goroot, "VERSION"))
	if err == nil {
		// Standard Go distribution has VERSION file inside its GOROOT, checking it
		// is the most efficient option.
		return string(v), nil
	}

	// Fall back to the "go version" command.
	cmd := exec.Command(filepath.Join(goroot, "bin", "go"), "version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("`go version` command failed: %w", err)
	}
	// Expected output: go version go1.18.1 linux/amd64
	parts := strings.Split(string(out), " ")
	if len(parts) != 4 {
		return "", fmt.Errorf("unexpected `go version` output %q, expected 4 words", string(out))
	}
	return parts[2], nil
}

// GoRelease does a best-effort to identify Go release we are building with.
// If unable to determin the precise version for the given GOROOT, falls back
// to the best guess available.
func GoRelease(goroot string) string {
	v, err := goRootVersion(goroot)
	if err == nil {
		// Prefer using the actual version of the GOROOT we are working with.
		return v
	}

	// If everything else fails, return just the Go version without patch level.
	return runtime.Version()
}
