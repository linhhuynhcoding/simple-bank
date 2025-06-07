package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/linhhuynhcoding/learn-go/util"
	"github.com/stretchr/testify/require"
	a "github.com/linhhuynhcoding/learn-go/db/accountdb"

)

func CreateRandomAccount(t *testing.T) a.Account {
	arg := a.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := TestQueries.Account.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	account2, err := TestQueries.Account.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.ID, account2.ID)
	require.Equal(t, account2.Balance, account2.Balance)
	require.Equal(t, account2.Owner, account2.Owner)
	require.Equal(t, account2.Currency, account2.Currency)
	require.WithinDuration(t, account2.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := a.UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	newAccount, err := TestQueries.Account.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newAccount)

	require.Equal(t, account.ID, newAccount.ID)
	require.Equal(t, arg.Balance, newAccount.Balance)
	require.Equal(t, account.Owner, newAccount.Owner)
	require.Equal(t, account.Currency, newAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, newAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	err := TestQueries.Account.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err2 := TestQueries.Account.GetAccount(context.Background(), account.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}
