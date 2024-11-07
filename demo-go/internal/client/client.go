// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Sample contains a simple client that periodically makes a simple http request
// to a server and exports to the OpenTelemetry service.
package client

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/metric"
)

func handleErr(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}

func Run(ctx context.Context, addr string) {
	tracer := otel.Tracer("demo-client-tracer")
	meter := otel.Meter("demo-client-meter")

	method, _ := baggage.NewMember("method", "repl")
	client, _ := baggage.NewMember("client", "cli")
	bag, _ := baggage.New(method, client)

	// labels represent additional key-value descriptors that can be bound to a
	// metric observer or recorder.
	// TODO: Use baggage when supported to extract labels from baggage.
	commonLabels := []attribute.KeyValue{
		attribute.String("method", "repl"),
		attribute.String("client", "cli"),
	}

	// Recorder metric example
	requestLatency, _ := meter.Float64Histogram(
		"demo_client/request_latency",
		metric.WithDescription("The latency of requests processed"),
	)

	// TODO: Use a view to just count number of measurements for requestLatency when available.
	requestCount, _ := meter.Int64Counter(
		"demo_client/request_counts",
		metric.WithDescription("The number of requests processed"),
	)

	lineLengths, _ := meter.Int64Histogram(
		"demo_client/line_lengths",
		metric.WithDescription("The lengths of the various lines in"),
	)

	// TODO: Use a view to just count number of measurements for lineLengths when available.
	lineCounts, _ := meter.Int64Counter(
		"demo_client/line_counts",
		metric.WithDescription("The counts of the lines in"),
	)

	defaultCtx := baggage.ContextWithBaggage(ctx, bag)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		startTime := time.Now()
		ctx, span := tracer.Start(defaultCtx, "ExecuteRequest")
		makeRequest(ctx, addr)
		span.End()
		latencyMs := float64(time.Since(startTime)) / 1e6
		nr := int(rng.Int31n(7))
		for i := 0; i < nr; i++ {
			randLineLength := rng.Int63n(999)
			lineCounts.Add(ctx, 1, metric.WithAttributes(commonLabels...))
			lineLengths.Record(ctx, randLineLength, metric.WithAttributes(commonLabels...))
			// fmt.Printf("#%d: LineLength: %dBy\n", i, randLineLength)
		}

		requestLatency.Record(ctx, latencyMs, metric.WithAttributes(commonLabels...))
		requestCount.Add(ctx, 1, metric.WithAttributes(commonLabels...))

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func makeRequest(ctx context.Context, addr string) {

	// Trace an HTTP client by wrapping the transport
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Make sure we pass the context to the request to avoid broken traces.
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s/rolldice/", addr), nil)
	if err != nil {
		handleErr(err, "failed to http request")
	}

	// All requests made with this client will create spans.
	res, err := client.Do(req)
	if err != nil {
		handleErr(err, "error on request")
	}
	res.Body.Close()
}
