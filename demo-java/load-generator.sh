#!/bin/bash

while :
do
  echo "GET http://localhost:8080/actuator/health"
  curl "http://localhost:8080/actuator/health" || true
  echo

  echo "GET http://localhost:8080/owners?lastName=l"
  curl -H "Accept: text/html" "http://localhost:8080/owners?lastName=" > /dev/null || true
  echo

    echo "GET http://localhost:8081/actuator/health"
    curl "http://localhost:8081/actuator/health" || true
    echo

    echo "GET http://localhost:8081/owners?lastName="
    curl -H "Accept: text/html" "http://localhost:8081/owners?lastName=" > /dev/null || true
    echo


  sleep 2
done