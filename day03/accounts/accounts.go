package accounts

import (
	"errors"
	"fmt"
)

type account struct {
	owner   string
	balance int
}

var errNoMoney error = errors.New("Can not widthDraw")

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}

	return &account
}

func (a *account) Deposit(amount int) {
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

func (a *account) WidthDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *account) ChangeOwner(owner string) {
	a.owner = owner
}

func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.owner, "'s account.\n Has: ", a.balance)
}
