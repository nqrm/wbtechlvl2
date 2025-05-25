package pattern

import "fmt"

/*
Реализовать паттерны, объяснить применимость каждого паттерна, плюсы и минусы, а также реальные примеры использования паттерна на практике.
1. Паттерн «Фасад».

Применимость:
	- Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
	- Когда необходимо разложить подсистему на отдельные слои.

Плюсы:
	- Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
	- Фасад может быть привязанным ко всем (или многим, даже к которым не совсем должен) структурам программы
*/

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &WalletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		notification: &Notification{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(accountID string, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	return nil
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.balance = w.balance - amount
	return nil
}

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	notification *Notification
}
