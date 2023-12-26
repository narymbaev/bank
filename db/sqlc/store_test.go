package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1, _ := store.GetAccount(context.Background(), int64(59))
	account2, _ := store.GetAccount(context.Background(), int64(60))

	fmt.Println("Balance before transactions ==>> account1: ", account1.Balance, " account2: ", account2.Balance)

	// run n concurrent transfer transactions
	n := 4
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i ++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		//if i % 2 == 1 {
		//	fromAccountID = account2.ID
		//	toAccountID = account1.ID
		//}


		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// TODO: check accountls' balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount1)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount2)

	fmt.Println("Updated Balance after transactions ==>> account1: ", updatedAccount1.Balance, " account2: ", updatedAccount2.Balance)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}