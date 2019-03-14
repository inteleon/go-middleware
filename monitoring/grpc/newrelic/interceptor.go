package newrelic

import (
	"context"

	nragent "github.com/newrelic/go-agent"
	"google.golang.org/grpc"

	monitoringctx "github.com/inteleon/go-middleware/monitoring/context"
)

type interceptor struct {
	newRelicApp nragent.Application
}

// NewInterceptor creates and returns a new New Relic interceptor.
func NewInterceptor(newRelicApp nragent.Application) *interceptor {
	return &interceptor{
		newRelicApp: newRelicApp,
	}
}

// Intercept is the grpc.UnaryServerInterceptor function for tracing GRPC requests using New Relic.
func (i *interceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	txn := i.newRelicApp.StartTransaction(info.FullMethod, nil, nil)
	defer txn.End()

	ctx = context.WithValue(ctx, monitoringctx.NewRelicKey, txn)

	res, err := handler(ctx, req)
	if err != nil {
		txn.NoticeError(err)
	}

	return res, err
}
