#!/bin/sh

podman run --rm -v "${PWD}:/local" \
  --env _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999 \
  docker.io/openapitools/openapi-generator-cli:latest \
  generate -c "/local/configs/oapi-codegen-v${1}.yaml"

go fmt ./...
