module github.com/dmitryax/opentelemetry-go/log

go 1.21

require (
	github.com/go-logr/logr v1.4.2
	github.com/stretchr/testify v1.9.0
	github.com/dmitryax/opentelemetry-go v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/dmitryax/opentelemetry-go/metric v1.28.0 // indirect
	github.com/dmitryax/opentelemetry-go/trace v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dmitryax/opentelemetry-go/metric => ../metric

replace github.com/dmitryax/opentelemetry-go => ../

replace github.com/dmitryax/opentelemetry-go/trace => ../trace
