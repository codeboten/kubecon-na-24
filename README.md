# Kubecon NA 2024 demo

This repository contains the demo app used during Kubecon NA co-located
Observability Day.

## Java Demo

Build the app image:

```shell
cd demo-java
docker build .
```

Run the app:
```shell
docker compose up -d
```

Generate load:
```shell
./load-generator.sh
```

Attache Otel terminal UI:
```shell
docker compose attach oteltui
```

## Go Demo

Run the app:

```shell
cd demo-go
make run
```

Generate load:
```shell
./load-generator.sh
```

## PHP Demo

Run the app:

```shell
cd demo-php
make run
```

Generate load:
```shell
./load-generator.sh
```

## Collector Demo

Run the collector:

```shell
cd demo-collector
make run
```

## Viewing data

View data in Jaeger, Prometheus:

* Jaeger UI: http://localhost:16686
* Prometheus UI: http://localhost:9090

## Tasks

- [x] update configuration to add second exporter
- [x] add "load" generator
- [x] add code to configure OTLP exporters programmatically
- [x] add OTel Collector example
- [x] update import to use PR instead of local copy
- [x] test the config in php: otlp export isn't configurable
- [x] Clients for sending traffic to all services (otel-cli)
