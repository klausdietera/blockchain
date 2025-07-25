package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"bitbucket.org/axelsheva/blockchain/models/types"
	"github.com/jamesruan/sodium"
)

const SALT_LENGTH = 32

type IVerifier interface {
	VerifyUnconfirmed(sender *Account) error
}

type IApplier interface {
	ApplyUnconfirmed(sender *Account) error
	UndoUnconfirmed(sender *Account) error
}

type IAsset interface {
	IVerifier
	IApplier
	json.Unmarshaler
	GetAmount() int64
	CalculateFee() int64
	GetBytes() []byte
}

type ITransaction interface {
	IVerifier
	IApplier
}

type Transaction struct {
	ID              string                `json:"id"`
	BlockID         string                `json:"blockId"`
	Type            types.TransactionType `json:"type"`
	SenderPublicKey string                `json:"senderPublicKey"`
	Fee             int64                 `json:"fee"`
	Signature       string                `json:"signature"`
	SecondSignature string                `json:"secondSignature"`
	CreatedAt       time.Time             `json:"createdAt"`
	Salt            string                `json:"salt"`
	Asset           IAsset                `json:"asset"`
}

func (transaction *Transaction) GetBytes(skipSignature bool, skipSecondSignature bool) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(transaction.Asset.GetBytes())

	err := binary.Write(buf, binary.LittleEndian, transaction.Type)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	err = binary.Write(buf, binary.LittleEndian, transaction.CreatedAt.Unix())
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	buf.Write([]byte(transaction.Salt))
	buf.Write([]byte(transaction.SenderPublicKey))

	if !skipSignature && transaction.Signature != "" {
		buf.Write([]byte(transaction.Signature))
	}

	if !skipSecondSignature && transaction.SecondSignature != "" {
		buf.Write([]byte(transaction.SecondSignature))
	}

	return buf.Bytes(), nil
}

func (transaction *Transaction) CalculateHash(skipSignature bool) ([32]byte, error) {
	b, err := transaction.GetBytes(skipSignature, skipSignature)
	if err != nil {
		var emptyBytes [32]byte
		return emptyBytes, err
	}

	return sha256.Sum256(b), nil
}

func (transaction *Transaction) CalculateID() (string, error) {
	hash, err := transaction.CalculateHash(false)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash[:]), nil
}

func (transaction *Transaction) CalculateSignature(keyPair sodium.SignKP) (string, error) {
	hash, err := transaction.CalculateHash(true)
	if err != nil {
		return "", err
	}

	b := sodium.Bytes(hash[:])
	signature := b.SignDetached(keyPair.SecretKey)
	return hex.EncodeToString(signature.Bytes), nil
}

func (transaction *Transaction) VerifyUnconfirmed(sender *Account) error {
	amount := transaction.Fee + transaction.Asset.GetAmount()
	if sender.Balance < amount {
		return errors.New("Not enough money")
	}
	return transaction.Asset.VerifyUnconfirmed(sender)
}

func (transaction *Transaction) ApplyUnconfirmed(sender *Account) {
	sender.Balance -= transaction.Fee

	transaction.Asset.ApplyUnconfirmed(sender)
}

func (transaction *Transaction) UndoUnconfirmed(sender *Account) {
	sender.Balance += transaction.Fee

	transaction.Asset.UndoUnconfirmed(sender)
}

// func (transaction *Transaction) UnmarshalJSON(data []byte) error {
// 	var tmp struct {
// 		Type types.Transaction `json:"type"`
// 	}

// 	err := json.Unmarshal(data, &tmp)
// 	if err != nil {
// 		return err
// 	}

// 	var asset IAsset
// 	switch tmp.Type {
// 	case types.SendType:
// 		asset = &SendAsset{}
// 	}

// 	transaction.Asset = asset

// 	return json.Unmarshal(data, transaction)
// 	// return nil
// }
