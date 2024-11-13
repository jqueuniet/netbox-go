#!/bin/sh

./patcher/pre_gen.py "${1}"

export _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999
openapi-generator generate -c "configs/oapi-codegen-v${1}.yaml"

./patcher/post_gen.sh "${1}"
