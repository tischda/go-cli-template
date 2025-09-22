//go:build !windows

package main

//go:generate sh -c 'go mod init github.com/${OWNER}/${PROJECT}'
//go:generate sh -c 'go run template.go --owner=${OWNER} --project=${PROJECT} --copyright="${COPYRIGHT}"'
