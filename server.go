package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func initTracer() func() {
	exporter, err := otlptracehttp.New(context.Background())
	if err != nil {
		log.Fatalf("Failed to create OTLP trace exporter: %v", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shut down tracer provider: %v", err)
		}
	}
}

func main() {
	shutdown := initTracer()
	defer shutdown()

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}), "HelloHandler"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
