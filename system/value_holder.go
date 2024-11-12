package system

import (
	"context"
	"sync"
	"sync/atomic"
)

// Value Holder is an object for lazy load.
// https://martinfowler.com/eaaCatalog/lazyLoad.html
type ValueHolder[T any] struct {
	newFunc func(ctx context.Context) (T, error)
	once    sync.Once
	loaded  atomic.Bool
	value   T
	err     error
}

// Get value. The first call triggers the load.
func (l *ValueHolder[T]) Value(ctx context.Context) (T, error) {
	l.once.Do(func() {
		l.value, l.err = l.newFunc(ctx)
		l.loaded.Store(true)
	})

	return l.value, l.err
}

func (l *ValueHolder[T]) Clone() *ValueHolder[T] {
	return &ValueHolder[T]{
		newFunc: l.newFunc,
		once:    sync.Once{},
		loaded:  atomic.Bool{},
	}
}

func NewValueHolder[T any](newFunc func(ctx context.Context) (T, error)) *ValueHolder[T] {
	return &ValueHolder[T]{
		newFunc: newFunc,
		once:    sync.Once{},
		loaded:  atomic.Bool{},
	}
}

func NewValueHolderWithResult[T any](result T) *ValueHolder[T] {
	holder := &ValueHolder[T]{
		newFunc: func(_ context.Context) (T, error) { return result, nil },
		once:    sync.Once{},
		loaded:  atomic.Bool{},
		value:   result,
	}
	_, _ = holder.Value(context.TODO())

	return holder
}
