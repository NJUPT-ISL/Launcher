#!/bin/zsh
echo "Release Linux amd64 version"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o launcher \
&& zip launcher-linux-x86_64.zip launcher && rm -f launcher
echo "Release Linux arm version"
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o launcher \
&& zip launcher-linux-x86_64.zip launcher && rm -f launcher
echo "Release MacOS version"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o launcher \
&& zip launcher-MacOS-x86_64.zip launcher && rm -f launcher
echo "Release Windows version"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o launcher.exe \
&& zip launcher-Windows-x86_64.zip launcher.exe && rm -f launcher.exe