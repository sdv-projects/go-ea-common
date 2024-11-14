package database

import (
	"context"
	"errors"
	"fmt"
	"reflect"
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

// TxRollback performs Rollback with error suppression
var TxRollback func(tx Tx) = func(tx Tx) { TxRollbackWithLog(tx, nil) }

// TxRollbackWithLog performs Rollback with error suppression and write error to log.
// If lw is nil, then the error isn't written to the log.
func TxRollbackWithLog(tx Tx, lw func(err error)) {
	if tx == nil {
		return
	}
	if err := tx.Rollback(); err != nil && lw != nil {
		lw(err)
	}
}

// CastTx is a helper generic function for casting the Tx interface
// to the target transaction type.
// T is the type of the target transaction instance.
func CastTx[T interface{ Tx }](tx Tx) (result T, err error) {
	if tx == nil {
		return result, errors.New("function was called with empty data")
	}

	var ok bool
	if result, ok = tx.(T); ok {
		return
	}

	return result, fmt.Errorf(
		"unable to cast object '%s' to '%s'",
		reflect.TypeOf(tx).String(),
		reflect.TypeOf(result).String(),
	)
}
