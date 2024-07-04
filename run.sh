#!/bin/bash

# Setup the Go environment
go mod tidy

# Compile the Go program
go build -o main main.go

# Run the Go program
./main

