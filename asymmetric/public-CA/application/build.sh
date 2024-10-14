#!/bin/bash

echo "Buiding the application...."
if [ -e main ]; then
    rm main
fi
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo  -o main .





