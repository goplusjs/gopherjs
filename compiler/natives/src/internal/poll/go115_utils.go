// +build js
// +build go1.15

package poll

import (
	"errors"
	"syscall"
)

var (
	ErrTimeout = errors.New("timeout")
)

func fcntl(fd int, cmd int, arg int) (int, error) {
	r, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg))
	if e != 0 {
		return int(r), syscall.Errno(e)
	}
	return int(r), nil
}
