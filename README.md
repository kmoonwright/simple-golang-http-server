# Simple Golang HTTP Server/Client with OpenTelemetry Instrumentation

## Setup 
1. Install dependencies
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace
go get go.opentelemetry.io/otel/sdk/trace
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp

2. Set your OpenTelemetry OTLP endpoint for exported telemetry. Since we're exporting with http and not grpc, the endpoint will automatically be concatenated with `/v1/traces`.
``` bash
export OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:4318"
```

3. Set the API Key for your observability platform, we'll use Honeycomb.io
``` bash
export HONEYCOMB_API_KEY="MY_SECRET_API_KEY"
```

4. Run the server and client
In one terminal, start the server:
``` bash
go run server.go
```

5. In another terminal, start the client.
``` bash
go run client.go
```

The server and client will send trace data to the OpenTelemetry Collector, which will then log the generated trace data, however both the server and Collector (detailed below) will need to be running to generate a trace. This example demonstrates how to instrument a Golang server and client using the net/http package with OpenTelemetry libraries and propagate the trace from the server to the client using an OTLP exporter for the generated trace data.


## Collector
This example uses the 0.75.0 Darwin ARM64 Collector binary. To start the OpenTelemetry Collector using the provided YAML configuration file, follow these steps:

1. Install OpenTelemetry Collector:

First, you need to install the OpenTelemetry Collector binary. You can download the latest release from the GitHub releases page: https://github.com/open-telemetry/opentelemetry-collector/releases. Choose the appropriate binary for your operating system (Windows, Linux, or macOS).

2. Create a configuration file:

Create a new file called config.yaml in your desired directory and paste the contents of the provided YAML configuration into the file. Save the file.

3. Start the OpenTelemetry Collector:

Open a terminal, navigate to the directory containing the OpenTelemetry Collector binary and the config.yaml file, and run the following command:

For Linux or macOS:
``` bash
./otelcol --config config.yaml
```

For Windows:
``` powershell
.\otelcol.exe --config config.yaml
```

This command will start the OpenTelemetry Collector using the provided YAML configuration file. The Collector will listen for incoming trace data on the specified OTLP receiver endpoint and forward the data to the configured exporter.

If you encounter any issues or need to modify the configuration, edit the config.yaml file and restart the OpenTelemetry Collector with the updated configuration.


## Analyzing generated trace
Go to [Honeycomb.io](https://ui.honeycomb.io) to see your recent traces.