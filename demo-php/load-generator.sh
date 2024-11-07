#!/bin/bash

while :
do
  OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 otel-cli exec --service php-load-generator --name "GET http://localhost:8080/rolldice" curl http://localhost:8080/rolldice
  sleep 2
done