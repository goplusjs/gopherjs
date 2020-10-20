// +build js

package time_test

import (
	"testing"
)

func TestSleep(t *testing.T) {
	t.Skip("time.Now() is not accurate enough for the test")
}

func TestEmbeddedTZData(t *testing.T) {
	t.Skip()
}

func TestEnvVarUsage(t *testing.T) {
	t.Skip()
}

func TestLoadLocationFromTZData(t *testing.T) {
	t.Skip()
}
