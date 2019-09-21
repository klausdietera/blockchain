package utils

import (
	"reflect"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
)

func TestCreateSendTransaction(t *testing.T) {
	keyPair := GenerateKeyPair("hen worry two thank unfair salmon smile oven gospel grab latin reason")

	createdAt, err := time.Parse(time.RFC3339, "2019-09-23T22:08:41+00:00")
	if err != nil {
		t.Fatalf("Cannot parse time. Error: %s", err)
	}

	transaction, err := CreateTransaction(models.Transaction{
		Type: types.TransactionSend,
		Asset: &assets.Send{
			Amount:             100000000,
			RecipientPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		},
		Salt:      "spDdryEuGzeNulzOBgZHekfOIByRllBe",
		CreatedAt: createdAt,
	}, keyPair, nil)

	if err != nil {
		t.Fatalf("Cannot create transaction. Error: %s", err)
	}

	expectedFee := int64(10000)
	if transaction.Fee != expectedFee {
		t.Errorf("Transaction fee is invalid. Expected: %d, actual: %d", expectedFee, transaction.Fee)
	}

	if transaction.Salt == "" {
		t.Errorf("Transaction salt is missing")
	}

	expectedSalt := "spDdryEuGzeNulzOBgZHekfOIByRllBe"
	if transaction.Salt != expectedSalt {
		t.Errorf("Transaction salt is invalid. Expected: %s, actual: %s", expectedSalt, transaction.Salt)
	}

	if len(transaction.Salt) != models.SALT_LENGTH {
		t.Errorf("Transaction salt has invalid length. Expected: %d, actual: %d", models.SALT_LENGTH, len(transaction.Salt))
	}

	if transaction.Type != types.TransactionSend {
		t.Errorf("Transaction type is invalid. Expected: %d, actual: %d", types.TransactionSend, transaction.Type)
	}

	expectedSignature := "fd058d3835ce4466579bedaa94ae7fcf26df6dd40fc0de63ca34c9eb01dd9c006d7c009c40a43d4087c11d15326ef9a7ffee8d43931681e99c11ffed00074702"
	if transaction.Signature != expectedSignature {
		t.Errorf("Transaction signature is invalid. Expected: %s, actual: %s", expectedSignature, transaction.Signature)
	}

	expectedSecondSignature := ""
	if transaction.SecondSignature != expectedSecondSignature {
		t.Errorf("Transaction second signature is invalid. Expected: %s, actual: %s", expectedSecondSignature, transaction.SecondSignature)
	}

	expectedID := "18b3535192e4585f34043477b956541748b41d42fa498261d223a6e14ac4d2f1"
	if transaction.ID != expectedID {
		t.Errorf("Transaction id is invalid. Expected: %s, actual: %s", expectedID, transaction.ID)
	}

	if transaction.CreatedAt != createdAt {
		t.Errorf("Transaction createdAt is invalid. Expected: %s, actual: %s", createdAt, transaction.CreatedAt)
	}

	expectedBlockID := ""
	if transaction.BlockID != expectedBlockID {
		t.Errorf("Transaction blockId is invalid. Expected: %s, actual: %s", expectedBlockID, transaction.BlockID)
	}

	expectedSenderPublicKey := "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2"
	if transaction.SenderPublicKey != expectedSenderPublicKey {
		t.Errorf("Transaction sender public key is invalid. Expected: %s, actual: %s", expectedSenderPublicKey, transaction.SenderPublicKey)
	}
}

func TestSendTransactionUnmarshalJSON(t *testing.T) {
	bytes := []byte(`{"id":"c7d80bf1bb220e62735bd388549a87c0cd93b8be30a1ae2f7291ce20d2a94b79","blockId":"cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0","type":10,"createdAt":"2019-01-01T00:00:00.000Z","senderPublicKey":"49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a","signature":"226ed984bf3d82b7c332ce48bc976fcc35930d22cb068b2e9de993a4fb3e402d4bdb7077d0923b8dd2c205e6a2473884752615c0787967b218143eec5df1390c","fee":10,"salt":"a7fdae234eeb416e31f5f02571f54a0c","asset":{"recipientPublicKey":"49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a","amount":4500000000000000}}`)

	transaction := models.Transaction{}
	err := UnmarshalTransaction(bytes, &transaction)
	if err != nil {
		panic(err)
	}

	createdAt, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00.000Z")
	if err != nil {
		panic(err)
	}
	expectedTransaction := models.Transaction{
		ID:              "c7d80bf1bb220e62735bd388549a87c0cd93b8be30a1ae2f7291ce20d2a94b79",
		BlockID:         "cbb9449abb9672d33fa2eb200b1c8b03db7c6572dfb6e59dc334c0ab82b63ab0",
		Type:            types.TransactionSend,
		CreatedAt:       createdAt,
		Signature:       "226ed984bf3d82b7c332ce48bc976fcc35930d22cb068b2e9de993a4fb3e402d4bdb7077d0923b8dd2c205e6a2473884752615c0787967b218143eec5df1390c",
		SenderPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		Salt:            "a7fdae234eeb416e31f5f02571f54a0c",
		Fee:             10,
		Asset: &assets.Send{
			Amount:             4500000000000000,
			RecipientPublicKey: "49a2b5e68f851a11058748269a276b0c0d36497215548fb40d4fe4e929d0283a",
		},
	}

	if !reflect.DeepEqual(transaction, expectedTransaction) {
		t.Errorf("Invalid send transaction unmarshal. Expected: %+v, actual: %+v", expectedTransaction, transaction)
	}
}
