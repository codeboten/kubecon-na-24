# Kubecon NA 2024 demo

This repository contains the demo app used during Kubecon NA co-located
Observability Day.

# Java Demo

Build the app image:

```shell
cd demo-java
docker build .
```

Run the app:
```shell
export NEW_RELIC_API_KEY=<INSERT_NEW_RELIC_API_KEY>
docker compose up
```

Generate load:
```shell
./load-generator.sh
```

View data New Relic, Jaeger, Prometheus:

* Jaeger UI: http://localhost:16686
* Prometheus UI: http://localhost:9090
* New Relic: https://one.newrelic.com/

## Tasks

- [ ] update configuration to add second exporter
- [x] add "load" generator
- [x] add code to configure OTLP exporters programmatically
- [x] add OTel Collector example
- [x] update import to use PR instead of local copy
- [x] test the config in php: otlp export isn't configurable
