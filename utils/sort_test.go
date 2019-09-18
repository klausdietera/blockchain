package utils

import (
	"reflect"
	"sort"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/types"
)

func TestBlockSort(t *testing.T) {
	transactions := []*models.Transaction{
		{Type: types.TransactionSend, CreatedAt: time.Unix(40, 0), ID: "b"},
		{Type: types.TransactionSend, CreatedAt: time.Unix(40, 0), ID: "a"},
		{Type: types.TransactionSend, CreatedAt: time.Unix(10, 0), ID: "h"},
		{Type: types.TransactionReferral, CreatedAt: time.Unix(50, 0), ID: "z"},
	}

	expected := []*models.Transaction{
		transactions[3],
		transactions[2],
		transactions[1],
		transactions[0],
	}

	sort.Sort(BlockSort(transactions))

	isEqual := reflect.DeepEqual(transactions, expected)
	if !isEqual {
		t.Error("Transactions sorting is incorrect")
	}
}

func TestDelegatesSort(t *testing.T) {
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

	expected := []*models.Account{
		delegates[2],
		delegates[3],
		delegates[1],
		delegates[0],
	}

	sort.Sort(sort.Reverse(ByVotes(delegates)))

	isEqual := reflect.DeepEqual(delegates, expected)
	if !isEqual {
		t.Error("Delegates sorting is incorrect")
	}
}
