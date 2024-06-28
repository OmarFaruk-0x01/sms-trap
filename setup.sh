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
    TEMP_DIR=$(mktemp -d)

    echo "Downloading $FILE_NAME..."
    wget "${BASE_URL}/${FILE_NAME}" -O "${TEMP_DIR}/${FILE_NAME}"

    echo "Downloading checksum file..."
    wget "${BASE_URL}/${CHECKSUM_FILE}" -O "${TEMP_DIR}/${CHECKSUM_FILE}"

    echo "Verifying checksum..."
    CHECKSUM=$(grep "${FILE_NAME}" "${TEMP_DIR}/${CHECKSUM_FILE}" | awk '{ print $1 }')
    echo "${CHECKSUM}  ${TEMP_DIR}/${FILE_NAME}" | sha256sum -c -

    if [[ $? -ne 0 ]]; then
        echo "Checksum verification failed!"
        rm -r "${TEMP_DIR}"
        exit 1
    fi

    echo "Extracting $FILE_NAME..."
    tar -xzf "${TEMP_DIR}/${FILE_NAME}" -C "${TEMP_DIR}"

    echo "Moving files to /usr/local/bin..."
    sudo mv "${TEMP_DIR}/sms-trap" /usr/local/bin/

    echo "Cleaning up..."
    rm -r "${TEMP_DIR}"

    echo "SMS Trap setup complete!"
}

# Main script execution
detect_os_arch
get_latest_version
download_and_setup