// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetSession(ctx context.Context, id pgtype.UUID) (Session, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
