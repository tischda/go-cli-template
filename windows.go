//go:build windows

package main

//go:generate cmd /c go mod init github.com/%OWNER%/%PROJECT%
//go:generate cmd /c go run template.go --owner=%OWNER% --project=%PROJECT% --copyright="%COPYRIGHT%"
