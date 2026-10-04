package services

import (
	"Mybankapp/models"
	"errors"
	"fmt"
)

func UpdateAccount(updated models.Account) error {
	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	for i := range accounts {
		if accounts[i].Username == updated.Username {
			accounts[i] = updated
			return SaveAccounts(accounts)
		}
	}

	return errors.New("Account does not exist.")
}
func Deposit(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("Account is nil.")
	}
	if amount <= 0 {
		return errors.New(
			"Deposit amount must be greater than zero",
		)
	}

	updated := *acc
	updated.Balance += amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Withdraw(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if amount <= 0 {
		return errors.New(
			"withdraw amount must be greater than zero",
		)
	}
	if amount > acc.Balance {
		return errors.New("account balance is not enough")
	}

	updated := *acc
	updated.Balance -= amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Transfer(from *models.Account, toUser string, amount float64) error {
	if from == nil {
		return errors.New("Sender account it NULL")
	}
	if amount <= 0 {
		return fmt.Errorf("")
	}
	if from.Username == toUser {
		return fmt.Errorf("kho n")
	}
	accounts, _ := LoadAccounts()
	var sender *models.Account
	var receiver *models.Account
	for i, acc := range accounts {
		if acc.Username == from.Username {
			sender = &accounts[i]
		}
		if acc.Username == toUser {
			receiver = &accounts[i]
		}
	}
	if sender == nil && receiver == nil {
		return fmt.Errorf("not found sender or receiver")
	}
	if sender.Balance < amount {
		return fmt.Errorf("not enough cash, please do not do this")
	}

	sender.Balance -= amount
	receiver.Balance += amount
	if err := SaveAccounts(accounts); err != nil {
		return err
	} 
	from.Balance = sender.Balance
	return nil
}
 