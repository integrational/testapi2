PORT=9004
OASDEF=https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore-expanded.yaml

docker run --rm --name apisprout -d -p $PORT:8000 danielgtaylor/apisprout --validate-request $OASDEF
