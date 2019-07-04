package log

import "os"

// Fun is a logging function type
type Fun func(msg string, ctx ...interface{})

// IsTestLogSet returns true when TEST_LOG environment variable is set
func IsTestLogSet() bool {
	return os.Getenv("TEST_LOG") != ""
}
