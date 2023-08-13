package db

import (
	"context"
	"testing"

	"github.com/Phadoo131/goginbig/util"
	"github.com/stretchr/testify/require"
)


func TestCreateInstore(t *testing.T) {
	uniqueBookName := util.RandomTitle() 
	bookcount := util.RandomInt(1, 10)

	instore, err := testQuerie.CreateInstore(context.Background(), CreateInstoreParams{
		Book:      uniqueBookName,
		Bookcount: bookcount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, instore)

	require.Equal(t, uniqueBookName, instore.Book)
}

func TestGetInstore(t *testing.T) {
	instore, err := testQuerie.GetInstore(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, instore)
}

func TestGetInstoreForUpdate(t *testing.T) {
	newBook := CreateRandomBookInStore(t)

	instore, err := testQuerie.GetInstoreForUpdate(context.Background(), newBook.Book)
	require.NoError(t, err)
	require.NotEmpty(t, instore)

	require.Equal(t, newBook.Book, instore.Book)
}

func TestListInstore(t *testing.T) {
	var limit int32 = 10
	var offset int32 = 0

	// Creating random instore entries
	for i := 0; i < int(limit); i++ {
		CreateRandomBookInStore(t)
	}

	instoreList, err := testQuerie.ListInstore(context.Background(), ListInstoreParams{
		Limit:  limit,
		Offset: offset,
	})
	require.NoError(t, err)
	require.NotEmpty(t, instoreList)
	require.Len(t, instoreList, int(limit))
}

func TestUpdateInstore(t *testing.T) {
	newBook := CreateRandomBookInStore(t)

	newBookcount := util.RandomInt(1, 10)

	updatedInstore, err := testQuerie.UpdateInstore(context.Background(), UpdateInstoreParams{
		Book:      newBook.Book,
		Bookcount: newBookcount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, updatedInstore)

	require.Equal(t, newBook.Book, updatedInstore.Book)
	require.Equal(t, newBookcount, updatedInstore.Bookcount)
}