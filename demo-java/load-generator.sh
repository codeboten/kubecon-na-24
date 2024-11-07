#!/bin/bash

while :
do
  echo "GET http://localhost:8080/actuator/health"
  curl "http://localhost:8080/actuator/health" || true
  echo

  echo "GET http://localhost:8080/vets.html"
  curl -H "Accept: text/html" "http://localhost:8080/vets.html" > /dev/null || true
  echo

    echo "GET http://localhost:8081/actuator/health"
    curl "http://localhost:8081/actuator/health" || true
    echo

    echo "GET http://localhost:8081/vets.html"
    curl -H "Accept: text/html" "http://localhost:8081/vets.html" > /dev/null || true
    echo


  sleep 2
done