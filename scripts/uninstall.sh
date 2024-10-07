#!/bin/sh

INSTALL_DIR="/usr/local/bin"
BINARY_NAME="kemit"

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

echo "Starting the uninstallation process..."

# Check if the binary exists
if command_exists $BINARY_NAME; then
    echo "Removing $BINARY_NAME from $INSTALL_DIR..."
    sudo rm "$INSTALL_DIR/$BINARY_NAME"
    if [ $? -eq 0 ]; then
        echo "$BINARY_NAME was successfully removed from $INSTALL_DIR."
    else
        echo "Failed to remove $BINARY_NAME from $INSTALL_DIR."
        exit 1
    fi
else
    echo "$BINARY_NAME is not installed."
fi

echo "Uninstallation complete."
