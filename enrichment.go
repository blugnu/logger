package unilog

import "context"

type EnrichmentFunc func(context.Context, Enricher) Enricher

var enrichmentFuncs []EnrichmentFunc

func WithEnrichment(d EnrichmentFunc) {
	enrichmentFuncs = append(enrichmentFuncs, d)
}
