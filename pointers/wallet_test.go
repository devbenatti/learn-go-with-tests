package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBalance := func(t *testing.T, wallet Wallet, expect Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expect {
			t.Errorf("result %s - expect %s", result, expect)
		}
	}

	checkError := func(t *testing.T, result error, expect error) {
		t.Helper()
		if result == nil {
			t.Fatal("Expected an error but none occurred")
		}

		if result != expect {
			t.Errorf("error result %s, error expect %s", result, expect)
		}
	}

	checkNonexistentError := func(t *testing.T, result error) {
		t.Helper()
		if result != nil {
			t.Fatal("unexpected error received")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw sufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))
		checkNonexistentError(t, err)
	})

	t.Run("Withdraw with balance insufficient", func(t *testing.T) {
		initialbalance := Bitcoin(20)
		wallet := Wallet{balance: Bitcoin(initialbalance)}
		err := wallet.Withdraw(Bitcoin(100))
		checkBalance(t, wallet, initialbalance)
		checkError(t, err, ErrInsufficientBalance)
	})
}
