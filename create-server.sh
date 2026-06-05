#!/bin/bash

set -e

if [ $# -ne 1 ]; then
    echo "Usage: $0 <service-name>"
    exit 1
fi

SERVICE_NAME="$1"

cd app || exit 1

if [ -d "$SERVICE_NAME" ]; then
    echo "Service '$SERVICE_NAME' already exists"
    exit 1
fi

mkdir -p "$SERVICE_NAME"

echo "Creating api..."
cd "$SERVICE_NAME"
goctl api new "${SERVICE_NAME}_api"
goctl rpc new "${SERVICE_NAME}rpc"

echo "Done."