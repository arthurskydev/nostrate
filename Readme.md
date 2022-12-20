# Lightweight Go library for building Nostr applications

### Creating and signing an event
```go
event, err := nostrate.NewEvent(PUBKEY, KIND, TAGS, CONTENT)
err = event.Sign(PRIVATE_KEY)
```