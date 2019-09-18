package services

import (
	"bytes"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

var (
	Transaction ITransactionService
)

func init() {
	Transaction = &TransactionService{
		accountRepository:     repositories.Accounts,
		transactionRepository: repositories.Transactions,
	}
}

type ITransaction interface {
	Create(data interface{}) *interface{}

	Validate(data interface{}) error

	Verify(data interface{}) error

	ApplyUnconfirmed(data interface{}) error
	UndoUnconfirmed(data interface{}) error

	GetBytes(data interface{}) bytes.Buffer

	Apply(data interface{}) error
	Undo(data interface{}) error
}

type ITransactionService interface {
	// ITransaction

	// CalculateID(tx *models.Transaction) string
	// CalculateHash(tx *models.Transaction) bytes.Buffer
	// ToBytes(tx *models.Transaction) bytes.Buffer

	// IsConfirmed(tx *models.Transaction) bool

	ApplyUnconfirmed(transaction models.Transaction) error
	UndoUnconfirmed(transaction models.Transaction) error
}

type TransactionService struct {
	accountRepository     repositories.IAccountsRepository
	transactionRepository repositories.ITransactionsRepository
}

func (s *TransactionService) Create(data interface{}) models.Transaction {
	return models.Transaction{}
}

func (s *TransactionService) ApplyUnconfirmed(transaction models.Transaction) error {
	// TODO: Add verification that the transaction has already been applied

	s.transactionRepository.AddOne(transaction)

	sender := s.accountRepository.Get(transaction.SenderPublicKey)
	sender.Balance -= transaction.Fee

	return transaction.Asset.ApplyUnconfirmed(sender)
}

func (s *TransactionService) UndoUnconfirmed(transaction models.Transaction) error {
	// TODO: Add check that transaction is not applied

	s.transactionRepository.RemoveByID(transaction.ID)

	sender := s.accountRepository.Get(transaction.SenderPublicKey)
	sender.Balance += transaction.Fee

	return transaction.Asset.UndoUnconfirmed(sender)
}
