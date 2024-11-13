#!/bin/sh

export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
openapi-generator generate \
  -c configs/oapi-codegen-v4_0.yaml