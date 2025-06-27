package webauthn

import (
	"github.com/go-webauthn/webauthn/internal/logger"
)

// DebugLog logs a message if debugging is enabled in the WebAuthn instance
func (webauthn *WebAuthn) DebugLog(component string, message string, values ...interface{}) {
	if webauthn.Config.Debug {
		logger.Debug(component, message, values...)
	}
}
