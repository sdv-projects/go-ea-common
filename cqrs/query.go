package cqrs

import (
	"context"
)

//go:generate mockgen --source=query.go --destination=mocks/query.go --package=mocks

// QueryHandler is a generic interface for the query handler in application layer (usecase)
// Q is a type of the query, R is a type of the result.
type QueryHandler[Q any, R any] interface {
	// Execute performs the query
	Execute(ctx context.Context, query Q) (R, error)
}
