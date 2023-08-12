package service

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/Phadoo131/goginbig/db/sqlc"
	"github.com/stretchr/testify/require"
)

var testdbConn *sql.DB
var testestqueries *db.Queries

func TestInit(t *testing.T) {
	var err error
	dbConn, err = sql.Open(DBDriver, DBSource)
	require.NoError(t, err)

	queries = db.New(dbConn)
}

func TestCreateUserAccount(t *testing.T){
	var newAccount *db.Account

	acc, err := queries.CreateAccount(context.Background(), newAccount.Owner)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, newAccount.ID)
	require.Equal(t, acc.Owner, newAccount.Owner)
	require.Equal(t, acc.CreatedAt, newAccount.CreatedAt)
}

func TestDeleteUserAccount(t *testing.T){
	var newAccount db.Account

	queries.DeleteAccount(context.Background(), newAccount.ID)

	require.Empty(t, newAccount.ID)
	require.Empty(t, newAccount.Owner)
	require.Empty(t, newAccount.CreatedAt)
}

func TestGetUserAccount(t *testing.T){
	var newAccount db.Account

	acc, err := queries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.ID, newAccount.ID)
	require.Equal(t, acc.Owner, newAccount.Owner)
	require.Equal(t, acc.CreatedAt, newAccount.CreatedAt)
}

func TestGetAllUserAccounts(t *testing.T){
	acc, err := queries.ListAccounts(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, acc)
}

