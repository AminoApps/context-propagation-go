[![Build Status](https://travis-ci.org/AminoApps/context-propagation-go.svg?branch=master)](https://travis-ci.org/AminoApps/context-propagation-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/AminoApps/context-propagation-go)](https://goreportcard.com/report/github.com/AminoApps/context-propagation-go)

# Context Propagation Go

This middleware is used for support propagate context between micro services.

For this version, we propagate context by [opentracing baggage](https://github.com/opentracing/specification/blob/master/specification.md) protocol.

Supported framework for auto inject and extract:

* [Gin](module/context-propagation-gin)
* [Standard Http Server](module/context-propagation-http)
* [Standard Http Client](module/context-propagation-http)
* [gRPC Server](module/context-propagation-grpc)
* [gRPC Client](module/context-propagation-grpc)


## How to use

### Operate data from context

> Before get or set data from context, you should enable auto inject and extract.

```shell
go get -u github.com/AminoApps/context-propagation-go
```

```go
package main

import cpg "github.com/AminoApps/context-propagation-go"


ctx = cpg.SetValueToContext(ctx, "my-key", "my-value")

valye := cpg.GetValueFromContext(context.Background(), "my-key")
```

### Auto inject and extract

#### Gin

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-gin
```

```go
package main 

import cpgin "github.com/AminoApps/context-propagation-go/module/context-propagation-gin"

e := gin.New()
e.Use(cpgin.Middleware())
```

#### Http Client/Server

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-http
```

```go
package main

import cphttp "github.com/AminoApps/context-propagation-go/module/context-propagation-http"

http.ListenAndServe(":8080", cphttp.Wrap(myHandler))

client := cphttp.WrapClient(&http.Client{})

// Please use the ctxhttp to wrap the request.
resp, err := ctxhttp.Get(ctx, client, "http://127.0.0.1:8080/test")
```

#### Grpc Client/Server

```shell
go get -u github.com/AminoApps/context-propagation-go/module/context-propagation-grpc
```

```go
package main

import cpgrpc "github.com/AminoApps/context-propagation-go/module/context-propagation-grpc"

server := grpc.NewServer(grpc.UnaryInterceptor(cpgrpc.NewUnaryServerInterceptor()))

client := grpc.Dial(address, grpc.WithUnaryInterceptor(cpgrpc.NewUnaryClientInterceptor()))
```