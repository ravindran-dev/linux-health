
#!/bin/sh

set -e

BIN_NAME=linux-health
INSTALL_DIR=/usr/local/bin

echo "Building $BIN_NAME..."
go build -o $BIN_NAME ./cmd/linux-health

echo "Installing to $INSTALL_DIR..."
sudo mv $BIN_NAME $INSTALL_DIR/

echo "Installed successfully."
