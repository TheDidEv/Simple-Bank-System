package db

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestTransferTx(t *testing.T) {
// 	// store := NewStore(testDB)

// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)
// 	fmt.Println(">> before", account1.Balance, account2.Balance)

// 	// run a concurrent transfer transactions
// 	n := 2
// 	amount := int64(10)

// 	errs := make(chan error)
// 	results := make(chan TransferTxResult)

// 	for i := 0; i < n; i++ {
// 		go func() {
// 			ctx := context.Background()
// 			result, err := testStore.TransferTx(ctx, TransferTxParams{
// 				FromAccountID: account1.ID,
// 				ToAccountID:   account2.ID,
// 				Amount:        amount,
// 			})

// 			errs <- err
// 			results <- result
// 		}()
// 	}

// 	existed := make(map[int]bool)
// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)

// 		result := <-results
// 		require.NotEmpty(t, result)

// 		// check transfer
// 		transfer := result.Transfer
// 		require.NotEmpty(t, transfer)
// 		require.Equal(t, account1.ID, transfer.FromAccountID)
// 		require.Equal(t, account2.ID, transfer.ToAccountID)
// 		require.Equal(t, amount, transfer.Amount)
// 		require.NotZero(t, transfer.ID)
// 		require.NotZero(t, transfer.CreatedAt)

// 		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
// 		require.NoError(t, err)

// 		// check entry
// 		fromEntry := result.FromEntry
// 		require.NotEmpty(t, fromEntry)
// 		require.Equal(t, account1.ID, fromEntry.Amount)
// 		require.Equal(t, -amount, fromEntry.Amount)
// 		require.NotZero(t, fromEntry.ID)
// 		require.NotZero(t, fromEntry.CreatedAt)

// 		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
// 		require.NoError(t, err)

// 		// to entry
// 		toEntry := result.ToEntry
// 		require.NotEmpty(t, toEntry)
// 		require.Equal(t, account1.ID, toEntry.Amount)
// 		require.Equal(t, amount, toEntry.Amount)
// 		require.NotZero(t, toEntry.ID)
// 		require.NotZero(t, toEntry.CreatedAt)

// 		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
// 		require.NoError(t, err)

// 		// check account
// 		fromAccount := result.FromAccount
// 		require.NotEmpty(t, fromAccount)
// 		require.Equal(t, account1.ID, fromAccount.ID)

// 		toAccount := result.ToAccount
// 		require.NotEmpty(t, toAccount)
// 		require.Equal(t, account2.ID, toAccount.ID)

// 		// check accounts' balance

// 		fmt.Println(">> tx", fromAccount.Balance, toAccount.Balance)
// 		diff1 := account1.Balance - fromAccount.Balance
// 		diff2 := toAccount.Balance - account2.Balance
// 		require.Equal(t, diff1, diff2)
// 		require.True(t, diff1 > 0)
// 		require.True(t, diff1%amount == 0)

// 		k := int(diff1 / amount)
// 		require.True(t, k >= 1 && k <= n)
// 		require.NotContains(t, existed, k)
// 		existed[k] = true

// 		// check the final update balance
// 		updateAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
// 		require.NoError(t, err)

// 		updateAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
// 		require.NoError(t, err)

// 		fmt.Println(">> after", updateAccount1.Balance, updateAccount2.Balance)
// 		require.Equal(t, account1.Balance-int64(n)*amount, updateAccount1.Balance)
// 		require.Equal(t, account2.Balance+int64(n)*amount, updateAccount2.Balance)

// 	}
// }

// func TestTransferTxDeadLock(t *testing.T) {
// 	// store := NewStore(testDB)

// 	account1 := createRandomAccount(t)
// 	account2 := createRandomAccount(t)
// 	fmt.Println(">> before", account1.Balance, account2.Balance)

// 	// run a concurrent transfer transactions
// 	n := 10
// 	amount := int64(10)

// 	errs := make(chan error)

// 	for i := 0; i < n; i++ {
// 		fromAccountId := account1.ID
// 		toAccountId := account2.ID

// 		if i%2 == 1 {
// 			fromAccountId = account2.ID
// 			toAccountId = account1.ID
// 		}

// 		go func() {
// 			ctx := context.Background()
// 			_, err := testStore.TransferTx(ctx, TransferTxParams{
// 				FromAccountID: fromAccountId,
// 				ToAccountID:   toAccountId,
// 				Amount:        amount,
// 			})

// 			errs <- err
// 		}()
// 	}

// 	for i := 0; i < n; i++ {
// 		err := <-errs
// 		require.NoError(t, err)

// 		// check the final update balance
// 		updateAccount1, err := testStore.GetAccount(context.Background(), account1.ID)
// 		require.NoError(t, err)

// 		updateAccount2, err := testStore.GetAccount(context.Background(), account2.ID)
// 		require.NoError(t, err)

// 		fmt.Println(">> after", updateAccount1.Balance, updateAccount2.Balance)
// 		require.Equal(t, account1.Balance, updateAccount1.Balance)
// 		require.Equal(t, account2.Balance, updateAccount2.Balance)

// 	}
// }
