#!/bin/bash

# Setup script for pakyus_commerce
echo "Setting up pakyus_commerce..."

# Check if config.json exists
if [ ! -f "config.json" ]; then
    echo "Creating config.json from example..."
    cp config.example.json config.json
    echo "Created config.json - Please update with your actual values"
else
    echo "config.json already exists"
fi

# Check if config.dev.json exists
if [ ! -f "config.dev.json" ]; then
    echo "config.dev.json not found - using default dev config"
else
    echo "config.dev.json found"
fi

echo "Environment setup:"
echo "  - Development: set APP_ENV=development"
echo "  - Testing: set APP_ENV=testing"
echo "  - Production: set APP_ENV=production (default)"
echo ""
echo "Environment variables (optional):"
echo "  - PAKYUS_DB_USERNAME"
echo "  - PAKYUS_DB_PASSWORD" 
echo "  - PAKYUS_DB_HOST"
echo "  - PAKYUS_DB_PORT"
echo "  - PAKYUS_DB_NAME"
echo "  - PAKYUS_WEB_PORT"
echo ""
echo "Setup complete!"