module github.com/te2ka/golang-interceptor-example/service/http

go 1.23

replace (
	github.com/te2ka/golang-interceptor-example/utils/log => ../../utils/log
	github.com/te2ka/golang-interceptor-example/utils/tracer => ../../utils/tracer
)

require (
	github.com/te2ka/golang-interceptor-example/utils/log v0.0.0-00010101000000-000000000000
	github.com/te2ka/golang-interceptor-example/utils/tracer v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
