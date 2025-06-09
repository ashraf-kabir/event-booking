#!/bin/bash

echo "Preparing to serve the application..."
cd "$(dirname "$0")/.."

OS=$(uname)

case "$OS" in
  "Darwin")
    echo "Detected macOS. Running ./build/macOS/event-booking..."
    ./build/macOS/event-booking
    ;;
  "Linux")
    echo "Detected Linux. Running ./build/linux/event-booking..."
    ./build/linux/event-booking
    ;;
  "MINGW"*|"MSYS"*|"CYGWIN"*)
    echo "Detected Windows. Running ./build/windows/event-booking.exe..."
    ./build/windows/event-booking.exe
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac
