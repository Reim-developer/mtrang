package utils

import (
	"fmt"
)

// Log, useful for display info
func Log(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

// Fatal, useful for display error
func Fatal(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}
