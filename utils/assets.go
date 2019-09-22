package utils

import (
	"encoding/json"
	"errors"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
)

func UnmarshalAsset(bytes []byte, t types.TransactionType) (models.IAsset, error) {
	switch t {
	case types.TransactionReferral:
		var asset assets.Referral
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case types.TransactionSend:
		var asset assets.Send
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	case types.TransactionDelegate:
		var asset assets.Delegate
		err := json.Unmarshal(bytes, &asset)
		return &asset, err
	}
	return nil, errors.New("Invalid transaction type")
}

func CreateSendAsset(recipientPublicKey string, amount int64) assets.Send {
	return assets.Send{
		RecipientPublicKey: recipientPublicKey,
		Amount:             amount,
	}
}
