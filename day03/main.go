package main

import (
	"fmt"

	"github.com/wmc15156/learngo/day03/accounts"
)

func main() {
	account := accounts.NewAccount("nmmm")

	account.Deposit(10)

	fmt.Println(account.Owner(), account.Balance())

	fmt.Println(account)
}
