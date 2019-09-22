package utils

import (
	"encoding/hex"
	"encoding/json"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
	"github.com/jamesruan/sodium"
)

func CreateTransaction(data models.Transaction, keyPair sodium.SignKP, secondKeyPair *sodium.SignKP) (models.Transaction, error) {
	var salt string
	if data.Salt != "" {
		salt = data.Salt
	} else {
		salt = RandStringBytesMask(models.SALT_LENGTH)
	}

	var createdAt time.Time
	if data.CreatedAt.IsZero() {
		createdAt = time.Now()
	} else {
		createdAt = data.CreatedAt
	}

	transaction := models.Transaction{
		Type:            data.Type,
		Fee:             data.Asset.CalculateFee(),
		Asset:           data.Asset,
		Salt:            salt,
		CreatedAt:       createdAt,
		SenderPublicKey: hex.EncodeToString(keyPair.PublicKey.Bytes),
	}

	signature, err := transaction.CalculateSignature(keyPair)
	if err != nil {
		return transaction, err
	}

	transaction.Signature = signature

	if secondKeyPair != nil {
		secondSignature, err := transaction.CalculateSignature(*secondKeyPair)
		if err != nil {
			return transaction, err
		}

		transaction.SecondSignature = secondSignature
	}

	id, err := transaction.CalculateID()
	if err != nil {
		return transaction, err
	}

	transaction.ID = id

	return transaction, nil
}

func UnmarshalTransaction(data []byte, transaction *models.Transaction) error {
	var tmp struct {
		Type types.TransactionType `json:"type"`
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	var asset models.IAsset
	switch tmp.Type {
	case types.TransactionSend:
		asset = &assets.Send{}
	}

	transaction.Asset = asset

	return json.Unmarshal(data, transaction)
}
