package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/codeboten/kubecon-na-24/demo/internal/client"
	"go.opentelemetry.io/contrib/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/log/global"
)

func main() {
	// load config file
	confFlag := flag.String("config", "config.yaml", "config file to parse")
	addrFlag := flag.String("addr", "localhost:8080", "destination")

	flag.Parse()
	b, err := os.ReadFile(*confFlag)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	// parse the config
	conf, err := config.ParseYAML(b)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	ctx := context.Background()
	sdk, err := config.NewSDK(config.WithContext(ctx), config.WithOpenTelemetryConfiguration(*conf))
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	defer func() {
		if err := sdk.Shutdown(context.Background()); err != nil {
			log.Printf("err: %v\n", err)
		}
	}()
	otel.SetTracerProvider(sdk.TracerProvider())
	otel.SetMeterProvider(sdk.MeterProvider())
	global.SetLoggerProvider(sdk.LoggerProvider())

	client.Run(ctx, *addrFlag)
}
