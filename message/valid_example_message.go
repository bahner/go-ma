package message

import "github.com/bahner/go-ma"

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "z2o53MHUTjgyaJ8upxousyh8RoHc7JMnTCTwchgdBV7j13p9kyDqbw1GyehDyPXoQ1ZPPZaWtL2QuVpnYhoHQjH4B2W9PvmHKYyAGyyjk5r6ba3YExiaKhCViqwosCowVg3YqAiZWo5Po7Yvejhp7vATzAx3SQUkMHTHQYH4M342m7Y7iBy2RU1sehYkkbuMrDGGC4VzbS15vioqCxmUmeazvx598PC9Z8NLnvaNhgpqoDWifNnyX4tuBaHX2KqHQFEee9bDRMeVgBKhEdUKT5kBQxzsmp1z87HPjpwn6zDJQR5r3qiv4xiq9oQs35eUvq7m"
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
