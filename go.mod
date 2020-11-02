module github.com/3Rivers/index

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/3Rivers/helloworld v0.0.0-20201030054725-78344f9d694f
	github.com/3Rivers/order v0.0.0-20201030085137-7ece0048b1e7
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.9.1
)
