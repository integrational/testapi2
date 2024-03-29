// Package gen provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
