package utils

import (
	"fmt"
)

// Log, useful for display info.
func Log(format string, args ...any) {
	fmt.Println(fmt.Sprintf(format, args...))
}

// Fatal, useful for display error.
func Fatal(format string, args ...any) {
	fmt.Println(fmt.Sprintf(format, args...))
}
