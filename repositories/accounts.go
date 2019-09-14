package repositories

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

var (
	Accounts IAccountsRepository
)

func init() {
	Accounts = &AccountsRepository{
		store: make(map[string]*models.Account),
	}
}

type IAccountsRepository interface {
	Add(account *models.Account)
	Get(publicKey string) *models.Account
	Remove(publicKey string)
}

type AccountsRepository struct {
	store map[string]*models.Account
}

func (r *AccountsRepository) Add(account *models.Account) {
	r.store[account.PublicKey] = account
}

func (r *AccountsRepository) Get(publicKey string) *models.Account {
	account := r.store[publicKey]

	if account == nil {
		account = &models.Account{
			PublicKey: publicKey,
		}
	}

	return account
}

func (r *AccountsRepository) Remove(publicKey string) {
	delete(r.store, publicKey)
}
