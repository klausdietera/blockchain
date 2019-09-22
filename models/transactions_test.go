package models_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
	"time"

	"bitbucket.org/axelsheva/blockchain/models"
	"bitbucket.org/axelsheva/blockchain/models/assets"
	"bitbucket.org/axelsheva/blockchain/models/types"
	"bitbucket.org/axelsheva/blockchain/repositories"
)

func TestSendApplyUnconfirmed(t *testing.T) {
	sender := models.Account{
		Balance:   100,
		PublicKey: "1",
	}
	recipient := models.Account{
		PublicKey: "2",
	}

	repositories.Accounts.Add(&recipient)

	transaction := models.Transaction{
		Fee: 1,
		Asset: &assets.Send{
			RecipientPublicKey: "2",
			Amount:             1,
		},
	}

	transaction.ApplyUnconfirmed(&sender)

	expectedSender := models.Account{
		Balance:   98,
		PublicKey: "1",
	}
	if sender.Balance != expectedSender.Balance {
		t.Errorf("Invalid sender balance. Actual %d, Expected %d", sender.Balance, expectedSender.Balance)
	}

	expectedRecipient := models.Account{
		Balance:   1,
		PublicKey: "2",
	}
	if recipient.Balance != expectedRecipient.Balance {
		t.Errorf("Invalid recipient balance. Actual %d, Expected %d", recipient.Balance, expectedRecipient.Balance)
	}
}

func TestGetBytes(t *testing.T) {
	transaction := models.Transaction{
		Type:            types.TransactionSend,
		Fee:             123,
		Salt:            "de69380839e87711a665aca932cb81d7",
		CreatedAt:       time.Now(),
		SenderPublicKey: "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc1",
		Asset: &assets.Send{
			Amount:             1234567890,
			RecipientPublicKey: "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2",
		},
	}

	b, err := transaction.GetBytes(true, true)
	if err != nil {
		t.Fatalf("Cannot get bytes from transaction. Error: %v.", err)
	}
	buf := bytes.NewReader(b)

	actualBufferSize := buf.Size()
	expectedSize := int64(180)
	if actualBufferSize != expectedSize {
		t.Errorf("Buffer size is incorrect. Actual %d, Expected %d", actualBufferSize, expectedSize)
	}

	var actualAmount int64
	expectedAmount := int64(1234567890)
	binary.Read(buf, binary.LittleEndian, &actualAmount)
	if actualAmount != expectedAmount {
		t.Errorf("Send amount is incorrect. Actual %d, Expected %d", actualAmount, expectedAmount)
	}

	actualRecipientPublicKey := string(b[8:72])
	expectedRecipientPublicKey := "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2"
	if actualRecipientPublicKey != expectedRecipientPublicKey {
		t.Errorf("Recipient public key is incorrect. Actual %s, Expected %s", actualRecipientPublicKey, expectedRecipientPublicKey)
	}

	var actualType types.TransactionType
	expectedType := types.TransactionSend
	buf.Seek(72, io.SeekStart)
	binary.Read(buf, binary.LittleEndian, &actualType)
	if actualType != expectedType {
		t.Errorf("Transaction type is incorrect. Actual %d, Expected %d", actualType, expectedType)
	}

	var actualCreatedAt int64
	expectedCreatedAt := transaction.CreatedAt.Unix()
	binary.Read(buf, binary.LittleEndian, &actualCreatedAt)
	if actualCreatedAt != expectedCreatedAt {
		t.Errorf("CreatedAt is incorrect. Actual %d, Expected %d", actualCreatedAt, expectedCreatedAt)
	}

	var actualSalt string = string(b[84:116])
	expectedSalt := transaction.Salt
	if actualSalt != expectedSalt {
		t.Errorf("Salt is incorrect. Actual %s, Expected %s", actualSalt, expectedSalt)
	}

	actualSenderPublicKey := string(b[116:180])
	expectedSenderPublicKey := "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc1"
	if actualSenderPublicKey != expectedSenderPublicKey {
		t.Errorf("Sender public key is incorrect. Actual %s, Expected %s", actualSenderPublicKey, expectedSenderPublicKey)
	}
}
