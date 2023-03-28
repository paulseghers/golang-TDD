package main

import (
	"fmt"
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

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		fmt.Printf("address os wallet in test: %v \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)
		wallet := Wallet{startingBalance}
		_ = wallet.Withdraw(Bitcoin(12.0))
		assertBalance(t, wallet, 8.0)

	})
	t.Run("withdrawTooMuch", func(t *testing.T) {
		startingBalance := Bitcoin(20.0)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(420))
		assertError(t, err)
		assertBalance(t, wallet, startingBalance)

	})

}
