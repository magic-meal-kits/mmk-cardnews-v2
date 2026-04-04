#!/bin/bash
set -e

# Build Go CLI if not already built
if [ ! -f "$CLAUDE_PROJECT_DIR/dist/mmk-cn" ]; then
  echo "Building mmk-cn CLI..."
  cd "$CLAUDE_PROJECT_DIR"
  make build
  echo "✓ mmk-cn built"
else
  echo "✓ mmk-cn already exists"
fi

# Check API keys
if [ -n "$GEMINI_API_KEY" ]; then
  echo "✓ GEMINI_API_KEY is set"
else
  echo "⚠ GEMINI_API_KEY is not set (AI image generation will be unavailable)"
fi

exit 0
