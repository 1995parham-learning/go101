# Go

## Introduction

The Go programming language, also known as Golang, has gained popularity among software practitioners since its release in 2009. This language offers various features and advantages that make it a preferred choice for many developers.
Golang is an awesome language, but it has sanctioned our country. It is similar to C programming language.

We can use it for writing code as terminal applications or servers. I had experience in using it for writing ReST API and GrapQL servers. For the terminal application there are many fun library out there like [`bubbletea`](https://github.com/charmbracelet/bubbletea) and [`tcell`](https://github.com/gdamore/tcell) which provide easy-to-use interfaces for creating interactive command-line applications.

## Go Modules

In Go 1.16, module-aware commands report an error after discovering a problem in `go.mod` or `go.sum` instead of attempting to fix the problem automatically.
In most cases, the error message recommends a command to fix the problem.

## `gotip`

One useful tool that comes with Golang is `gotip`, which allows you to test upcoming changes and experimental features.

```bash
go install golang.org/dl/gotip@latest
gotip download
```

And then use the `gotip` command as if it were your normal go command.

After installation, use the `gotip` command instead of your normal go command to have latest features.
To update, run `gotip download` again. This will always download the main branch.
To download an alternative branch, run `gotip download BRANCH` and to download a specific CL, run `gotip download NUMBER`.

## `libc`

You can control to use `cgo` with `CGO_ENABLED` flag in go build.
We have different implementation of C library.

- GNU C Library (`glibc`)
- `musl`
- Microsoft C Runtime Library

## Testing with `testify`

To test with go, [testify](https://pkg.go.dev/github.com/stretchr/testify) is an awesome library.
It has suite, require and assert.

Always use `_test` prefix on packages for writing tests but in case of internal tests
in which you need to access private package members use `_internal_test.go` as filename.

```go
// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
    suite.Suite
    VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
    suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
    suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(ExampleTestSuite))
}
```

```go
// SetupAllSuite has a SetupSuite method, which will run before the
// tests in the suite are run.
type SetupAllSuite interface {
 SetupSuite()
}

// SetupTestSuite has a SetupTest method, which will run before each
// test in the suite.
type SetupTestSuite interface {
 SetupTest()
}

// TearDownAllSuite has a TearDownSuite method, which will run after
// all the tests in the suite have been run.
type TearDownAllSuite interface {
 TearDownSuite()
}

// TearDownTestSuite has a TearDownTest method, which will run after
// each test in the suite.
type TearDownTestSuite interface {
 TearDownTest()
}

// BeforeTest has a function to be executed right before the test
// starts and receives the suite and test names as input
type BeforeTest interface {
 BeforeTest(suiteName, testName string)
}

// AfterTest has a function to be executed right after the test
// finishes and receives the suite and test names as input
type AfterTest interface {
 AfterTest(suiteName, testName string)
}

// WithStats implements HandleStats, a function that will be executed
// when a test suite is finished. The stats contain information about
// the execution of that suite and its tests.
type WithStats interface {
 HandleStats(suiteName string, stats *SuiteInformation)
}
```

Writing tests with `testify` is awesome, so use them. Also, I write tests with mock for the higher modules and tests the low level one with the real dependencies.
For application that has really great mocks like `redis`, I have used them instead of real one.

## Logging

As I said there are mainly two types of applications developed by Golang.
In the terminal applications, it is better to use [pterm](https://pterm.sh) to show the logs and other information
instead of simply printing them.
In case of the server applications, [zap](https://github.com/uber-go/zap) is a better choice because it has structure logging,
and you can also write the logs on console and syslog at the same time.

Structure logging increase the search efficiency for when you want to search among your
logs on your log aggregation system.

## Telemetry

The current status of the major functional components for OpenTelemetry Go is as follows:

| Tracing | Metrics | Logging             |
| ------- | ------- | ------------------- |
| Stable  | Alpha   | Not Yet Implemented |

With this release we are introducing a split in module versions.
The tracing API and SDK are entering the v1.0.0 Release Candidate phase with v1.0.0-RC1 while the experimental metrics API and SDK continue with v0.x releases at v0.21.0.
Modules at major version 1 or greater will not depend on modules with major version 0.

[OpenTelemetry Go API and SDK](https://github.com/open-telemetry/opentelemetry-go)

### Exporter

The SDK requires an exporter to be created. Exporters are packages that allow telemetry data to be emitted somewhere - either to the console (which is what we’re doing here), or to a remote system or collector for further analysis and/or enrichment. OpenTelemetry supports a variety of exporters through its ecosystem including popular open source tools like Jaeger, Zipkin, and Prometheus.

```go
import stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

func main() {
 exporter, err := stdout.New(
  stdout.WithPrettyPrint(),
 )
 if err != nil {
  log.Fatalf("failed to initialize stdout export pipeline: %v", err)
 }
}
```

Exporter depend on the task and the service for example the Jaeger exporter is in the following package:

```go
import "go.opentelemetry.io/otel/exporters/jaeger"
```

### Tracer Provider

This block of code will create a new batch span processor, a type of span processor that batches up multiple spans over a period of time, that writes to the exporter we created in the previous step.

```go
 ctx := context.Background()

 bsp := sdktrace.NewBatchSpanProcessor(exporter)
 s := sdktrace.AlwaysSampler()
 tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp), sdktrace.WithSampler(s))

 // Handle this error in a sensible manner where possible
 defer func() { _ = tp.Shutdown(ctx) }()
```

OpenTelemetry requires a trace provider to be initialized in order to generate traces. A trace provider can have multiple span processors, which are components that allow for span data to be modified or exported after it’s created.

### Quick Start

First, we’re asking the global trace provider for an instance of a tracer, which is the object that manages spans for our service.

```go
tracer := otel.Tracer("ex.com/basic")
ctx = baggage.ContextWithValues(ctx,
    fooKey.String("foo1"),
    barKey.String("bar1"),
)

func(ctx context.Context) {
    var span trace.Span
    ctx, span = tracer.Start(ctx, "operation")
    defer span.End()

    span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
    span.SetAttributes(anotherKey.String("yes"))

    meter.RecordBatch(
        // Note: call-site variables added as context Entries:
        baggage.ContextWithValues(ctx, anotherKey.String("xyz")),
        commonAttributes,

        valueRecorder.Measurement(2.0),
    )

    func(ctx context.Context) {
        var span trace.Span
        ctx, span = tracer.Start(ctx, "Sub operation...")
        defer span.End()

        span.SetAttributes(lemonsKey.String("five"))
        span.AddEvent("Sub span event")
        boundRecorder.Record(ctx, 1.3)
    }(ctx)
}(ctx)
```

### Contribution

There are tracers for internals of go libraries available at opentelemetry-go-contrib with experimental version v0.21.0:

- [mongodb](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo)
- [echo](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/instrumentation/github.com/labstack/echo/otelecho)

## RabbitMQ

There is an official library for AMQP in Golang:

[An AMQP 0-9-1 Go client maintained by the RabbitMQ team](https://github.com/rabbitmq/amqp091-go)

Also, there is another official library for Streams in Golang:

[A client library for RabbitMQ streams](https://github.com/rabbitmq/rabbitmq-stream-go-client)

## Projects

### [Array](array)

Creates arrays and slices. It tries to demo some features of Go in Array.

### [Beehive Hello](beehive-hello)

Hello world example in [beehive](https://github.com/kandoo/beehive).

### [Cast](cast)

Type switch and Interface cast example.

### [Closures](closures)

How do closures work with upper scope variables?

### [Conditional Variable](condvar)

Introduction to go conditional variables.

### [defer](defer)

Go Defer Here and Now!

### [echo-server](echo-server)

Says back everything you say to it.

### [fibonacci](fibonacci)

Fibonacci sequence in Go.

### [go-c](go-c)

Adds some C to Go. :yum:

### [Monte Carlo](monte-carlo)

Monte-Carlo method for estimating the PI. check [here](https://academo.org/demos/estimating-pi-monte-carlo/) for more information.

### [Once](once)

Go uses zero initiation but sometimes you want some initiation and you want it lazy ...

### [Redis](redis)

Let's use redis as a cache with MessagePack coding

### [roy](roy)

In the Memory of Roya Taheri.

### [rwp](rwp)

The reader-writer-problem in Go.

### [twiddle](twiddle)

[Stackoverflow](https://stackoverflow.com/questions/47778453/generate-combinations-permutation-of-specific-length)
