package main

import (
	"context"
	"fmt"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"log"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/graphite"
	"contrib.go.opencensus.io/exporter/zipkin"
	openzipkin "github.com/openzipkin/zipkin-go"
	zhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"go.opencensus.io/trace"
)

var (
	FooMessage = stats.Int64("foo", "测试一下 foo", stats.UnitDimensionless)
	BarMessage = stats.Int64("bar", "测试一下 bar", stats.UnitDimensionless)
	BazMessage = stats.Int64("baz", "测试一下 baz", stats.UnitDimensionless)
)

func main() {
	localEndpoint, err := openzipkin.NewEndpoint("example-server", "[myip]:5454")
	if err != nil {
		log.Println(err)
	}

	reporter := zhttp.NewReporter("http://localhost:9411/api/v2/spans")
	defer reporter.Close()

	exporterTrace := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(exporterTrace)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	exporterView, err := graphite.NewExporter(graphite.Options{Namespace: "opencensus", Port: 9109})
	if err != nil {
		log.Fatal(err)
	}
	view.RegisterExporter(exporterView)

	if err = view.Register(
		&view.View{
			Name:        "foo_view",
			Description: "foo message",
			Measure:     FooMessage,
			TagKeys:     nil,
			Aggregation: view.Count(),
		},
		&view.View{
			Name:        "bar_view",
			Description: "bar message",
			Measure:     BarMessage,
			TagKeys:     nil,
			Aggregation: view.Count(),
		},
		&view.View{
			Name:        "baz_view",
			Description: "baz message",
			Measure:     BazMessage,
			TagKeys:     nil,
			Aggregation: view.Count(),
		},
	); err != nil {
		log.Fatalf("Cannot register the view: %v", err)
	}

	view.SetReportingPeriod(1 * time.Second)

	ctx := context.Background()
	osk, _ := tag.NewKey("liangc-macos")
	uid, _ := tag.NewKey("uid-liangc")
	ctx, err = tag.New(ctx,
		tag.Insert(osk, "macOS-10.12.5"),
		tag.Upsert(uid, "cc14514"),
	)
	if err != nil {
		log.Fatal(err)
	}

	go foo(ctx)
	fmt.Println("ok.")

	addr := ":2004"
	log.Printf("Serving at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func foo(ctx context.Context) {

	// Name the current span "/foo"
	stats.Record(ctx, FooMessage.M(1))
	ctx, span := trace.StartSpan(ctx, "/foo")
	defer span.End()
	go func() {
		for {
			// Foo calls bar and baz
			bar(ctx)
			baz(ctx)
		}
	}()
}

func bar(ctx context.Context) {
	stats.Record(ctx, BarMessage.M(1))
	ctx, span := trace.StartSpan(ctx, "/bar")
	defer span.End()

	// Do bar
	time.Sleep(1 * time.Second)
}

func baz(ctx context.Context) {
	stats.Record(ctx, BazMessage.M(1))
	ctx, span := trace.StartSpan(ctx, "/baz")
	defer span.End()

	// Do baz
	time.Sleep(2 * time.Second)
}
