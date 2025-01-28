#!/bin/bash

set -e

APP_NAME="tabman"
VERSION="v0.0.1"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH="amd64"

# # Determine the correct binary
# if [[ "$OS" == "darwin" ]]; then
#     FILE="${APP_NAME}-${VERSION}-macos-${ARCH}.zip"
# # elif [[ "$OS" == "linux" ]]; then
# #     FILE="${APP_NAME}-${VERSION}-linux-${ARCH}.tar.gz"
# else
#     echo "Unsupported OS: $OS"
#     exit 1
# fi

# Download the binary
echo "Downloading $APP_NAME..."

curl -LO https://github.com/abhishek71994/tab-manager-cli/releases/download/v0.0.1/tabman | bash

# Extract and install
# if [[ "$FILE" == *.zip ]]; then
#     unzip -o "$FILE"
# else
#     tar -xzf "$FILE"
# fi
chmod +x "$APP_NAME"
sudo mv "$APP_NAME" "$INSTALL_DIR"

# Clean up

echo "$APP_NAME installed successfully to $INSTALL_DIR"
