package message

import "github.com/bahner/go-ma"

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "zDrMedu2xgJrJ74Ht9BTscCyRX4wtYq9UjfZbNMK5s3CNiwqHBrWpGMzRBvpVz5AaSYYYWHV1eU3yKQSo7AdXhQgwamttGZorxp3ihXMjW6wzKP3sEzahgb2PfooQJrSQdfQhCKiLUB6kYFxxgxvaWQWsxABtxopz988GQvcxpiRsfAcH6SiydJLtYXbbgAiGnUMaX9pZHve5E6WTVaQKtm9BhkhJKwLE5s8FtTBKQ9WVm6GBRHwboSS8mGdg4aUZxuXmeuJa3KvHD5gw9oA2curF6EKnD3NxHjeWLbTB5rgtZYMfCcTzyUfmARBWBosY4eKDYfuEZKycTNJK8yMrzWb"
)

func ValidExampleMessage() *Message {

	msg := &Message{
		ID:           "CT6EklGVDpQpaYrth_O80",
		MimeType:     ma.MESSAGE_MIME_TYPE,
		From:         "did:ma:k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd",
		To:           "did:ma:k51qzi5uqu5dk3pkcowsu2jqmnby0ry551xud502v000dzftwf4bj68384j84l",
		Created:      1698684192,
		Expires:      1698687792,
		BodyMimeType: "text/plain",
		Body:         "Share and Enjoy!",
		Version:      ma.VERSION,
		Signature:    "",
	}

	return msg

}
