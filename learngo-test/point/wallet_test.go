package point

import "testing"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
