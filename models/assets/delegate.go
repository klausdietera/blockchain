package assets

import (
	"bytes"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

type Delegate struct {
	Base
	Username string `json:"username"`
}

func (*Delegate) CalculateFee() int64 {
	return 1000000000
}

func (asset *Delegate) VerifyUnconfirmed(sender *models.Account) error {
	return nil
}

func (asset *Delegate) ApplyUnconfirmed(sender *models.Account) error {
	delegate := models.Delegate{
		Username: asset.Username,
	}

	sender.Delegate = &delegate

	repositories.Delegates.Add(sender)

	return nil
}

func (asset *Delegate) UndoUnconfirmed(sender *models.Account) error {
	repositories.Delegates.Remove(sender.PublicKey)

	sender.Delegate = nil

	return nil
}

func (asset *Delegate) GetBytes() []byte {
	buf := new(bytes.Buffer)

	return buf.Bytes()
}
