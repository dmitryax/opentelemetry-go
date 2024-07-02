module github.com/dmitryax/opentelemetry-go/example/opencensus

go 1.21

replace (
	github.com/dmitryax/opentelemetry-go => ../..
	github.com/dmitryax/opentelemetry-go/bridge/opencensus => ../../bridge/opencensus
	github.com/dmitryax/opentelemetry-go/sdk => ../../sdk
)

require (
	github.com/dmitryax/opentelemetry-go v1.28.0
	github.com/dmitryax/opentelemetry-go/bridge/opencensus v1.28.0
	github.com/dmitryax/opentelemetry-go/exporters/stdout/stdoutmetric v1.28.0
	github.com/dmitryax/opentelemetry-go/exporters/stdout/stdouttrace v1.28.0
	github.com/dmitryax/opentelemetry-go/sdk v1.28.0
	github.com/dmitryax/opentelemetry-go/sdk/metric v1.28.0
	go.opencensus.io v0.24.0
)

require (
	github.com/dmitryax/opentelemetry-go/metric v1.28.0 // indirect
	github.com/dmitryax/opentelemetry-go/trace v1.28.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.6.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace github.com/dmitryax/opentelemetry-go/metric => ../../metric

replace github.com/dmitryax/opentelemetry-go/sdk/metric => ../../sdk/metric

replace github.com/dmitryax/opentelemetry-go/trace => ../../trace

replace github.com/dmitryax/opentelemetry-go/exporters/stdout/stdoutmetric => ../../exporters/stdout/stdoutmetric

replace github.com/dmitryax/opentelemetry-go/exporters/stdout/stdouttrace => ../../exporters/stdout/stdouttrace
