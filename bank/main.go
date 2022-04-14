package main

import (
	"bank/banking"
	"fmt"
)

func main() {
	account := banking.NewAccount("nico")
	fmt.Println(account)

	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	//golang에서 string + int contantenation
	output := fmt.Sprintf("%s%d", "balance: ", account.Balance())
	fmt.Println(output)

}
