package types

type TransactionType int32

const (
	TransactionReferral  TransactionType = 0
	TransactionSend      TransactionType = 10
	TransactionSignature TransactionType = 20
	TransactionDelegate  TransactionType = 30
)
