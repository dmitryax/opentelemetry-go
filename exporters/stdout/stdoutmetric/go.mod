module github.com/dmitryax/opentelemetry-go/exporters/stdout/stdoutmetric

go 1.21

require (
	github.com/dmitryax/opentelemetry-go v1.28.0
	github.com/dmitryax/opentelemetry-go/sdk v1.28.0
	github.com/dmitryax/opentelemetry-go/sdk/metric v1.28.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dmitryax/opentelemetry-go/metric v1.28.0 // indirect
	github.com/dmitryax/opentelemetry-go/trace v1.28.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dmitryax/opentelemetry-go/metric => ../../../metric

replace github.com/dmitryax/opentelemetry-go => ../../..

replace github.com/dmitryax/opentelemetry-go/sdk/metric => ../../../sdk/metric

replace github.com/dmitryax/opentelemetry-go/trace => ../../../trace

replace github.com/dmitryax/opentelemetry-go/sdk => ../../../sdk
