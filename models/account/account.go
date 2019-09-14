package account

import (
	"encoding/hex"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/utils"
)

func Create(secret string) models.Account {
	keyPair := utils.GenerateKeyPair(secret)
	publicKey := hex.EncodeToString(keyPair.PublicKey.Bytes)

	return models.Account{
		PublicKey: publicKey,
	}
}
