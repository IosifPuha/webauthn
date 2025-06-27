package protocol

import (
	"fmt"

	"github.com/go-webauthn/webauthn/internal/logger"
)

// LogDebug is a helper function to log debug messages if debug mode is enabled.
// This avoids import cycles by directly using the internal logger package.
func LogDebug(component string, message string, values ...interface{}) {
	if logger.IsEnabled() {
		logger.Debug(component, message, values...)
	}
}

// LogValue is a helper that simply prints a debug value with a clear prefix.
// Use this only for temporary debugging, not for production code.
func LogValue(component string, message string, values ...interface{}) {
	fmt.Printf("🟡 %s 🟡 %s: ", component, message)
	for _, value := range values {
		fmt.Printf("%+v ", value)
	}
	fmt.Println()
}
