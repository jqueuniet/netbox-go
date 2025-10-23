#!/bin/sh

./patcher/pre_gen.py "${1}"

podman run --rm \
  -v "${PWD}:/local:Z" \
  --userns=keep-id \
  --workdir /local \
  --env _JAVA_OPTIONS='-DmaxYamlCodePoints=99999999' \
  docker.io/openapitools/openapi-generator-cli:latest \
  generate -c configs/oapi-codegen-v${1}.yaml

./patcher/post_gen.sh "${1}"
