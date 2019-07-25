package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("%#v:, Got %s, wanted %s", wallet, got, want)
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted error, didnt get one :(")
		}

		if err != want {
			t.Errorf("wanted %q, got %q", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got error, dindnt want one!")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("withdraw with poor person", func(t *testing.T) {
		start := Bitcoin(10)
		wallet := Wallet{balance: start}
		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, start)
		assertError(t, err, ErrNoMoney)
	})
}
