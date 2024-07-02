module github.com/dmitryax/opentelemetry-go/example/zipkin

go 1.21

replace (
	github.com/dmitryax/opentelemetry-go => ../..
	github.com/dmitryax/opentelemetry-go/exporters/zipkin => ../../exporters/zipkin
	github.com/dmitryax/opentelemetry-go/sdk => ../../sdk
)

require (
	github.com/dmitryax/opentelemetry-go v1.28.0
	github.com/dmitryax/opentelemetry-go/exporters/zipkin v1.28.0
	github.com/dmitryax/opentelemetry-go/sdk v1.28.0
	github.com/dmitryax/opentelemetry-go/trace v1.28.0
)

require (
	github.com/dmitryax/opentelemetry-go/metric v1.28.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace github.com/dmitryax/opentelemetry-go/trace => ../../trace

replace github.com/dmitryax/opentelemetry-go/metric => ../../metric
