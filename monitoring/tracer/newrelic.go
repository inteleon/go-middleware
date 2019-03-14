package tracer

import (
	"context"

	newrelic "github.com/newrelic/go-agent"

	monitoringctx "github.com/inteleon/go-middleware/monitoring/context"
)

type newRelic struct {
	trace string
	nrtx  newrelic.Transaction
	nrseg *newrelic.Segment
}

// NewNewRelic creates and returns a new New Relic tracer.
func NewNewRelic(ctx context.Context, trace string) *newRelic {
	nrtx, ok := ctx.Value(monitoringctx.NewRelicKey).(newrelic.Transaction)
	if !ok {
		nrtx = nil
	}

	return &newRelic{
		trace: trace,
		nrtx:  nrtx,
	}
}

// Begin starts tracing a New Relic transaction segment.
func (n *newRelic) Begin() {
	if n.nrtx == nil {
		return
	}

	seg := newrelic.StartSegment(n.nrtx, n.trace)
	n.nrseg = seg
}

// End ends the tracing of a New Relic transaction segment.
func (n *newRelic) End() {
	if n.nrseg == nil {
		return
	}

	n.nrseg.End()
}
