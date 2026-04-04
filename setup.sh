#!/bin/bash
set -e

echo "=== mmk-cardnews-v2 setup ==="

# 1. Check Node.js
if ! command -v node &>/dev/null; then
  echo "ERROR: Node.js is required. Install from https://nodejs.org"
  exit 1
fi
echo "✓ Node.js $(node --version)"

# 2. Build Go CLI
if command -v go &>/dev/null; then
  echo "Building mmk-cn from source..."
  mkdir -p dist
  go build -o dist/mmk-cn ./cmd/mmk-cn
  echo "✓ Built dist/mmk-cn"
else
  echo "Go not found. Downloading prebuilt binary..."
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  ARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/')

  LATEST=$(curl -sI "https://github.com/pureugong/mmk-cardnews-v2/releases/latest" | grep -i location | sed 's/.*tag\///' | tr -d '\r\n')
  if [ -z "$LATEST" ]; then
    echo "ERROR: Could not determine latest release. Install Go and run: make build"
    exit 1
  fi

  mkdir -p dist
  URL="https://github.com/pureugong/mmk-cardnews-v2/releases/download/${LATEST}/mmk-cn-${OS}-${ARCH}"
  echo "Downloading ${URL}..."
  curl -sL "$URL" -o dist/mmk-cn
  chmod +x dist/mmk-cn
  echo "✓ Downloaded dist/mmk-cn"
fi

echo ""
echo "Setup complete! Usage:"
echo "  ./dist/mmk-cn new \"Your Topic\" --slides 7"
echo ""
echo "Optional: Set GEMINI_API_KEY for AI image generation"
echo "  cp .env.example .env && edit .env"
