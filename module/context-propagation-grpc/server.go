package cpgrpc

import (
	"github.com/AminoApps/context-propagation-go"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//NewUnaryServerInterceptor wrap for grpc server
func NewUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)

		if ok {
			headersWithFirst := make(map[string]string, md.Len())
			for k, v := range md {
				if len(v) > 0 {
					headersWithFirst[k] = v[0]
				}
			}

			carrier := cp.Extract(headersWithFirst)
			if len(carrier) > 0 {
				ctx = context.WithValue(ctx, cp.InternalContextKey{}, carrier)
			}

		}

		return handler(ctx, req)
	}
}
