package logger

import (
	"fmt"
	"sync"
)

// Logger provides a simple logging interface for the WebAuthn library
type Logger struct {
	enabled bool
	mu      sync.RWMutex
}

// globalLogger is the default logger instance used when not explicitly setting one
var globalLogger = &Logger{
	enabled: false,
}

// Enable turns on logging
func Enable() {
	globalLogger.Enable()
}

// Disable turns off logging
func Disable() {
	globalLogger.Disable()
}

// IsEnabled returns if logging is enabled
func IsEnabled() bool {
	return globalLogger.IsEnabled()
}

// Debug logs a debug message if logging is enabled
func Debug(component string, message string, values ...interface{}) {
	globalLogger.Debug(component, message, values...)
}

// Enable turns on logging for this logger instance
func (l *Logger) Enable() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = true
}

// Disable turns off logging for this logger instance
func (l *Logger) Disable() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.enabled = false
}

// IsEnabled returns if logging is enabled for this logger instance
func (l *Logger) IsEnabled() bool {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.enabled
}

// Debug logs a debug message if logging is enabled for this logger instance
func (l *Logger) Debug(component string, message string, values ...interface{}) {
	if !l.IsEnabled() {
		return
	}

	fmt.Printf("🟡 WebAuthn Debug 🟡 %s 🟡 %s: ", component, message)
	fmt.Println()
	for i := 0; i < len(values); i += 2 {
		// Log current element
		fmt.Printf("🟡%+v: ", values[i])

		// If there's a next element, log on the same line
		if i+1 < len(values) {
			fmt.Printf("%+v", values[i+1])
		}

		// New line only after loggin the pair (key, value)
		fmt.Println()
	}
}
