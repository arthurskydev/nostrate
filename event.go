package nostrate

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/goccy/go-json"
)

type Event struct {
	Id        string    `json:"id"`
	PubKey    string    `json:"pubkey"`
	CreatedAt time.Time `json:"created_at"`
	Kind      Kind      `json:"kind"`
	Tags      Tags      `json:"tags"`
	Content   string    `json:"content"`
	Signature string    `json:"sig,omitempty"`
}

type Tags map[string][]string

// Kind describes that action or meaning of an event.
type Kind int

const (
	SetMetadata            Kind = 0
	TextNote               Kind = 1
	RecommendServer        Kind = 2
	ContactList            Kind = 3
	EncryptedDirectMessage Kind = 4
	Deletion               Kind = 5
	Boost                  Kind = 6
	Reaction               Kind = 7
	ChannelCreation        Kind = 40
	ChannelMetadata        Kind = 41
	ChannelMessage         Kind = 42
	ChannelHideMessage     Kind = 43
	ChannelMuteUser        Kind = 44
)

func (e *Event) GetHash() ([]byte, error) {
	buf, err := json.Marshal(&e)
	if err != nil {
		return nil, err
	}
	h := sha256.Sum256(buf)
	return h[:], err
}

func NewEvent(pubKey string, kind Kind, tags Tags, content string) (*Event, error) {
	ev := Event{PubKey: pubKey, Kind: kind, Tags: tags, Content: content}
	ev.CreatedAt = time.Now()
	return &ev, nil
}

func (e *Event) Sign(privateKey string) error {
	h, err := e.GetHash()
	if err != nil {
		return err
	}
	pks, err := hex.DecodeString(privateKey)
	if err != nil {
		return fmt.Errorf("attempted to sign event with an invalid private key %w", err)
	}
	pk, _ := btcec.PrivKeyFromBytes(pks)
	sig, err := schnorr.Sign(pk, h)
	if err != nil {
		return err
	}
	e.Id = hex.EncodeToString(h)
	e.Signature = hex.EncodeToString(sig.Serialize())
	return nil
}
