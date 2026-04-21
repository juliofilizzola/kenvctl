#!/bin/bash
set -e

# Limpa builds antigos
echo "Limpando builds antigos..."
make clean || true

# Build para cada sistema

echo "Buildando para Windows..."
GOOS=windows GOARCH=amd64 go build -o kenvctl.exe ./cmd

echo "Buildando para Linux..."
GOOS=linux GOARCH=amd64 go build -o kenvctl-linux-amd64 ./cmd

echo "Buildando para Mac..."
GOOS=darwin GOARCH=amd64 go build -o kenvctl-mac-amd64 ./cmd

# Empacota cada binário

echo "Empacotando Windows..."
zip kenvctl-windows-amd64.zip kenvctl.exe README.md LICENSE || true

echo "Empacotando Linux..."
tar czvf kenvctl-linux-amd64.tar.gz kenvctl-linux-amd64 README.md LICENSE || true

echo "Empacotando Mac..."
tar czvf kenvctl-mac-amd64.tar.gz kenvctl-mac-amd64 README.md LICENSE || true

echo "Pronto! Os arquivos para release estão na pasta atual."

