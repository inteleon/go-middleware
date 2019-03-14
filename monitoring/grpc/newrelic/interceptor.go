package newrelic

import (
	"context"

	nragent "github.com/newrelic/go-agent"
	"google.golang.org/grpc"

	monitoringctx "github.com/inteleon/go-middleware/monitoring/context"
)

// Interceptor is the New Relic GRPC interceptor.
type Interceptor struct {
	newRelicApp nragent.Application
}

// Intercept is the grpc.UnaryServerInterceptor function for tracing GRPC requests using New Relic.
func (i *Interceptor) Intercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	txn := i.newRelicApp.StartTransaction(info.FullMethod, nil, nil)
	defer txn.End()

	ctx = context.WithValue(ctx, monitoringctx.NewRelicKey, txn)

	res, err := handler(ctx, req)
	if err != nil {
		txn.NoticeError(err)
	}

	return res, err
}
