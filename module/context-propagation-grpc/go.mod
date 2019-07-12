module github.com/AminoApps/context-propagation-go/module/context-propagation-grpc

go 1.12

require (
	github.com/AminoApps/context-propagation-go v0.0.0-20190712093550-b461a62a3c33
	github.com/golang/protobuf v1.3.2
	github.com/opentracing-contrib/go-grpc v0.0.0-20180928155321-4b5a12d3ff02
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7
	google.golang.org/grpc v1.22.0
)

replace github.com/AminoApps/context-propagation-go => ../..
