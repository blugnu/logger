package logger

import "context"

type EnrichmentFunc func(context.Context, *Logger) *Logger

var enrichmentFuncs []EnrichmentFunc

func WithEnrichment(d EnrichmentFunc) {
	enrichmentFuncs = append(enrichmentFuncs, d)
}
