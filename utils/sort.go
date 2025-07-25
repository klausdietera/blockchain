package utils

import (
	"bitbucket.org/axelsheva/blockchain/models"
)

type BlockSort []models.Transaction

func (a BlockSort) Len() int      { return len(a) }
func (a BlockSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BlockSort) Less(i, j int) bool {
	if a[i].Type < a[j].Type {
		return a[i].Type < a[j].Type
	}
	if a[i].Type > a[j].Type {
		return a[i].Type > a[j].Type
	}
	if a[i].CreatedAt.Unix() < a[j].CreatedAt.Unix() {
		return a[i].CreatedAt.Unix() < a[j].CreatedAt.Unix()
	}
	if a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix() {
		return a[i].CreatedAt.Unix() > a[j].CreatedAt.Unix()
	}
	if a[i].ID < a[j].ID {
		return a[i].ID < a[j].ID
	}
	if a[i].ID > a[j].ID {
		return a[i].ID > a[j].ID
	}
	return false
}

type ByVotes []*models.Account

func (a ByVotes) Len() int           { return len(a) }
func (a ByVotes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVotes) Less(i, j int) bool { return a[i].Delegate.Votes < a[j].Delegate.Votes }

type HashSort []*models.HashList

func (a HashSort) Len() int           { return len(a) }
func (a HashSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a HashSort) Less(i, j int) bool { return a[i].Hash < a[j].Hash }
