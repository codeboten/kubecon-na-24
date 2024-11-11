#!/bin/bash


export OTEL_SERVICE_VERSION=0.4
export OTEL_SERVICE_NAME=client 
CLIENT='go run ../demo-go/cmd/client/main.go --config ../config.yaml'

while :
do
  ${CLIENT} --addr http://localhost:8080/rolldice 
  ${CLIENT} --addr http://localhost:19999/rolldice/
  OTEL_EXPORTER_OTLP_ENDPOINT=localhost:24317 otel-cli exec --service collector-trace-generator --name "GET http://localhost:19999/" curl http://localhost:19999
  sleep 2
done