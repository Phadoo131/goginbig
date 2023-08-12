package db

import (
	"context"
	"testing"

	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := Account{
		Owner: util.RandomOwner(),
	}

	account, err := testQuerie.CreateAccount(context.Background(), arg.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func CreateRandomBookInStore(t *testing.T) Instore {
	arg := CreateInstoreParams{
		Book:      util.RandomTitle(),
		Bookcount: util.RandomInt(1, 10),
	}

	bookInstore, err := testQuerie.CreateInstore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bookInstore)

	require.Equal(t, arg.Book, bookInstore.Book)
	require.Equal(t, arg.Bookcount, bookInstore.Bookcount)

	require.NotZero(t, bookInstore.CreatedAt)

	return bookInstore
}

func TestCreateAccount(t *testing.T) {
	CreateRandomBookInStore(t)
	CreateRandomAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	newacc := CreateRandomAccount(t)
	require.NotEmpty(t, newacc)

	err := testQuerie.DeleteAccount(context.Background(), newacc.ID)
	require.NoError(t, err)
}

func TestGetAccount(t *testing.T){
	newacc := CreateRandomAccount(t)
	require.NotEmpty(t, newacc)

	acc, err := testQuerie.GetAccount(context.Background(), newacc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, newacc.ID)
	require.Equal(t, acc.Owner, newacc.Owner)
	require.Equal(t, acc.CreatedAt, newacc.CreatedAt)
}

func TestListAccounts(t *testing.T){
	accounts, err := testQuerie.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}
