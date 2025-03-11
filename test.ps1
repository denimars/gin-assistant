# PowerShell script to build and move the executable

# Build the Go project
go build

if (Test-Path test_) {
    Remove-Item -Recurse -Force test_
}

New-Item -ItemType Directory -Path test_

Move-Item -Path gin-assistant.exe -Destination test_
