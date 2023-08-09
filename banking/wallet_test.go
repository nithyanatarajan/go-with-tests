package banking_test

import (
	"testing"

	"github.com/nithyanatarajan/go-with-tests/banking"
	. "github.com/nithyanatarajan/go-with-tests/test"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := banking.Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := banking.Bitcoin(10)

		AssertEqual(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := banking.Wallet{}
		wallet.Deposit(10)

		err := wallet.Withdraw(9)

		got := wallet.Balance()
		want := banking.Bitcoin(1)

		AssertEqual(t, got, want)
		AssertNoError(t, err)
	})

	t.Run("withdraw to error on insufficient funds", func(t *testing.T) {
		wallet := banking.Wallet{}
		wallet.Deposit(10)

		err := wallet.Withdraw(20)

		AssertError(t, err)
		AssertEqual(t, err, banking.ErrInsufficientFunds)
		AssertEqual(t, wallet.Balance(), banking.Bitcoin(10))
	})
}
