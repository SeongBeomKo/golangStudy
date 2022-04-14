package banking

import (
	"errors"
	"fmt"
)

//private화 시킨 object 개념
type Account struct {
	owner   string
	balance int
}

//*Account --> point to the Account
//NewAcocunt creates Account
// 주소를 리턴해서 객체화
func NewAccount(newOwner string) *Account {

	account := Account{owner: newOwner, balance: 0}
	return &account
}

//a account = receiver
//*Account -> use the account calling the account without making a copy
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

//no exception handler in go
func (a *Account) Withdraw(amount int) error {
	//error handling
	if a.balance < amount {
		return errors.New("cant withdraw you are poor")
	}
	a.balance -= amount

	// errors return errors or nil
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

//account := banking.NewAccount("nico")
//	fmt.Println(account)
// account를 print하면 자동으로 내부적으로 호출된다
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
