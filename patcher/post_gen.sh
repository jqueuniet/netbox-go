#!/bin/sh

goimports -w apiv"${1}"
# Needs to be run twice for some reason
go fmt apiv"${1}"/*.go
go fmt apiv"${1}"/*.go
go mod tidy
