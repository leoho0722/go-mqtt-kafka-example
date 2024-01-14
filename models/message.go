package models

type Message struct {

	// Body of the message
	Body []byte
}

var message Message

// CacheMessage caches the message struct
func CacheMessage(payload []byte) {
	message.Body = payload
}

// GetMessage returns the message struct
func GetMessage() Message {
	return message
}
