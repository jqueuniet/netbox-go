#!/bin/sh

podman run --rm -v $PWD:/local openapitools/openapi-generator-cli:latest generate -i /local/${1} -g go -o /local/ --git-user-id jqueuniet --git-repo-id netbox-go -c /local/config-oapi-codegen.yaml
