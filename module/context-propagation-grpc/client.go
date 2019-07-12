package cpgrpc

import (
	"github.com/AminoApps/context-propagation-go"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//NewUnaryClientInterceptor wrap for grpc client
func NewUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, resp interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		carrier := ctx.Value(cp.InternalContextKey{})
		headers := cp.Inject(carrier)

		md, ok := metadata.FromOutgoingContext(ctx)

		if !ok {
			md = metadata.New(headers)
		} else {
			md = md.Copy()
			for k, v := range headers {
				md.Set(k, v)
			}
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, resp, cc, opts...)
	}
}
