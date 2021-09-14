package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertionError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't one.")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertionError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("Got an error but didn't one.")
	}
}