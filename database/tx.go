package database

import (
	"context"
)

//go:generate mockgen -source=tx.go -destination=mocks/tx.go -package=database_mocks

// Tx is an in-progress repository transaction.
type Tx interface {
	// Commit the transaction.
	Commit() error
	// Rollback aborts the transaction.
	Rollback() error
}

// TxFactory is a factory interface to create transactions.
type TxFactory interface {
	// BeginTx creates new transaction.
	BeginTx(ctx context.Context) (Tx, error)
}

// TxRollback performs Rollback with error suppression and write error to log.
// If logFunc is nil, then the error isn't written to the log.
func TxRollback(tx Tx, logFunc func(err error)) {
	if tx == nil {
		return
	}
	if err := tx.Rollback(); err != nil && logFunc != nil {
		logFunc(err)
	}
}
