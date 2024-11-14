package unitofwork

import (
	"context"
	"reflect"

	database "github.com/sdv-projects/go-ea-common/database"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go -package=mocks

// UnitOfWork is a interface of the unit of work pattern by Martin Fowler.
// https://martinfowler.com/eaaCatalog/unitOfWork.html
type UnitOfWork interface {
	// RegisterNew adds and marks the entity as new
	RegisterNew(entity any)
	// RegisterDirty adds and marks the entity as updated
	RegisterDirty(entity any)
	// RegisterDeleted adds and marks the entity as deleted
	RegisterDeleted(entity any)
	// Commit applies changes to the repository in a single transaction.
	Commit(ctx context.Context) error
}

// UoWRepositoryFactory is a factory interface for getting UoWRepository by Entity type.
type UoWRepositoryFactory interface {
	// Get gets UoWRepository by the Entity type
	Get(ctx context.Context, et reflect.Type) (UoWRepository, error)
}

// UoWRepository is an interface to apply changes to a storage.
type UoWRepository interface {
	// TxInsert implements the logic for inserting new entites into the storage
	TxInsert(ctx context.Context, tx database.Tx, entity ...any) error
	// TxUpdate implements the logic for updating entites into the storage
	TxUpdate(ctx context.Context, tx database.Tx, entity ...any) error
	// TxDelete implements the logic for deleting entites from the storage
	TxDelete(ctx context.Context, tx database.Tx, entity ...any) error
}
