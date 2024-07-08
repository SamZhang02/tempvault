#!/bin/bash

BINARY_NAME="tempvault"

echo "Building tempvault.."
go build -o $BINARY_NAME

if [ $? -ne 0 ]; then
  echo "Build failed. Please check your Go project for errors."
  exit 1
fi

INSTALL_DIR="/usr/local/bin"

echo "Installing the binary to $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR

if [ $? -ne 0 ]; then
  echo "Installation failed. Please check your permissions and try again."
  exit 1
fi

if command -v $BINARY_NAME &> /dev/null; then
  echo "Installation successful. You can now use '$BINARY_NAME' from the command line."
else
  echo "Installation failed. '$BINARY_NAME' is not in your PATH."
  exit 1
fi
