package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(11))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got error, but do not want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("missing required error")
	}

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
