package types

type TransactionType int

const (
	TransactionReferral  TransactionType = 0
	TransactionSend                      = 10
	TransactionSignature                 = 20
	TransactionDelegate                  = 30
)
