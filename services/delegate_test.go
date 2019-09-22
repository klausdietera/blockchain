package services

import (
	"testing"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

func TestGetActive(t *testing.T) {
	activeDelegatesCount := uint16(3)
	Delegate = NewDelegate("", activeDelegatesCount)

	delegates := []*models.Account{
		{
			PublicKey: "1",
			Delegate: &models.Delegate{
				Votes: 10,
			},
		},
		{
			PublicKey: "2",
			Delegate: &models.Delegate{
				Votes: 20,
			},
		},
		{
			PublicKey: "3",
			Delegate: &models.Delegate{
				Votes: 40,
			},
		},
		{
			PublicKey: "4",
			Delegate: &models.Delegate{
				Votes: 30,
			},
		},
	}

	for _, delegate := range delegates {
		repositories.Delegates.Add(delegate)
	}

	active := Delegate.GetActive()

	actualLen := len(active)
	if actualLen != int(activeDelegatesCount) {
		t.Errorf("Incorrect active delegates count. Expected %d, actual %d", activeDelegatesCount, actualLen)
	}
}
