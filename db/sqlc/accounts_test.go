package db

import (
	"context"
	"testing"

	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Book:  util.RandomTitle(),
	}

	account, err := testQuerie.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Book, account.Book)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func CreateRandomBookInStore(t *testing.T) Instore{
	arg := CreateInstoreParams{
		Book: util.RandomTitle(),
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
