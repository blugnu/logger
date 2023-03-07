package logger

import (
	"context"
	"testing"
)

func TestDecorator(t *testing.T) {
	// ARRANGE
	oef := enrichmentFuncs
	defer func() { enrichmentFuncs = oef }()

	f := func(ctx context.Context, log LogEntry) LogEntry { return log }

	// ACT
	if len(oef) != 0 {
		t.Fatal("`decorators` is not empty")
	}
	WithEnrichment(f)

	// ASSERT
	wanted := 1
	got := len(enrichmentFuncs)
	if wanted != got {
		t.Errorf("wanted %v, got %v", wanted, got)
	}
}
