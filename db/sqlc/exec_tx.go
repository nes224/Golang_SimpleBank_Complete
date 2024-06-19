package db

import (
	"context"
	"fmt"
)

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx) // connection the database
	if err != nil {
		return err
	}

	q := New(tx) // created transaction
	err = fn(q)  // call the input function with that queries, and get back an error.
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx) // If all operations in the transaction are successful, we simply commit the transaction with tx.Commit(), and return its error to the caller.
}


