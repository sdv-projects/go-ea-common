package cqrs

import (
	"context"
)

//go:generate mockgen --source=command.go --destination=mocks/command.go --package=mocks

// CommandHandler is a generic interface for the command handler in application layer (usecase).
// C is a type of the command.
type CommandHandler[C any] interface {
	// Handle executes the command.
	Execute(ctx context.Context, command C) error
}

// CommandHandlerWithResult is a generic interface for the command handler with result.
// C is a type of the command, R is a type of the result.
type CommandHandlerWithResult[C any, R any] interface {
	// Handle executes the command.
	Execute(ctx context.Context, command C) (R, error)
}
