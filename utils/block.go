package utils

import (
	"encoding/hex"

	"bitbucket.org/axelsheva/blockchain/models"
	"github.com/jamesruan/sodium"
)

func NewBlock(data models.BlockData, lastBlock models.Block, keyPair sodium.SignKP) (models.Block, error) {
	var totalFee int64
	var totalAmount int64
	for _, transaction := range data.Transactions {
		totalFee += transaction.Fee
		totalAmount += transaction.Asset.GetAmount()
	}

	block := models.Block{
		Transactions:       data.Transactions,
		Fee:                totalFee,
		Amount:             totalAmount,
		CreatedAt:          data.CreatedAt,
		TransactionCount:   int32(len(data.Transactions)),
		GeneratorPublicKey: hex.EncodeToString(keyPair.PublicKey.Bytes),
		PreviousBlockID:    lastBlock.ID,
		Height:             lastBlock.Height + 1,
	}

	payloadHash, err := block.CalculatePayloadHash()
	if err != nil {
		return block, err
	}
	block.PayloadHash = payloadHash

	signature, err := block.CalculateSignature(keyPair)
	if err != nil {
		return block, err
	}
	block.Signature = signature

	id, err := block.CalculateID()
	if err != nil {
		return block, err
	}
	block.ID = id

	return block, nil
}
