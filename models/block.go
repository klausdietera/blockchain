package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/jamesruan/sodium"
)

type BlockData struct {
	Transactions []Transaction `json:"transactions"`
	CreatedAt    time.Time     `json:"createdAt"`
}

type Block struct {
	ID                 string        `json:"id"`
	PreviousBlockID    string        `json:"previousBlockId"`
	PayloadHash        string        `json:"payloadHash"`
	Signature          string        `json:"signature"`
	GeneratorPublicKey string        `json:"generatorPublicKey"`
	Height             uint64        `json:"height"`
	Amount             int64         `json:"amount"`
	Fee                int64         `json:"fee"`
	Transactions       []Transaction `json:"transactions"`
	TransactionCount   int32         `json:"transactionCount"`
	Version            uint8         `json:"version"`
	CreatedAt          time.Time     `json:"createdAt"`
}

type IBlockController interface {
	OnReceive(block *Block) error
	Generate() (*Block, error)
}

func (b *Block) CalculatePayloadHash() (string, error) {
	buf := new(bytes.Buffer)
	for _, transaction := range b.Transactions {
		b, err := transaction.GetBytes(false, false)
		if err != nil {
			return "", err
		}

		buf.Write(b)
	}

	hash := sha256.Sum256(buf.Bytes())

	return hex.EncodeToString(hash[:]), nil
}

func (b *Block) GetBytes(skipSignature bool) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, b.Version)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	err = binary.Write(buf, binary.LittleEndian, b.CreatedAt.Unix())
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	err = binary.Write(buf, binary.LittleEndian, b.TransactionCount)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	err = binary.Write(buf, binary.LittleEndian, b.Amount)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	err = binary.Write(buf, binary.LittleEndian, b.Fee)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
		return buf.Bytes(), err
	}

	buf.Write([]byte(b.PreviousBlockID))
	buf.Write([]byte(b.PayloadHash))
	buf.Write([]byte(b.GeneratorPublicKey))

	if !skipSignature {
		buf.Write([]byte(b.Signature))
	}

	return buf.Bytes(), nil
}

func (s *Block) CalculateHash(skipSignature bool) ([32]byte, error) {
	b, err := s.GetBytes(skipSignature)
	if err != nil {
		var emptyBytes [32]byte
		return emptyBytes, err
	}

	return sha256.Sum256(b), nil
}

func (b *Block) CalculateSignature(keyPair sodium.SignKP) (string, error) {
	hash, err := b.CalculateHash(true)
	if err != nil {
		return "", err
	}

	bytes := sodium.Bytes(hash[:])
	signature := bytes.SignDetached(keyPair.SecretKey)
	return hex.EncodeToString(signature.Bytes), nil
}

func (s *Block) CalculateID() (string, error) {
	hash, err := s.CalculateHash(false)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash[:]), nil
}
