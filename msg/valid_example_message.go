package msg

import "github.com/bahner/go-ma"

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "zF4oNn5r5xe8hR5ttaes2sfM6ZD5oZNrhiXMVw3Lty7ZJgYxigKpkqETHx5N4tYUeCjn6RrkHMZfu4uC3GQCXBxzSCdSQpam3QfK3Xdkz6YSUyKd9mqzy87W76Nkce7RTbdMmtw4WdyqMPUPLPtoKKqrpbpvw2GnDoMBytnotDEvNDsP2BS4WU4E5epYj1hZM7Zbns7PV8pbwDtWdToWdHhpMBkbfYowWVtjCkkKeHCqpgoMrD1NEkyd9MBytgpNjwfCHNgFWJzQJZsMwx6LC3G7Wja5B9vgqrfN91PeL5HnPKCzpxeBLSYp3op9jdZJwXjG85YB8htDPjnBZdDLk6avNftNmHABkVVD3rX"
)

func ValidExampleMessage() *Message {

	msg := &Message{
		ID:           "CT6EklGVDpQpaYrth_O80",
		MimeType:     MIME_TYPE,
		From:         "did:ma:k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd#bahner",
		To:           "did:ma:k51qzi5uqu5dk3pkcowsu2jqmnby0ry551xud502v000dzftwf4bj68384j84l#job",
		Created:      1698684192,
		Expires:      1698687792,
		BodyMimeType: "text/plain",
		Body:         []byte("Share and Enjoy!"),
		Version:      ma.VERSION,
		Signature:    "",
	}

	return msg

}
