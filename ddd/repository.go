package ddd

import (
	"context"
	"errors"
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
	UpdateRange(ctx context.Context, list []*Entity) ([]*Entity, error)
	// AddOrUpdate inserts or updates entity.
	AddOrUpdate(ctx context.Context, e *Entity) (*Entity, error)
	// AddOrUpdateRange  inserts or updates the list of entities.
	AddOrUpdateRange(ctx context.Context, list []*Entity) ([]*Entity, error)
	// Delete removes entity.
	Delete(ctx context.Context, e *Entity) error
	// DeleteRange removes the list of entities.
	DeleteRange(ctx context.Context, list []*Entity) error
}

type UnimplementedReadOnlyRepository[ID any, Entity any] struct {
}

type UnimplementedRepository[ID any, Entity any] struct {
	UnimplementedReadOnlyRepository[ID, Entity]
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) GetByID(ctx context.Context, id ID) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) Find(
	ctx context.Context,
	spec Specification[Entity]) ([]*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) Single(
	ctx context.Context,
	spec Specification[Entity]) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) SingleOrDefault(
	ctx context.Context,
	spec Specification[Entity]) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) First(
	ctx context.Context,
	spec Specification[Entity]) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) FirstOrDefault(
	ctx context.Context,
	spec Specification[Entity]) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) Count(
	ctx context.Context,
	spec Specification[Entity]) (int, error) {
	return 0, errors.ErrUnsupported
}

func (*UnimplementedReadOnlyRepository[ID, Entity]) IsExist(
	ctx context.Context,
	spec Specification[Entity]) (bool, error) {
	return false, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) Add(
	ctx context.Context,
	e *Entity) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) AddRange(ctx context.Context, list []*Entity) ([]*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) Update(ctx context.Context, e *Entity) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) UpdateRange(ctx context.Context, e []*Entity) ([]*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) AddOrUpdate(ctx context.Context, e *Entity) (*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) AddOrUpdateRange(ctx context.Context, e []*Entity) ([]*Entity, error) {
	return nil, errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) Delete(ctx context.Context, e *Entity) error {
	return errors.ErrUnsupported
}

func (*UnimplementedRepository[ID, Entity]) DeleteRange(ctx context.Context, e []*Entity) error {
	return errors.ErrUnsupported
}
