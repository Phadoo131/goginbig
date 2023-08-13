package service

import (
	"context"
	"database/sql"

	"testing"

	db "github.com/Phadoo131/goginbig/db/sqlc"
	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)

var testdbConn *sql.DB
var testQueries *db.Queries

func TestInit(t *testing.T) {
	var err error
	dbConn, err = sql.Open(DBDriver, DBSource)
	require.NoError(t, err, "Failed to connect to the database")

	testQueries = db.New(dbConn)
}

func TestCreateUserAccount(t *testing.T) {

	newAccount := &db.Account{
		Owner: util.RandomOwner(), 
	}

	acc, err := testQueries.CreateAccount(context.Background(), newAccount.Owner)
	require.NoError(t, err)

	require.Equal(t, newAccount.Owner, acc.Owner)

}


func TestDeleteUserAccount(t *testing.T) {
    newAccount := &db.Account{}

	testQueries.DeleteAccount(context.Background(), newAccount.ID)

	require.Empty(t, newAccount.ID)
	require.Empty(t, newAccount.Owner)
	require.Empty(t, newAccount.CreatedAt)
}

func TestGetUserAccount(t *testing.T) {
    newAccount := &db.Account{}

	acc, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, newAccount.ID)
	require.Equal(t, acc.Owner, newAccount.Owner)
	require.Equal(t, acc.CreatedAt, newAccount.CreatedAt)
}

func TestGetAllUserAccounts(t *testing.T) {
	acc, err := testQueries.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, acc)
}
