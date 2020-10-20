// +build js
package doc

import "testing"

func Test(t *testing.T) {
	t.Skip("runtime error: native function not implemented: internal/poll.fcntl")
}
