#!/bin/sh

export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
openapi-generator generate \
  -c "configs/oapi-codegen-v${1}.yaml"

# Needs to be run twice for some reason
go fmt v"${1}"/*.go
go fmt v"${1}"/*.go
