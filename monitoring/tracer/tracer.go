package tracer

// Tracer defines the interface for all tracers to implement.
type Tracer interface {
	Begin(trace string)
	End()
}
