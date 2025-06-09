#!/bin/bash

echo "Preparing the build environment..."
cd "$(dirname "$0")/.."

OS=$(uname)

case "$OS" in
  "Darwin")
    echo "Detected macOS. Building binaries..."
    go build -o build/macOS/event-booking
    ;;
  "Linux")
    echo "Detected Linux. Building binaries..."
    go build -o build/linux/event-booking
    ;;
  "MINGW"*|"MSYS"*|"CYGWIN"*)
    echo "Detected Windows. Building binaries..."
    go build -o build/windows/event-booking.exe
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac
