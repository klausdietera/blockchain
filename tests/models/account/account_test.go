package account_test

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/account"
)

func TestCreate(t *testing.T) {
	actualAccount := account.Create("hen worry two thank unfair salmon smile oven gospel grab latin reason")

	expectedAccount := models.Account{
		PublicKey: "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2",
		Balance:   0,
	}

	if expectedAccount.PublicKey != actualAccount.PublicKey {
		t.Errorf("Invalid account public key. Actual %s, Expected %s", actualAccount.PublicKey, expectedAccount.PublicKey)
	}

	if expectedAccount.Balance != actualAccount.Balance {
		t.Errorf("Invalid account balance. Actual %d, Expected %d", actualAccount.Balance, expectedAccount.Balance)
	}
}
