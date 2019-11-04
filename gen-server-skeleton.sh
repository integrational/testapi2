OASDEFURL=https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore-expanded.yaml
BASEPATH=api
GENPACK=gen
GENPATH=$BASEPATH/$GENPACK
OASDEF=$BASEPATH/oasdef.yaml

curl -o $OASDEF $OASDEFURL

go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

mkdir -p $GENPATH
oapi-codegen -package $GENPACK -generate types -o $GENPATH/types.go $OASDEF
oapi-codegen -package $GENPACK -generate server -o $GENPATH/server.go $OASDEF
oapi-codegen -package $GENPACK -generate spec -o $GENPATH/spec.go $OASDEF

mkdir -p $BASEPATH/impl # for the API implementation proper
mkdir -p $BASEPATH/cmd  # for the main that launches the API implementation server
