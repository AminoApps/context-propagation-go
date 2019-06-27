# Context Propagation Go

This middleware is used for support propagate context between micro services.

For this version, we propagate context by [opentracing baggage](https://github.com/opentracing/specification/blob/master/specification.md) protocol.

Supported framework:

* [Gin](module/context-propagation-gin)
* [Standard Http Client](module/context-propagation-http)
* [gRPC Server](module/context-propagation-grpc)
* [gRPC Client](module/context-propagation-grpc)


## How to use

### Get data from context

```shell
go get -u github.com/AminoApps/context-propagation-go
```

```go
package main

import cpg "github.com/AminoApps/context-propagation-go"


ctx = cpg.SetValueToContext(ctx, "my-key", "my-value")

valye := cpg.GetValueFromContext(context.Background(), "my-key")
```

### Gin

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-gin
```

```go
package main 

import cpgin "github.com/AminoApps/context-propagation-go/module/context-propagation-gin"

e := gin.New()
e.Use(cpgin.Middleware())
```

### Http Client

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-http
```

```go
package main

import cphttp "github.com/AminoApps/context-propagation-go/module/context-propagation-http"

client := cphttp.WrapClient(&http.Client{})

// Please use the ctxhttp to wrap the request.
resp, err := ctxhttp.Get(ctx, client, "http://test.com")
```

### Grpc Client/Server

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-grpc
```

```go
package main

import cpgrpc "github.com/AminoApps/context-propagation-go/module/context-propagation-grpc"

server := grpc.NewServer(grpc.UnaryInterceptor(cpgrpc.NewUnaryServerInterceptor()))

client := grpc.Dial(address, grpc.WithUnaryInterceptor(cpgrpc.NewUnaryClientInterceptor()))
```