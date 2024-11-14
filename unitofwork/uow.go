package unitofwork

import (
	"context"
	"reflect"
	"sync"

	database "github.com/sdv-projects/go-ea-common/database"
)

// operationType is an enumeration that describes an operation on an entity
type operationType uint8

const (
	otInsert operationType = iota
	otUpdate
	otDelete
)

// entityOperation contains entity and operation
type entityOperation struct {
	Type   operationType
	Entity any
}

var _ UnitOfWork = (*uow)(nil)

// uow implements UnitOfWork interface.
// The struct contains the list of entitites operation/
type uow struct {
	tf database.TxFactory
	rf UoWRepositoryFactory

	operations []entityOperation
	mu         sync.Mutex
}

// register adds entity with operation to list
func (u *uow) register(entity any, opType operationType) {
	if entity == nil {
		return
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	u.operations = append(u.operations, entityOperation{
		Type:   opType,
		Entity: entity,
	})
}

// RegisterNew adds and marks the entity as new
func (u *uow) RegisterNew(entity any) {
	u.register(entity, otInsert)
}

// RegisterDirty adds and marks the entity as updated
func (u *uow) RegisterDirty(entity any) {
	u.register(entity, otUpdate)
}

// RegisterDeleted adds and marks the entity as deleted
func (u *uow) RegisterDeleted(entity any) {
	u.register(entity, otDelete)
}

// Commit applies changes to the repository in a single transaction.
func (u *uow) Commit(ctx context.Context) (err error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	tx, err := u.tf.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			database.TxRollback(tx)
		}
	}()

	for _, op := range u.operations {
		t := reflect.TypeOf(op.Entity)
		repository, err := u.rf.Get(ctx, t)
		if err != nil {
			return err
		}

		var do func(ctx context.Context, tx database.Tx, entity ...any) error
		switch op.Type {
		case otInsert:
			do = repository.TxInsert
		case otUpdate:
			do = repository.TxUpdate
		case otDelete:
			do = repository.TxDelete
		}

		if err := do(ctx, tx, op.Entity); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func NewUoW(tf database.TxFactory, rf UoWRepositoryFactory) *uow {
	return &uow{
		tf:         tf,
		rf:         rf,
		mu:         sync.Mutex{},
		operations: make([]entityOperation, 0),
	}
}
