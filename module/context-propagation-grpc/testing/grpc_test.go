package testing

import (
	"context"
	"github.com/AminoApps/context-propagation-go"
	"github.com/AminoApps/context-propagation-go/module/context-propagation-grpc"
	_ "github.com/opentracing-contrib/go-grpc"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"net"
	"testing"
	"time"
)

type grpcTestServer struct {
}

func (s *grpcTestServer) GetRole(ctx context.Context, msg *Msg) (*Result, error) {
	return &Result{Role: cp.GetValueFromContext(ctx, "auth-role")}, nil
}

func TestGrpc(t *testing.T) {
	go func() {
		lis, err := net.Listen("tcp", ":9080")
		assert.Nil(t, err)
		grpcServer := grpc.NewServer(grpc.UnaryInterceptor(cpgrpc.NewUnaryServerInterceptor()))
		RegisterGrpcTestServer(grpcServer, &grpcTestServer{})
		err = grpcServer.Serve(lis)
		assert.Nil(t, err)
	}()

	time.Sleep(time.Second)

	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure(), grpc.WithUnaryInterceptor(cpgrpc.NewUnaryClientInterceptor()))
	assert.Nil(t, err)
	defer conn.Close()

	client := NewGrpcTestClient(conn)
	ctx := context.Background()
	ctx = cp.SetValueToContext(ctx, "auth-role", "1")
	result, err := client.GetRole(ctx, &Msg{})
	assert.Nil(t, err)

	assert.Equal(t, "1", result.Role)
}
