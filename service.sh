#!/bin/bash

# Source profile untuk mendapatkan environment variables
source ~/.profile
source ~/.nvm/nvm.sh

# Lokasi aplikasi
APP_DIR="~/api"
cd $APP_DIR

# Get PM2 path dynamically
PM2_PATH=$(which pm2)
if [ -z "$PM2_PATH" ]; then
    echo "Error: PM2 not found in PATH"
    exit 1
fi

case "$1" in
    start)
        echo "Starting genuicomm-api..."
        $PM2_PATH start ecosystem.config.js
        ;;
    stop)
        echo "Stopping genuicomm-api..."
        $PM2_PATH stop genuicomm-api
        ;;
    restart)
        echo "Restarting genuicomm-api..."
        $PM2_PATH restart genuicomm-api
        ;;
    status)
        echo "Checking genuicomm-api status..."
        $PM2_PATH show genuicomm-api
        ;;
    logs)
        echo "Showing genuicomm-api logs..."
        $PM2_PATH logs genuicomm-api
        ;;
    save)
        echo "Saving genuicomm-api..."
        $PM2_PATH save
        ;;
    *)
        echo "Usage: $0 {start|stop|restart|status|logs|save}"
        exit 1
        ;;
esac

exit 0 