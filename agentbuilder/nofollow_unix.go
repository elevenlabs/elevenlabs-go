//go:build unix

package agentbuilder

import (
	"errors"
	"syscall"
)

// oNoFollow makes opening a symlink fail, so a file is never written through one.
const oNoFollow = syscall.O_NOFOLLOW

func isSymlinkError(err error) bool { return errors.Is(err, syscall.ELOOP) }
