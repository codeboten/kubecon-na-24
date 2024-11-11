#!/bin/bash


export OTEL_SERVICE_VERSION=0.4
export OTEL_SERVICE_NAME=client 
CLIENT='go run ../demo-go/cmd/client/main.go --config ../config.yaml'

while :
do
  ${CLIENT} --addr http://localhost:8080/rolldice 
  ${CLIENT} --addr http://localhost:19999/rolldice/
  sleep 2
done