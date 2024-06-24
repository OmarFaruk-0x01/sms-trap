#!/bin/bash

# Function to detect OS and architecture
detect_os_arch() {
    OS="$(uname | tr '[:upper:]' '[:lower:]')"
    ARCH="$(uname -m)"

    case $ARCH in
        x86_64)
            ARCH="amd64"
            ;;
        arm64)
            ARCH="arm64"
            ;;
        i386|i686)
            ARCH="386"
            ;;
        *)
            echo "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac

    if [[ "$OS" == "darwin" ]]; then
        OS="darwin"
    elif [[ "$OS" == "linux" ]]; then
        OS="linux"
    else
        echo "Unsupported OS: $OS"
        exit 1
    fi
}

# Function to get the latest version
get_latest_version() {
    LATEST_VERSION=$(curl -s https://api.github.com/repos/OmarFaruk-0x01/sms-trap/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    if [[ -z "$LATEST_VERSION" ]]; then
        echo "Failed to fetch the latest version."
        exit 1
    fi
}

# Function to download and extract the appropriate binary
download_and_setup() {
    BASE_URL="https://github.com/OmarFaruk-0x01/sms-trap/releases/download/${LATEST_VERSION}"
    FILE_NAME="sms-trap_${LATEST_VERSION:1}_${OS}_${ARCH}.tar.gz"
    CHECKSUM_FILE="sms-trap_${LATEST_VERSION:1}_checksums.txt"

    echo "Downloading $FILE_NAME..."
    wget "${BASE_URL}/${FILE_NAME}" -O "/tmp/${FILE_NAME}"

    echo "Downloading checksum file..."
    wget "${BASE_URL}/${CHECKSUM_FILE}" -O "/tmp/${CHECKSUM_FILE}"

    echo "Verifying checksum..."
    CHECKSUM=$(grep "${FILE_NAME}" "/tmp/${CHECKSUM_FILE}" | awk '{ print $1 }')
    echo "${CHECKSUM}  /tmp/${FILE_NAME}" | sha256sum -c -

    if [[ $? -ne 0 ]]; then
        echo "Checksum verification failed!"
        exit 1
    fi

    echo "Extracting $FILE_NAME..."
    tar -xzf "/tmp/${FILE_NAME}" -C /usr/local/bin

    echo "Cleaning up..."
    rm "/tmp/${FILE_NAME}" "/tmp/${CHECKSUM_FILE}"

    echo "SMS Trap setup complete!"
}

# Main script execution
detect_os_arch
get_latest_version
download_and_setup