package db

import (
	"context"
	"database/sql"
	"fmt"

	a "github.com/linhhuynhcoding/learn-go/db/accountdb"
	e "github.com/linhhuynhcoding/learn-go/db/entrydb"
	t "github.com/linhhuynhcoding/learn-go/db/transferdb"
)

// Store provide all functions to excute db queries
type Store struct {
	Account  *a.Queries
	Entry    *e.Queries
	Transfer *t.Queries
	db       *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:       db,
		Account:  a.New(db),
		Entry:    e.New(db),
		Transfer: t.New(db),
	}
}

// TxStore provide all functions to excute db transactions
type TxStore struct {
	Account  *a.Queries
	Entry    *e.Queries
	Transfer *t.Queries
}

// NewTxStore create a new NewTxStore
func NewTxStore(tx *sql.Tx) *TxStore {
	return &TxStore{
		Account:  a.New(tx),
		Entry:    e.New(tx),
		Transfer: t.New(tx),
	}
}

// execTx excutes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*TxStore) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := NewTxStore(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    t.Transfer `json:"transfer"`
	FromAccount a.Account  `json:"from_account"`
	ToAccount   a.Account  `json:"to_account"`
	FromEntry   e.Entry    `json:"from_entry"`
	ToEntry     e.Entry    `json:"to_entry"`
}

// TransferTX performs a money transfer from one account to the orther
// It creates a transfer record, add account entries, and update accounts's balance within a single db transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *TxStore) error {
		var err error

		result.Transfer, err = q.Transfer.CreateTransfer(ctx, t.CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.Entry.CreateEntry(ctx, e.CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry, err = q.Entry.CreateEntry(ctx, e.CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		// TODO: update accounts's balance

		return nil
	})

	return result, err
}
