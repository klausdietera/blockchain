package utils

import (
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestNewBlock(t *testing.T) {
	createdAt, err := time.Parse(time.RFC3339, "2019-09-23T22:08:41+00:00")
	if err != nil {
		t.Fatalf("Cannot parse time. Error: %s", err)
	}

	keyPair := GenerateKeyPair("hen worry two thank unfair salmon smile oven gospel grab latin reason")
	blockData := models.BlockData{
		CreatedAt: createdAt,
		Transactions: []models.Transaction{
			models.Transaction{
				Fee: 10000,
				Asset: &assets.Send{
					Amount:             100000000,
					RecipientPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
				},
			},
			models.Transaction{
				Fee: 20000,
				Asset: &assets.Send{
					Amount:             200000000,
					RecipientPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
				},
			},
		},
	}
	lastBlock := models.Block{
		ID:     "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		Height: 7,
	}

	block, err := NewBlock(blockData, lastBlock, keyPair)
	if err != nil {
		t.Fatalf("Cannot create block. Error: %s", err)
	}

	expectedID := "ee1a1c49e08021f5bc3bc90fb1f8179b79d59556487f74299cb23ccac9cdf390"
	if block.ID != expectedID {
		t.Errorf("Block id is invalid. Expected: %s, actual: %s", expectedID, block.ID)
	}

	expectedTransactionCount := int32(2)
	if block.TransactionCount != expectedTransactionCount {
		t.Errorf("Block transaction count is invalid. Expected: %d, actual: %d", expectedTransactionCount, block.TransactionCount)
	}

	if len(block.Transactions) != len(blockData.Transactions) {
		t.Errorf("Block transactions len is invalid. Expected: %d, actual: %d", len(blockData.Transactions), len(block.Transactions))
	}

	expectedSignature := "daf165dff511bd31d53677b29bc5b80f9bd137d27304496d731bdfa7b0d2a77fc9ed14cc135379ae0e99e27c5346446c0e59e79697be2d8eb64a06d6c0207804"
	if block.Signature != expectedSignature {
		t.Errorf("Block signature is invalid. Expected: %s, actual: %s", expectedSignature, block.Signature)
	}

	expectedGeneratorPublicKey := "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2"
	if block.GeneratorPublicKey != expectedGeneratorPublicKey {
		t.Errorf("Block generator public key is invalid. Expected: %s, actual: %s", expectedGeneratorPublicKey, block.GeneratorPublicKey)
	}

	expectedPayloadHash := "68c2ce1476056623c2046ad252255bc14bbc6333ab3c59f2e9e529f7668da7e3"
	if block.PayloadHash != expectedPayloadHash {
		t.Errorf("Block payload hash is invalid. Expected: %s, actual: %s", expectedPayloadHash, block.PayloadHash)
	}

	if block.PreviousBlockID != lastBlock.ID {
		t.Errorf("Block previous block id is invalid. Expected: %s, actual: %s", lastBlock.ID, block.PreviousBlockID)
	}

	if block.CreatedAt != blockData.CreatedAt {
		t.Errorf("Block createdAt is invalid. Expected: %s, actual: %s", blockData.CreatedAt, block.CreatedAt)
	}

	expectedAmount := int64(300000000)
	if block.Amount != expectedAmount {
		t.Errorf("Block amount is invalid. Expected: %d, actual: %d", expectedAmount, block.Amount)
	}

	expectedFee := int64(30000)
	if block.Fee != expectedFee {
		t.Errorf("Block fee is invalid. Expected: %d, actual: %d", expectedFee, block.Fee)
	}

	expectedHeight := uint64(8)
	if block.Height != expectedHeight {
		t.Errorf("Block height is invalid. Expected: %d, actual: %d", expectedHeight, block.Height)
	}
}
