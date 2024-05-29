#!/bin/sh

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

if !command_exists curl; then
  echo "curl is not installed. Please install curl and try again."
  exit 1
fi

if !command_exists tar; then
  echo "tar is not installed. Please install tar and try again."
  exit 1
fi

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux)
        case "$ARCH" in
            x86_64)
                SUFFIX="linux-amd64"
                ;;
            aarch64)
                SUFFIX="linux-arm64"
                ;;
            *)
                echo "Unsupported architecture: $ARCH"
                exit 1
                ;;
        esac
        ;;
    Darwin)
        case "$ARCH" in
            x86_64)
                SUFFIX="darwin-amd64"
                ;;
            arm64)
                SUFFIX="darwin-arm64"
                ;;
            *)
                echo "Unsupported architecture: $ARCH"
                exit 1
                ;;
        esac
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
      ;;
esac

REPO="zaidfadhil/kemit"
INSTALL_DIR="/usr/local/bin"
LOCAL_BINARY="bin/kemit"

if [ -f "$LOCAL_BINARY" ]; then
  echo "Local binary found. Using it for installation."
  BINARY_PATH="$LOCAL_BINARY"
else
  RELEASE_URL="https://api.github.com/repos/$REPO/releases/latest"
  echo "Fetching the latest release URL for $OS $ARCH..."
  DOWNLOAD_URL=$(curl -s $RELEASE_URL | grep "browser_download_url.*$SUFFIX.tar.gz" | cut -d '"' -f 4)

  if [ -z "$DOWNLOAD_URL" ]; then
      echo "Failed to fetch the download URL."
      exit 1
  fi

  echo "Downloading the latest release..."
  curl -L -o kemit.tar.gz "$DOWNLOAD_URL"

  echo "Extracting the downloaded tarball..."
  tar -xzf kemit.tar.gz
  BINARY_PATH="kemit"
fi

echo "Installing the application..."
mv "$BINARY_PATH" "$INSTALL_DIR"

if command_exists kemit; then
  echo "kemit was successfully installed!"
else
  echo "Failed to install kemit."
  exit 1
fi

if [ -f "kemit.tar.gz" ]; then
  rm kemit.tar.gz
fi
if [ -f "kemit" ]; then
  rm kemit
fi

echo "Installation complete."
