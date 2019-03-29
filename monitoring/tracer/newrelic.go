package tracer

import (
	"context"

	newrelic "github.com/newrelic/go-agent"

	monitoringctx "github.com/inteleon/go-middleware/monitoring/context"
)

type newRelic struct {
	nrtx  newrelic.Transaction
	nrseg *newrelic.Segment
}

// NewNewRelic creates and returns a new New Relic tracer.
func NewNewRelic(ctx context.Context) *newRelic {
	nrtx, ok := ctx.Value(monitoringctx.NewRelicKey).(newrelic.Transaction)
	if !ok {
		nrtx = nil
	}

	return &newRelic{
		nrtx: nrtx,
	}
}

// Begin starts tracing a New Relic transaction segment.
func (n *newRelic) Begin(trace string) {
	if n.nrtx == nil {
		return
	}

	n.nrseg = newrelic.StartSegment(n.nrtx, trace)
}

// End ends the tracing of a New Relic transaction segment.
func (n *newRelic) End() {
	if n.nrseg == nil {
		return
	}

	n.nrseg.End()
}

// Transaction returns the New Relic transaction object.
func (n *newRelic) Transaction() newrelic.Transaction {
	return n.nrtx
}
