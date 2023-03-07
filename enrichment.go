package logger

import "context"

type EnrichmentFunc func(context.Context, LogEntry) LogEntry

var enrichmentFuncs []EnrichmentFunc

func WithEnrichment(d EnrichmentFunc) {
	enrichmentFuncs = append(enrichmentFuncs, d)
}
