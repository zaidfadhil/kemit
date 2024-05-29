#!/bin/sh

command_exists() {
  command -v "$1" >/dev/null 2>&1
}

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux)
        case "$ARCH" in
            x86_64)
                SUFFIX="linux_amd64"
                ;;
            aarch64)
                SUFFIX="linux_arm64"
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
                SUFFIX="darwin_amd64"
                ;;
            arm64)
                SUFFIX="darwin_arm64"
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
  echo "Fetching the latest release tag..."
  RELEASE_TAG=$(curl -s $RELEASE_URL | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

  if [ -z "$RELEASE_TAG" ]; then
    echo "Failed to fetch the release tag."
    exit 1
  fi

  DOWNLOAD_URL="https://github.com/$REPO/releases/download/$RELEASE_TAG/kemit_${RELEASE_TAG#v}_${SUFFIX}.tar.gz"

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
