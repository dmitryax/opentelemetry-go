module github.com/dmitryax/opentelemetry-go/trace

go 1.21

replace github.com/dmitryax/opentelemetry-go => ../

require (
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.9.0
	github.com/dmitryax/opentelemetry-go v1.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dmitryax/opentelemetry-go/metric => ../metric
