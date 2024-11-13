#!/bin/sh

podman run --rm -v $PWD:/local \
  --env _JAVA_OPTIONS=-DmaxYamlCodePoints=99999999 \
  docker.io/openapitools/openapi-generator-cli:latest \
  generate -i "/local/${1}" -o /local/v4_0 \
  -c /local/config-oapi-codegen.yaml
