#!/bin/sh

goimports -w v"${1}"
# Needs to be run twice for some reason
go fmt v"${1}"/*.go
go fmt v"${1}"/*.go
go mod tidy
