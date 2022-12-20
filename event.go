package nostrate

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/goccy/go-json"
	"time"
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
	KindSetMetadata            Kind = 0
	KindTextNote               Kind = 1
	KindRecommendServer        Kind = 2
	KindContactList            Kind = 3
	KindEncryptedDirectMessage Kind = 4
	KindDeletion               Kind = 5
	KindBoost                  Kind = 6
	KindReaction               Kind = 7
	KindChannelCreation        Kind = 40
	KindChannelMetadata        Kind = 41
	KindChannelMessage         Kind = 42
	KindChannelHideMessage     Kind = 43
	KindChannelMuteUser        Kind = 44
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
