module github.com/dmitryax/opentelemetry-go/sdk

go 1.21

replace github.com/dmitryax/opentelemetry-go => ../

require (
	github.com/go-logr/logr v1.4.2
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	github.com/dmitryax/opentelemetry-go v1.28.0
	github.com/dmitryax/opentelemetry-go/trace v1.28.0
	golang.org/x/sys v0.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/dmitryax/opentelemetry-go/metric v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dmitryax/opentelemetry-go/trace => ../trace

replace github.com/dmitryax/opentelemetry-go/metric => ../metric
