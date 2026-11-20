//go:build !unix

package agentbuilder

// oNoFollow is 0 where the platform has no O_NOFOLLOW; writeProjectFiles checks for a symlink
// with Lstat instead.
const oNoFollow = 0

func isSymlinkError(error) bool { return false }
