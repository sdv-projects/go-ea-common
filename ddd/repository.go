package ddd

import (
	"context"
)

//go:generate mockgen --source=repository.go --destination=mocks/repository.go --package=mocks

// ReadOnlyRepository is a generic repository interface for read-only repository access.
type ReadOnlyRepository[ID any, Entity any] interface {
	// GetByID gets an *Entity by identifier.
	GetByID(ctx context.Context, id ID) (*Entity, error)
	// Find gets the list of entities by a
	Find(ctx context.Context, spec Specification[Entity]) ([]*Entity, error)
	// Single gets a one entity by
	Single(ctx context.Context, spec Specification[Entity]) (*Entity, error)
	// SingleOrDefault gets one object by  It returns the default entity if not found.
	SingleOrDefault(ctx context.Context, spec Specification[Entity]) (*Entity, error)
	// First gets the first entity by
	First(ctx context.Context, spec Specification[Entity]) (*Entity, error)
	// FirstOrDefault gets the first entity by  It returns the default entity if not found.
	FirstOrDefault(ctx context.Context, spec Specification[Entity]) (*Entity, error)
	// Count get the number of entities by
	Count(ctx context.Context, spec Specification[Entity]) (int, error)
	// IsExist checks whether there are entities according to the
	IsExist(ctx context.Context, spec Specification[Entity]) (bool, error)
}

// Repository is a generic repository interface for accessing the repository for change.
type Repository[ID any, Entity any] interface {
	// Repository include access methods from ReadOnlyRepository.
	ReadOnlyRepository[ID, Entity]
	// Add inserts new entity.
	Add(ctx context.Context, e *Entity) (*Entity, error)
	// AddRange inserts the list of new entity.
	AddRange(ctx context.Context, list []*Entity) ([]*Entity, error)
	// Update updates entity.
	Update(ctx context.Context, e *Entity) (*Entity, error)
	// UpdateRange updates the list of entities.
	UpdateRange(ctx context.Context, e []*Entity) ([]*Entity, error)
	// AddOrUpdate inserts or updates entity.
	AddOrUpdate(ctx context.Context, e *Entity) (*Entity, error)
	// AddOrUpdateRange  inserts or updates the list of entities.
	AddOrUpdateRange(ctx context.Context, e []*Entity) ([]*Entity, error)
	// Delete removes entity.
	Delete(ctx context.Context, e *Entity) error
	// DeleteRange removes the list of entities.
	DeleteRange(ctx context.Context, e []*Entity) error
}
