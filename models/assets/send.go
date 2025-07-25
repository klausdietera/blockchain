package assets

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Send struct {
	Base
	RecipientPublicKey string `json:"recipientPublicKey"`
	Amount             int64  `json:"amount"`
}

func (asset *Send) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Send) GetAmount() int64 {
	return asset.Amount
}

func (asset *Send) CalculateFee() int64 {
	return asset.Amount / 10000
}

func (asset *Send) ApplyUnconfirmed(sender *models.Account) error {
	sender.Balance -= asset.Amount

	recipient := repositories.Accounts.Get(asset.RecipientPublicKey)
	recipient.Balance += asset.Amount

	return nil
}

func (asset *Send) UndoUnconfirmed(sender *models.Account) error {
	sender.Balance += asset.Amount

	recipient := repositories.Accounts.Get(asset.RecipientPublicKey)
	recipient.Balance -= asset.Amount

	return nil
}

func (asset *Send) UnmarshalJSON(data []byte) error {
	var tmp struct {
		RecipientPublicKey string `json:"recipientPublicKey"`
		Amount             int64  `json:"amount"`
	}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	asset.Amount = tmp.Amount
	asset.RecipientPublicKey = tmp.RecipientPublicKey

	return nil
}

func (asset *Send) GetBytes() []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, asset.Amount)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	buf.Write([]byte(asset.RecipientPublicKey))

	return buf.Bytes()
}
