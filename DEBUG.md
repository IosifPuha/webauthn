# Debugging in go-webauthn

This document explains how to use the built-in logging capabilities of the go-webauthn library to debug WebAuthn flows.

## Enabling Debug Logging

Debug logging can be enabled by setting the `Debug` flag to `true` in your WebAuthn configuration:

```go
config := &webauthn.Config{
    RPDisplayName: "Example RP",
    RPID:         "example.com",
    RPOrigins:    []string{"https://example.com"},
    
    // Enable debug logging
    Debug:        true,
}

// Initialize WebAuthn
_, err := webauthn.New(config)
```

## What Gets Logged

When debug logging is enabled, the library will automatically log:

1. Method calls with their parameters
2. Key verification steps
3. Data transformations
4. Authentication and attestation details

## Logging in Protocol Package

For components in the protocol package, use the `LogDebug` helper function:

```go
import "github.com/go-webauthn/webauthn/protocol"

// Log debugging information
protocol.LogDebug("component.name", "Message description", 
    "paramName1", paramValue1,
    "paramName2", paramValue2)
```

## Logging in Webauthn Package

For components in the webauthn package, use the `DebugLog` method on the WebAuthn instance:

```go
// Using the WebAuthn instance directly
webauthn.DebugLog("component.name", "Message description", 
    "paramName1", paramValue1,
    "paramName2", paramValue2)
```

## Temporary Debugging

For quick troubleshooting outside the regular logging system, you can use `LogValue`:

```go
protocol.LogValue("quickcheck", "Examining value", myValue)
```

This is intended for temporary debugging only and prints regardless of the Debug setting.
