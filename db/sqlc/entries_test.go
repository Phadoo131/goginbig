package db

import (
	"context"
	"testing"

	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntries(t *testing.T) Entry {
	book := CreateRandomBookInStore(t)
	acc := CreateRandomAccount(t)
	
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Book: book.Book,
		Amount: util.RandomInt(0, 10),
	}

	entry, err := testQuerie.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.Book, entry.Book)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T){
	CreateRandomEntries(t)
}

func TestGetEntry(t *testing.T){
	arg := CreateRandomEntries(t)
	require.NotEmpty(t, arg)

	entry, err := testQuerie.GetEntry(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.Book, entry.Book)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreatedAt)
}

func TestGetEntryForUpdate(t *testing.T){
	arg := CreateRandomEntries(t)
	require.NotEmpty(t, arg)

	entry, err := testQuerie.GetEntryForUpdate(context.Background(), arg.AccountID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.Book, entry.Book)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	book := CreateRandomBookInStore(t)
	acc := CreateRandomAccount(t)

	var testEntryIDs []int64
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: acc.ID,
			Book:      book.Book,
			Amount:    util.RandomInt(0, 19),
		}
		entry, err := testQuerie.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
		testEntryIDs = append(testEntryIDs, entry.ID)
	}

	listEntriesParams := ListEntriesParams{
		AccountID: acc.ID,
		Offset:    0,
		Limit:     5,
	}

	entries, err := testQuerie.ListEntries(context.Background(), listEntriesParams)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.Equal(t, acc.ID, entry.AccountID)
		require.Contains(t, testEntryIDs, entry.ID)
	}
}

func TestUpdateEntries(t *testing.T) {
	book := CreateRandomBookInStore(t)
	acc := CreateRandomAccount(t)

	var testEntryIDs []int64
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: acc.ID,
			Book:      book.Book,
			Amount:    util.RandomInt(0, 19),
		}
		entry, err := testQuerie.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)
		testEntryIDs = append(testEntryIDs, entry.ID)
	}
	newAmount := util.RandomInt(0, 1)

	updated := UpdateEntriesParams{
		AccountID: acc.ID,
		Book:  book.Book,
		Amount: newAmount,
	}

	updatedEntry, err := testQuerie.UpdateEntries(context.Background(), updated)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, newAmount, updatedEntry.Amount)
}