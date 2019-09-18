package assets

import (
	"bytes"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Referral struct {
	Base
	Referral string `json:"referral"`
}

func (asset *Referral) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Referral) ApplyUnconfirmed(sender *models.Account) error {
	referral := repositories.Accounts.Get(asset.Referral)

	sender.Referral = referral

	return nil
}

func (asset *Referral) UndoUnconfirmed(sender *models.Account) error {
	sender.Referral = nil

	return nil
}

func (asset *Referral) GetBytes() []byte {
	buf := new(bytes.Buffer)

	return buf.Bytes()
}
