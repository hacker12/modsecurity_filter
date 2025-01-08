#!/bin/bash

VERSION="2.0.0"

echo "Compiling for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o modsecurity_filter_windows_amd64.exe modsecurity_filter.go

echo "Compiling for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o modsecurity_filter_linux_amd64 modsecurity_filter.go

echo "Compiling for Linux (arm)..."
GOOS=linux GOARCH=arm go build -ldflags "-X main.Version=$VERSION" -o modsecurity_filter_linux_arm modsecurity_filter.go

echo "Compilation completed!"
