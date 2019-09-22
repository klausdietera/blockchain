package assets_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"bitbucket.org/axelsheva/blockchain/models/assets"
)

func TestSendCalculateFee(t *testing.T) {
	send := assets.Send{
		Amount: 100000000,
	}

	got := send.CalculateFee()
	want := int64(10000)

	if got != want {
		t.Errorf("Send fee calculation is incorrect. Actual %d, Expected %d", got, want)
	}
}

func TestSendGetBytes(t *testing.T) {
	sendAsset := assets.Send{
		Amount:             1234567890,
		RecipientPublicKey: "f4ae589b02f97e9ab5bce61cf187bcc96cfb3fdf9a11333703a682b7d47c8dc2",
	}

	b := sendAsset.GetBytes()
	buf := bytes.NewReader(b)

	var actualAmount int64
	binary.Read(buf, binary.LittleEndian, &actualAmount)
	if sendAsset.Amount != actualAmount {
		t.Errorf("Send amount is incorrect. Actual %d, Expected %d", actualAmount, sendAsset.Amount)
	}

	actualRecipientPublicKey := string(b[8:])
	if sendAsset.RecipientPublicKey != actualRecipientPublicKey {
		t.Errorf("Recipient public key is incorrect. Actual %s, Expected %s", actualRecipientPublicKey, sendAsset.RecipientPublicKey)
	}
}
