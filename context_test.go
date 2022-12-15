package logger

// import (
// 	"context"
// 	"testing"

// 	"github.com/sirupsen/logrus"
// )

// func TestFromContext(t *testing.T) {
// 	// ARRANGE
// 	type key string
// 	ctx := context.WithValue(context.Background(), key("foo"), "bar")

// 	// use a logrus logger to test since this is capable of log enrichment
// 	logger := logrus.New()
// 	var log Logger = Logger{ctx, &LogrusAdapter{logger}}

// 	// a decorator is needed for FromContext to get the keyed values from
// 	// the context and add them to the log entry
// 	od := decorators
// 	defer func() { decorators = od }()

// 	Decorator(func(ctx context.Context, log Logger) Logger {
// 		v := ctx.Value(key("foo"))
// 		return log.WithField("foo", v)
// 	})

// 	// ACT
// 	log = log.fromContext()
// 	entry := log.Adapter.(*logrusEntry)

// 	// ASSERT
// 	keys := make([]string, 0, len(entry.Data))
// 	for k := range entry.Data {
// 		keys = append(keys, k)
// 	}

// 	t.Run("context keys", func(t *testing.T) {
// 		wanted := 1
// 		got := len(keys)
// 		if wanted != got {
// 			t.Errorf("wanted %v, got %v", wanted, got)
// 		}
// 	})

// 	k := keys[0]
// 	v := entry.Data[k]

// 	t.Run("field key", func(t *testing.T) {
// 		wanted := "foo"
// 		got := k
// 		if wanted != got {
// 			t.Errorf("wanted %v, got %v", wanted, got)
// 		}
// 	})

// 	t.Run("field value", func(t *testing.T) {
// 		wanted := "bar"
// 		got := v
// 		if wanted != got {
// 			t.Errorf("wanted %v, got %v", wanted, got)
// 		}
// 	})
// }
